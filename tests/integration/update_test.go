package integration

import (
	"testing"

	"net/http"

	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"github.com/supervised-io/kov/gen/client/operations"
	"github.com/supervised-io/kov/gen/models"
)

func assertUpdateDefaultErrorResponse(t *testing.T, resp *operations.UpdateClusterAccepted, err error, expCode int) {
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	if assert.IsType(t, &operations.UpdateClusterDefault{}, err) {
		assert.Equal(t, http.StatusUnprocessableEntity, err.(*operations.UpdateClusterDefault).Code())
	}
}

func getValidUpdateConfig(clusterName string) *models.ClusterUpdateConfig {
	return &models.ClusterUpdateConfig{
		Name:        clusterName,
		NoOfMasters: swag.Int32(1),
		MinNodes:    swag.Int32(1),
	}
}

func TestUpdateNoConfig(t *testing.T) {
	cli := kovClient()
	params := operations.NewUpdateClusterParams().WithName("dummy")
	resp, err := cli.Operations.UpdateCluster(params)
	assertUpdateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestUpdateNoName(t *testing.T) {
	cli := kovClient()
	cfg := getValidUpdateConfig("cluster1")
	cfg.Name = ""
	params := operations.NewUpdateClusterParams().WithClusterUpdateConfig(cfg).WithName("dummy")
	resp, err := cli.Operations.UpdateCluster(params)
	assertUpdateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestUpdateInvalidName(t *testing.T) {
	cli := kovClient()
	cfg := getValidUpdateConfig("cluster1")
	cfg.Name = "_whateverinvalidname"
	params := operations.NewUpdateClusterParams().WithClusterUpdateConfig(cfg).WithName("dummy")
	resp, err := cli.Operations.UpdateCluster(params)
	assertUpdateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestUpdateNoMasterNodes(t *testing.T) {
	cli := kovClient()
	cfg := getValidUpdateConfig("cluster1")
	cfg.NoOfMasters = nil
	params := operations.NewUpdateClusterParams().WithClusterUpdateConfig(cfg).WithName("dummy")
	resp, err := cli.Operations.UpdateCluster(params)
	assertUpdateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestUpdateZeroMasterNodes(t *testing.T) {
	cli := kovClient()
	cfg := getValidUpdateConfig("cluster1")
	cfg.NoOfMasters = swag.Int32(0)
	params := operations.NewUpdateClusterParams().WithClusterUpdateConfig(cfg).WithName("dummy")
	resp, err := cli.Operations.UpdateCluster(params)
	assertUpdateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestUpdateNegativeMasterNodes(t *testing.T) {
	cli := kovClient()
	cfg := getValidUpdateConfig("cluster1")
	cfg.NoOfMasters = swag.Int32(-5)
	params := operations.NewUpdateClusterParams().WithClusterUpdateConfig(cfg).WithName("dummy")
	resp, err := cli.Operations.UpdateCluster(params)
	assertUpdateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestUpdateNoMinNodes(t *testing.T) {
	cli := kovClient()
	cfg := getValidUpdateConfig("cluster1")
	cfg.MinNodes = nil
	params := operations.NewUpdateClusterParams().WithClusterUpdateConfig(cfg).WithName("dummy")
	resp, err := cli.Operations.UpdateCluster(params)
	assertUpdateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestUpdateZeroMinNodes(t *testing.T) {
	cli := kovClient()
	cfg := getValidUpdateConfig("cluster1")
	cfg.MinNodes = swag.Int32(0)
	params := operations.NewUpdateClusterParams().WithClusterUpdateConfig(cfg).WithName("dummy")
	resp, err := cli.Operations.UpdateCluster(params)
	assertUpdateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}

func TestUpdateNegativeMinNodes(t *testing.T) {
	cli := kovClient()
	cfg := getValidUpdateConfig("cluster1")
	cfg.MinNodes = swag.Int32(-12)
	params := operations.NewUpdateClusterParams().WithClusterUpdateConfig(cfg).WithName("dummy")
	resp, err := cli.Operations.UpdateCluster(params)
	assertUpdateDefaultErrorResponse(t, resp, err, http.StatusUnprocessableEntity)
}
