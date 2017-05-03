///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

// NO TEST

package cluster

import (
	"errors"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/supervised-io/kov/gen/client/operations"
	"github.com/supervised-io/kov/gen/models"
	"github.com/supervised-io/kov/pkg/kovclient"
)

// ClusterAPI interface
type ClusterAPI interface {
	CreateCluster(string, *models.ClusterConfig) (*string, error)
	GetTaskStatus(string, string) (bool, error)
}

// Client Cluster client for sending cluster request
type client struct {
}

// NewClusterClient return a cluster client
func NewClusterClient() ClusterAPI {
	return &client{}
}

// CreateCluster Sent create cluster request and return task id or error
func (c *client) CreateCluster(url string, clusterConfig *models.ClusterConfig) (*string, error) {
	params := operations.NewCreateClusterParams().WithClusterConfig(clusterConfig)

	kovClient, err := kovclient.GetClient(url)
	if err != nil {
		return nil, err
	}

	// send request
	resp, err := kovClient.CreateCluster(params)

	if err != nil {
		var payload *models.Error
		switch etp := err.(type) {
		case *operations.CreateClusterConflict:
			payload = etp.Payload
		default:
			return nil, fmt.Errorf("%s: %s \n", "Failed cluster create", err.Error())
		}

		if swag.StringValue(payload.Message) == "" {
			msg := err.Error()
			payload.Message = swag.String(msg)
		}
		return nil, fmt.Errorf(err.Error(), swag.StringValue(payload.Message), swag.Int64Value(payload.Code))
	}

	taskID := string(resp.Payload)
	return &taskID, nil

}

// GetTaskStatus process task status response
func (c *client) GetTaskStatus(url, taskID string) (bool, error) {
	params := operations.NewGetTaskParams().WithTaskid(taskID)
	kovClient, err := kovclient.GetClient(url)
	if err != nil {
		return false, err
	}

	resp, err := kovClient.GetTask(params)
	if err != nil {
		var payload *models.Error
		switch etp := err.(type) {
		case *operations.GetTaskNotFound:
			payload = etp.Payload
			return false, nil
		case *operations.GetTaskDefault:
			payload = etp.Payload
			code := etp.Code()
			if code/100 != 2 {
				return false, errors.New("Error get task: GetTaskDefault")
			}
		default:
			return false, err
		}
		if swag.StringValue(payload.Message) == "" {
			msg := err.Error()
			payload.Message = swag.String(msg)
		}

		return false, fmt.Errorf("Error : %s Code: %d", swag.StringValue(payload.Message), swag.Int64Value(payload.Code))
	}
	// if get response without error
	switch resp.Payload.State {
	case models.TaskStateProcessing:
		{
			return false, nil
		}
	case models.TaskStateFailed:
		{
			if resp.Payload.Context != nil {
				if resp.Payload.Context.Log == "" {
					return true, errors.New(resp.Payload.Context.Cause)
				}
				return true, errors.New(resp.Payload.Context.Log)
			}
			return true, fmt.Errorf(resp.Error())
		}
	case models.TaskStateCompleted:
		{
			return true, nil
		}
	}
	return false, nil
}
