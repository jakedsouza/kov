///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

//NO TESTS

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
	CreateCluster(*models.ClusterConfig) (*string, error)
	GetTaskStatus(string) (bool, error)
}

// Client Cluster client for sending cluster request
type client struct {
	kovClient *operations.Client
}

// NewClusterClient return a cluster client
func NewClusterClient(url string) (ClusterAPI, error) {
	kovClient, err := kovclient.GetClient(url)
	if err != nil {
		return nil, err
	}
	return &client{kovClient: kovClient}, nil
}

// CreateCluster Sent create cluster request and return task id or error
func (c *client) CreateCluster(clusterConfig *models.ClusterConfig) (*string, error) {
	params := operations.NewCreateClusterParams().WithClusterConfig(clusterConfig)

	// send request
	resp, err := c.kovClient.CreateCluster(params)

	if err != nil {
		var payload *models.Error
		switch etp := err.(type) {
		case *operations.CreateClusterDefault:
			payload = etp.Payload
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
func (c *client) GetTaskStatus(taskID string) (bool, error) {
	params := operations.NewGetTaskParams().WithTaskid(taskID)

	resp, err := c.kovClient.GetTask(params)
	if err != nil {
		var payload *models.Error
		switch etp := err.(type) {
		case *operations.GetTaskNotFound:
			payload = etp.Payload
		case *operations.GetTaskDefault:
			payload = etp.Payload
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
