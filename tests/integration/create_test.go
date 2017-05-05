///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package integration

import (
	"testing"

	"net/http"

	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"github.com/supervised-io/kov/gen/client/operations"
	"github.com/supervised-io/kov/gen/models"
)

func assertCreateDefaultErrorResponse(t *testing.T, resp *operations.CreateClusterAccepted, err error, expCode int) {
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	if assert.IsType(t, &operations.CreateClusterDefault{}, err) {
		assert.Equal(t, http.StatusUnprocessableEntity, err.(*operations.CreateClusterDefault).Code())
	}
}

func getValidClusterConfig(clusterName string) *models.ClusterConfig {
	return &models.ClusterConfig{
		Name:              swag.String(clusterName),
		ManagementNetwork: swag.String("dummy"),
		NoOfMasters:       swag.Int32(1),
		MinNodes:          swag.Int32(1),
		ResourcePool:      "pool",
	}
}

func TestCreateZeroMinNode(t *testing.T) {
	cli := kovClient()
	cfg := getValidClusterConfig("cluster1")
	cfg.MinNodes = swag.Int32(0)
	params := operations.NewCreateClusterParams().WithClusterConfig(cfg)
	resp, err := cli.Operations.CreateCluster(params)
	assertCreateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestCreateNegativeMinNode(t *testing.T) {
	cli := kovClient()
	cfg := getValidClusterConfig("cluster1")
	cfg.MinNodes = swag.Int32(-5)
	params := operations.NewCreateClusterParams().WithClusterConfig(cfg)
	resp, err := cli.Operations.CreateCluster(params)
	assertCreateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestCreateNoMinNode(t *testing.T) {
	cli := kovClient()
	cfg := getValidClusterConfig("cluster1")
	cfg.MinNodes = nil
	params := operations.NewCreateClusterParams().WithClusterConfig(cfg)
	resp, err := cli.Operations.CreateCluster(params)
	assertCreateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestCreateNoConfig(t *testing.T) {
	cli := kovClient()
	params := operations.NewCreateClusterParams()
	resp, err := cli.Operations.CreateCluster(params)
	assertCreateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestCreateNoManagementNetwork(t *testing.T) {
	cli := kovClient()
	cfg := getValidClusterConfig("cluster1")
	cfg.ManagementNetwork = nil
	params := operations.NewCreateClusterParams().WithClusterConfig(cfg)
	resp, err := cli.Operations.CreateCluster(params)
	assertCreateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestCreateNoName(t *testing.T) {
	cli := kovClient()
	cfg := getValidClusterConfig("cluster1")
	cfg.Name = nil
	params := operations.NewCreateClusterParams().WithClusterConfig(cfg)
	resp, err := cli.Operations.CreateCluster(params)
	assertCreateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestCreateNoNumberMasters(t *testing.T) {
	cli := kovClient()
	cfg := getValidClusterConfig("cluster1")
	cfg.NoOfMasters = nil
	params := operations.NewCreateClusterParams().WithClusterConfig(cfg)
	resp, err := cli.Operations.CreateCluster(params)
	assertCreateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestCreateNoResourcePool(t *testing.T) {
	cli := kovClient()
	cfg := getValidClusterConfig("cluster1")
	cfg.ResourcePool = ""
	params := operations.NewCreateClusterParams().WithClusterConfig(cfg)
	resp, err := cli.Operations.CreateCluster(params)
	assertCreateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestCreateZeroMasters(t *testing.T) {
	cli := kovClient()
	cfg := getValidClusterConfig("cluster1")
	cfg.NoOfMasters = swag.Int32(0)
	params := operations.NewCreateClusterParams().WithClusterConfig(cfg)
	resp, err := cli.Operations.CreateCluster(params)
	assertCreateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestCreateNegativeMasters(t *testing.T) {
	cli := kovClient()
	cfg := getValidClusterConfig("cluster1")
	cfg.NoOfMasters = swag.Int32(-1)
	params := operations.NewCreateClusterParams().WithClusterConfig(cfg)
	resp, err := cli.Operations.CreateCluster(params)
	assertCreateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestCreateInvalidName(t *testing.T) {
	cli := kovClient()
	cfg := getValidClusterConfig("_ywegduywe")
	params := operations.NewCreateClusterParams().WithClusterConfig(cfg)
	resp, err := cli.Operations.CreateCluster(params)
	assertCreateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}
