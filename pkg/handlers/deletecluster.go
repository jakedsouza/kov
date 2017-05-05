///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/supervised-io/kov/gen/models"
	"github.com/supervised-io/kov/gen/restapi/operations"
)

// DeleteCluster is the handler of the create-cluster operation
type DeleteCluster struct {
	//app kov.Application
}

// Handle is the handler for the create cluser operation
func (h *DeleteCluster) Handle(params operations.DeleteClusterParams) middleware.Responder {
	// this is a dummy/stub implementation
	tid := models.TaskID("1111-2222-3333-4444")
	return operations.NewDeleteClusterAccepted().WithPayload(tid)
}
