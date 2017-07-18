package login

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
)

// NewGetLoginAuthorizeParams creates a new GetLoginAuthorizeParams object
// with the default values initialized.
func NewGetLoginAuthorizeParams() *GetLoginAuthorizeParams {
	var ()
	return &GetLoginAuthorizeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoginAuthorizeParamsWithTimeout creates a new GetLoginAuthorizeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoginAuthorizeParamsWithTimeout(timeout time.Duration) *GetLoginAuthorizeParams {
	var ()
	return &GetLoginAuthorizeParams{

		timeout: timeout,
	}
}

// NewGetLoginAuthorizeParamsWithContext creates a new GetLoginAuthorizeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoginAuthorizeParamsWithContext(ctx context.Context) *GetLoginAuthorizeParams {
	var ()
	return &GetLoginAuthorizeParams{

		Context: ctx,
	}
}

// NewGetLoginAuthorizeParamsWithHTTPClient creates a new GetLoginAuthorizeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLoginAuthorizeParamsWithHTTPClient(client *http.Client) *GetLoginAuthorizeParams {
	var ()
	return &GetLoginAuthorizeParams{
		HTTPClient: client,
	}
}

/*GetLoginAuthorizeParams contains all the parameters to send to the API endpoint
for the get login authorize operation typically these are written to a http.Request
*/
type GetLoginAuthorizeParams struct {

	/*Provider*/
	Provider string
	/*RedirectURL*/
	RedirectURL *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get login authorize params
func (o *GetLoginAuthorizeParams) WithTimeout(timeout time.Duration) *GetLoginAuthorizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get login authorize params
func (o *GetLoginAuthorizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get login authorize params
func (o *GetLoginAuthorizeParams) WithContext(ctx context.Context) *GetLoginAuthorizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get login authorize params
func (o *GetLoginAuthorizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get login authorize params
func (o *GetLoginAuthorizeParams) WithHTTPClient(client *http.Client) *GetLoginAuthorizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get login authorize params
func (o *GetLoginAuthorizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProvider adds the provider to the get login authorize params
func (o *GetLoginAuthorizeParams) WithProvider(provider string) *GetLoginAuthorizeParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the get login authorize params
func (o *GetLoginAuthorizeParams) SetProvider(provider string) {
	o.Provider = provider
}

// WithRedirectURL adds the redirectURL to the get login authorize params
func (o *GetLoginAuthorizeParams) WithRedirectURL(redirectURL *string) *GetLoginAuthorizeParams {
	o.SetRedirectURL(redirectURL)
	return o
}

// SetRedirectURL adds the redirectUrl to the get login authorize params
func (o *GetLoginAuthorizeParams) SetRedirectURL(redirectURL *string) {
	o.RedirectURL = redirectURL
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoginAuthorizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param provider
	qrProvider := o.Provider
	qProvider := qrProvider
	if qProvider != "" {
		if err := r.SetQueryParam("provider", qProvider); err != nil {
			return err
		}
	}

	if o.RedirectURL != nil {

		// query param redirect_url
		var qrRedirectURL string
		if o.RedirectURL != nil {
			qrRedirectURL = *o.RedirectURL
		}
		qRedirectURL := qrRedirectURL
		if qRedirectURL != "" {
			if err := r.SetQueryParam("redirect_url", qRedirectURL); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
