// Code generated by go-swagger; DO NOT EDIT.

package context

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

	"github.com/Polarishq/logface-sdk-go/models"
)

// NewPostContextsParams creates a new PostContextsParams object
// with the default values initialized.
func NewPostContextsParams() *PostContextsParams {
	var ()
	return &PostContextsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostContextsParamsWithTimeout creates a new PostContextsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostContextsParamsWithTimeout(timeout time.Duration) *PostContextsParams {
	var ()
	return &PostContextsParams{

		timeout: timeout,
	}
}

// NewPostContextsParamsWithContext creates a new PostContextsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostContextsParamsWithContext(ctx context.Context) *PostContextsParams {
	var ()
	return &PostContextsParams{

		Context: ctx,
	}
}

// NewPostContextsParamsWithHTTPClient creates a new PostContextsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostContextsParamsWithHTTPClient(client *http.Client) *PostContextsParams {
	var ()
	return &PostContextsParams{
		HTTPClient: client,
	}
}

/*PostContextsParams contains all the parameters to send to the API endpoint
for the post contexts operation typically these are written to a http.Request
*/
type PostContextsParams struct {

	/*SearchContext*/
	SearchContext *models.SearchContext
	/*TenantID
	  ID of the tenant

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post contexts params
func (o *PostContextsParams) WithTimeout(timeout time.Duration) *PostContextsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post contexts params
func (o *PostContextsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post contexts params
func (o *PostContextsParams) WithContext(ctx context.Context) *PostContextsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post contexts params
func (o *PostContextsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post contexts params
func (o *PostContextsParams) WithHTTPClient(client *http.Client) *PostContextsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post contexts params
func (o *PostContextsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearchContext adds the searchContext to the post contexts params
func (o *PostContextsParams) WithSearchContext(searchContext *models.SearchContext) *PostContextsParams {
	o.SetSearchContext(searchContext)
	return o
}

// SetSearchContext adds the searchContext to the post contexts params
func (o *PostContextsParams) SetSearchContext(searchContext *models.SearchContext) {
	o.SearchContext = searchContext
}

// WithTenantID adds the tenantID to the post contexts params
func (o *PostContextsParams) WithTenantID(tenantID string) *PostContextsParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the post contexts params
func (o *PostContextsParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *PostContextsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SearchContext == nil {
		o.SearchContext = new(models.SearchContext)
	}

	if err := r.SetBodyParam(o.SearchContext); err != nil {
		return err
	}

	// header param tenantID
	if err := r.SetHeaderParam("tenantID", o.TenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
