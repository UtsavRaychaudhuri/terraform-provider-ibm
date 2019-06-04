// Code generated by go-swagger; DO NOT EDIT.

package geography

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRegionsRegionNameZonesParams creates a new GetRegionsRegionNameZonesParams object
// with the default values initialized.
func NewGetRegionsRegionNameZonesParams() *GetRegionsRegionNameZonesParams {
	var ()
	return &GetRegionsRegionNameZonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegionsRegionNameZonesParamsWithTimeout creates a new GetRegionsRegionNameZonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRegionsRegionNameZonesParamsWithTimeout(timeout time.Duration) *GetRegionsRegionNameZonesParams {
	var ()
	return &GetRegionsRegionNameZonesParams{

		timeout: timeout,
	}
}

// NewGetRegionsRegionNameZonesParamsWithContext creates a new GetRegionsRegionNameZonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRegionsRegionNameZonesParamsWithContext(ctx context.Context) *GetRegionsRegionNameZonesParams {
	var ()
	return &GetRegionsRegionNameZonesParams{

		Context: ctx,
	}
}

// NewGetRegionsRegionNameZonesParamsWithHTTPClient creates a new GetRegionsRegionNameZonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRegionsRegionNameZonesParamsWithHTTPClient(client *http.Client) *GetRegionsRegionNameZonesParams {
	var ()
	return &GetRegionsRegionNameZonesParams{
		HTTPClient: client,
	}
}

/*GetRegionsRegionNameZonesParams contains all the parameters to send to the API endpoint
for the get regions region name zones operation typically these are written to a http.Request
*/
type GetRegionsRegionNameZonesParams struct {

	/*RegionName
	  The region name

	*/
	RegionName string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get regions region name zones params
func (o *GetRegionsRegionNameZonesParams) WithTimeout(timeout time.Duration) *GetRegionsRegionNameZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get regions region name zones params
func (o *GetRegionsRegionNameZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get regions region name zones params
func (o *GetRegionsRegionNameZonesParams) WithContext(ctx context.Context) *GetRegionsRegionNameZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get regions region name zones params
func (o *GetRegionsRegionNameZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get regions region name zones params
func (o *GetRegionsRegionNameZonesParams) WithHTTPClient(client *http.Client) *GetRegionsRegionNameZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get regions region name zones params
func (o *GetRegionsRegionNameZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegionName adds the regionName to the get regions region name zones params
func (o *GetRegionsRegionNameZonesParams) WithRegionName(regionName string) *GetRegionsRegionNameZonesParams {
	o.SetRegionName(regionName)
	return o
}

// SetRegionName adds the regionName to the get regions region name zones params
func (o *GetRegionsRegionNameZonesParams) SetRegionName(regionName string) {
	o.RegionName = regionName
}

// WithVersion adds the version to the get regions region name zones params
func (o *GetRegionsRegionNameZonesParams) WithVersion(version string) *GetRegionsRegionNameZonesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get regions region name zones params
func (o *GetRegionsRegionNameZonesParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegionsRegionNameZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param region_name
	if err := r.SetPathParam("region_name", o.RegionName); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
