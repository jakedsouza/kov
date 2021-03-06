///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/supervised-io/kov/gen/models"
)

// GetTaskReader is a Reader for the GetTask structure.
type GetTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTaskOK creates a GetTaskOK with default headers values
func NewGetTaskOK() *GetTaskOK {
	return &GetTaskOK{}
}

/*GetTaskOK handles this case with default header values.

the task for the given task id
*/
type GetTaskOK struct {
	Payload *models.Task
}

func (o *GetTaskOK) Error() string {
	return fmt.Sprintf("[GET /tasks/{taskid}][%d] getTaskOK  %+v", 200, o.Payload)
}

func (o *GetTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskNotFound creates a GetTaskNotFound with default headers values
func NewGetTaskNotFound() *GetTaskNotFound {
	return &GetTaskNotFound{}
}

/*GetTaskNotFound handles this case with default header values.

The task was not found
*/
type GetTaskNotFound struct {
	Payload *models.Error
}

func (o *GetTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /tasks/{taskid}][%d] getTaskNotFound  %+v", 404, o.Payload)
}

func (o *GetTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskDefault creates a GetTaskDefault with default headers values
func NewGetTaskDefault(code int) *GetTaskDefault {
	return &GetTaskDefault{
		_statusCode: code,
	}
}

/*GetTaskDefault handles this case with default header values.

Error
*/
type GetTaskDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get task default response
func (o *GetTaskDefault) Code() int {
	return o._statusCode
}

func (o *GetTaskDefault) Error() string {
	return fmt.Sprintf("[GET /tasks/{taskid}][%d] getTask default  %+v", o._statusCode, o.Payload)
}

func (o *GetTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
