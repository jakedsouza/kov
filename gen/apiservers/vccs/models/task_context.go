package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaskContext the context for a task, contains data to describe what this job pertained to.
// swagger:model taskContext
type TaskContext struct {

	// cause
	Cause string `json:"cause,omitempty"`

	// cluster name
	// Max Length: 63
	// Min Length: 3
	// Pattern: ^[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?(\.[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?)*$
	ClusterName string `json:"clusterName,omitempty"`

	// log
	Log string `json:"log,omitempty"`

	// timeout
	Timeout *int64 `json:"timeout,omitempty"`
}

// Validate validates this task context
func (m *TaskContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskContext) validateClusterName(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterName) { // not required
		return nil
	}

	if err := validate.MinLength("clusterName", "body", string(m.ClusterName), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("clusterName", "body", string(m.ClusterName), 63); err != nil {
		return err
	}

	if err := validate.Pattern("clusterName", "body", string(m.ClusterName), `^[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?(\.[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?)*$`); err != nil {
		return err
	}

	return nil
}