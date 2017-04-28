package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/supervised-io/kov/gen/models"
	"github.com/supervised-io/kov/gen/restapi/operations"
)

// UpdateCluster is the handler of the create-cluster operation
type UpdateCluster struct {
	//app kov.Application
}

// Handle is the handler for the create cluser operation
func (h *UpdateCluster) Handle(params operations.UpdateClusterParams) middleware.Responder {
	// this is a dummy/stub implementation
	tid := models.TaskID("5555-5555-5555-5555")
	return operations.NewUpdateClusterAccepted().WithPayload(tid)
}
