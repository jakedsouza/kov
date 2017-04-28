package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/supervised-io/kov/gen/models"
	"github.com/supervised-io/kov/gen/restapi/operations"
)

// ListClusters is the handler of the list-cluster operation
type ListClusters struct {
	//app kov.Application
}

// Handle is the handler for the list cluser operation
func (h *ListClusters) Handle(params operations.ListClustersParams) middleware.Responder {
	// this is a dummy/stub implementation
	clusters := []*models.Cluster{
		{
			Status: models.ClusterStatusActive,
			Config: &models.ClusterConfig{
				Name:       swag.String("cluster1"),
				MasterSize: models.InstanceSizeHuge,
				MaxNodes:   5,
				MinNodes:   swag.Int32(2),
			},
		},
		{
			Status: models.ClusterStatusActive,
			Config: &models.ClusterConfig{
				Name:       swag.String("cluster2"),
				MasterSize: models.InstanceSizeSmall,
				MaxNodes:   3,
				MinNodes:   swag.Int32(1),
			},
		},
	}

	return operations.NewListClustersOK().WithPayload(clusters)
}
