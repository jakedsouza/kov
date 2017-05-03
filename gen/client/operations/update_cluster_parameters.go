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

// NewUpdateClusterParams creates a new UpdateClusterParams object
// with the default values initialized.
func NewUpdateClusterParams() *UpdateClusterParams {
	var ()
	return &UpdateClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateClusterParamsWithTimeout creates a new UpdateClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateClusterParamsWithTimeout(timeout time.Duration) *UpdateClusterParams {
	var ()
	return &UpdateClusterParams{

		timeout: timeout,
	}
}

// NewUpdateClusterParamsWithContext creates a new UpdateClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateClusterParamsWithContext(ctx context.Context) *UpdateClusterParams {
	var ()
	return &UpdateClusterParams{

		Context: ctx,
	}
}

// NewUpdateClusterParamsWithHTTPClient creates a new UpdateClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateClusterParamsWithHTTPClient(client *http.Client) *UpdateClusterParams {
	var ()
	return &UpdateClusterParams{
		HTTPClient: client,
	}
}

/*UpdateClusterParams contains all the parameters to send to the API endpoint
for the update cluster operation typically these are written to a http.Request
*/
type UpdateClusterParams struct {

	/*XRequestID
	  A unique UUID for the request

	*/
	XRequestID *string
	/*ClusterUpdateConfig
	  the new config of the cluster to be updated

	*/
	ClusterUpdateConfig *models.ClusterUpdateConfig
	/*Name
	  the cluster name to be updated

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update cluster params
func (o *UpdateClusterParams) WithTimeout(timeout time.Duration) *UpdateClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cluster params
func (o *UpdateClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cluster params
func (o *UpdateClusterParams) WithContext(ctx context.Context) *UpdateClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cluster params
func (o *UpdateClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cluster params
func (o *UpdateClusterParams) WithHTTPClient(client *http.Client) *UpdateClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cluster params
func (o *UpdateClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update cluster params
func (o *UpdateClusterParams) WithXRequestID(xRequestID *string) *UpdateClusterParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update cluster params
func (o *UpdateClusterParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithClusterUpdateConfig adds the clusterUpdateConfig to the update cluster params
func (o *UpdateClusterParams) WithClusterUpdateConfig(clusterUpdateConfig *models.ClusterUpdateConfig) *UpdateClusterParams {
	o.SetClusterUpdateConfig(clusterUpdateConfig)
	return o
}

// SetClusterUpdateConfig adds the clusterUpdateConfig to the update cluster params
func (o *UpdateClusterParams) SetClusterUpdateConfig(clusterUpdateConfig *models.ClusterUpdateConfig) {
	o.ClusterUpdateConfig = clusterUpdateConfig
}

// WithName adds the name to the update cluster params
func (o *UpdateClusterParams) WithName(name string) *UpdateClusterParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update cluster params
func (o *UpdateClusterParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ClusterUpdateConfig == nil {
		o.ClusterUpdateConfig = new(models.ClusterUpdateConfig)
	}

	if err := r.SetBodyParam(o.ClusterUpdateConfig); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
