package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/supervised-io/kov/gen/models"
)

// NewUpdateClusterParams creates a new UpdateClusterParams object
// with the default values initialized.
func NewUpdateClusterParams() UpdateClusterParams {
	var ()
	return UpdateClusterParams{}
}

// UpdateClusterParams contains all the bound params for the update cluster operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateCluster
type UpdateClusterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*A unique UUID for the request
	  Min Length: 1
	  In: header
	*/
	XRequestID *string
	/*the new config of the cluster to be updated
	  Required: true
	  In: body
	*/
	ClusterUpdateConfig *models.ClusterUpdateConfig
	/*the cluster name to be updated
	  Required: true
	  In: path
	*/
	Name string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateClusterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if err := o.bindXRequestID(r.Header[http.CanonicalHeaderKey("X-Request-Id")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ClusterUpdateConfig
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("clusterUpdateConfig", "body"))
			} else {
				res = append(res, errors.NewParseError("clusterUpdateConfig", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ClusterUpdateConfig = &body
			}
		}

	} else {
		res = append(res, errors.Required("clusterUpdateConfig", "body"))
	}

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateClusterParams) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.XRequestID = &raw

	if err := o.validateXRequestID(formats); err != nil {
		return err
	}

	return nil
}

func (o *UpdateClusterParams) validateXRequestID(formats strfmt.Registry) error {

	if err := validate.MinLength("X-Request-Id", "header", (*o.XRequestID), 1); err != nil {
		return err
	}

	return nil
}

func (o *UpdateClusterParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Name = raw

	return nil
}
