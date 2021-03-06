///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/supervised-io/kov/gen/models"
)

// NewCreateClusterParams creates a new CreateClusterParams object
// with the default values initialized.
func NewCreateClusterParams() *CreateClusterParams {
	var ()
	return &CreateClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterParamsWithTimeout creates a new CreateClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateClusterParamsWithTimeout(timeout time.Duration) *CreateClusterParams {
	var ()
	return &CreateClusterParams{

		timeout: timeout,
	}
}

// NewCreateClusterParamsWithContext creates a new CreateClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateClusterParamsWithContext(ctx context.Context) *CreateClusterParams {
	var ()
	return &CreateClusterParams{

		Context: ctx,
	}
}

// NewCreateClusterParamsWithHTTPClient creates a new CreateClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateClusterParamsWithHTTPClient(client *http.Client) *CreateClusterParams {
	var ()
	return &CreateClusterParams{
		HTTPClient: client,
	}
}

/*CreateClusterParams contains all the parameters to send to the API endpoint
for the create cluster operation typically these are written to a http.Request
*/
type CreateClusterParams struct {

	/*XRequestID
	  A unique UUID for the request

	*/
	XRequestID *string
	/*ClusterConfig
	  the config of the cluster to be created

	*/
	ClusterConfig *models.ClusterConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create cluster params
func (o *CreateClusterParams) WithTimeout(timeout time.Duration) *CreateClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster params
func (o *CreateClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster params
func (o *CreateClusterParams) WithContext(ctx context.Context) *CreateClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster params
func (o *CreateClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster params
func (o *CreateClusterParams) WithHTTPClient(client *http.Client) *CreateClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster params
func (o *CreateClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create cluster params
func (o *CreateClusterParams) WithXRequestID(xRequestID *string) *CreateClusterParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create cluster params
func (o *CreateClusterParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithClusterConfig adds the clusterConfig to the create cluster params
func (o *CreateClusterParams) WithClusterConfig(clusterConfig *models.ClusterConfig) *CreateClusterParams {
	o.SetClusterConfig(clusterConfig)
	return o
}

// SetClusterConfig adds the clusterConfig to the create cluster params
func (o *CreateClusterParams) SetClusterConfig(clusterConfig *models.ClusterConfig) {
	o.ClusterConfig = clusterConfig
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if o.ClusterConfig == nil {
		o.ClusterConfig = new(models.ClusterConfig)
	}

	if err := r.SetBodyParam(o.ClusterConfig); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
