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

// CreateCluster is the handler of the create-cluster operation
type CreateCluster struct {
	//app kov.Application
}

// Handle is the handler for the create cluser operation
func (h *CreateCluster) Handle(params operations.CreateClusterParams) middleware.Responder {
	// this is a dummy/stub implementation
	tid := models.TaskID("1234-5678-1234-5678")
	return operations.NewCreateClusterAccepted().WithPayload(tid)
}
