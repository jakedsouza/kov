package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/supervised-io/kov/gen/models"
	"github.com/supervised-io/kov/gen/restapi/operations"
)

// GetTask is the handler of the get-task operation
type GetTask struct {
	//app kov.Application
}

// Handle is the handler for the get-task operation
func (h *GetTask) Handle(params operations.GetTaskParams) middleware.Responder {
	// this is a dummy/stub implementation
	task := &models.Task{
		ID:       models.TaskID(params.Taskid),
		TaskType: models.TaskTypeCreate,
		State:    models.TaskStateProcessing,
	}
	return operations.NewGetTaskOK().WithPayload(task)
}
