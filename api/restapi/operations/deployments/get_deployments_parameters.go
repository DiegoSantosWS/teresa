package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDeploymentsParams creates a new GetDeploymentsParams object
// with the default values initialized.
func NewGetDeploymentsParams() GetDeploymentsParams {
	var (
		limitDefault int64 = int64(20)
		sinceDefault int64 = int64(0)
	)
	return GetDeploymentsParams{
		Limit: &limitDefault,

		Since: &sinceDefault,
	}
}

// GetDeploymentsParams contains all the bound params for the get deployments operation
// typically these are obtained from a http.Request
//
// swagger:parameters getDeployments
type GetDeploymentsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*App ID
	  Required: true
	  In: path
	*/
	AppID string
	/*Limit
	  In: query
	  Default: 20
	*/
	Limit *int64
	/*Offset
	  In: query
	  Default: 0
	*/
	Since *int64
	/*Team ID
	  Required: true
	  In: path
	*/
	TeamID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetDeploymentsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rAppID, rhkAppID, _ := route.Params.GetOK("app_id")
	if err := o.bindAppID(rAppID, rhkAppID, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qSince, qhkSince, _ := qs.GetOK("since")
	if err := o.bindSince(qSince, qhkSince, route.Formats); err != nil {
		res = append(res, err)
	}

	rTeamID, rhkTeamID, _ := route.Params.GetOK("team_id")
	if err := o.bindTeamID(rTeamID, rhkTeamID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDeploymentsParams) bindAppID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.AppID = raw

	return nil
}

func (o *GetDeploymentsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var limitDefault int64 = int64(20)
		o.Limit = &limitDefault
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

func (o *GetDeploymentsParams) bindSince(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var sinceDefault int64 = int64(0)
		o.Since = &sinceDefault
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("since", "query", "int64", raw)
	}
	o.Since = &value

	return nil
}

func (o *GetDeploymentsParams) bindTeamID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.TeamID = raw

	return nil
}