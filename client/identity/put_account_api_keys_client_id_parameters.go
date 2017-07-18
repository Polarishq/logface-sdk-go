package identity

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

// NewPutAccountAPIKeysClientIDParams creates a new PutAccountAPIKeysClientIDParams object
// with the default values initialized.
func NewPutAccountAPIKeysClientIDParams() *PutAccountAPIKeysClientIDParams {
	var ()
	return &PutAccountAPIKeysClientIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAccountAPIKeysClientIDParamsWithTimeout creates a new PutAccountAPIKeysClientIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAccountAPIKeysClientIDParamsWithTimeout(timeout time.Duration) *PutAccountAPIKeysClientIDParams {
	var ()
	return &PutAccountAPIKeysClientIDParams{

		timeout: timeout,
	}
}

// NewPutAccountAPIKeysClientIDParamsWithContext creates a new PutAccountAPIKeysClientIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAccountAPIKeysClientIDParamsWithContext(ctx context.Context) *PutAccountAPIKeysClientIDParams {
	var ()
	return &PutAccountAPIKeysClientIDParams{

		Context: ctx,
	}
}

// NewPutAccountAPIKeysClientIDParamsWithHTTPClient creates a new PutAccountAPIKeysClientIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAccountAPIKeysClientIDParamsWithHTTPClient(client *http.Client) *PutAccountAPIKeysClientIDParams {
	var ()
	return &PutAccountAPIKeysClientIDParams{
		HTTPClient: client,
	}
}

/*PutAccountAPIKeysClientIDParams contains all the parameters to send to the API endpoint
for the put account API keys client ID operation typically these are written to a http.Request
*/
type PutAccountAPIKeysClientIDParams struct {

	/*Body*/
	Body PutAccountAPIKeysClientIDBody
	/*ClientID
	  the client ID of the API to update

	*/
	ClientID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put account API keys client ID params
func (o *PutAccountAPIKeysClientIDParams) WithTimeout(timeout time.Duration) *PutAccountAPIKeysClientIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put account API keys client ID params
func (o *PutAccountAPIKeysClientIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put account API keys client ID params
func (o *PutAccountAPIKeysClientIDParams) WithContext(ctx context.Context) *PutAccountAPIKeysClientIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put account API keys client ID params
func (o *PutAccountAPIKeysClientIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put account API keys client ID params
func (o *PutAccountAPIKeysClientIDParams) WithHTTPClient(client *http.Client) *PutAccountAPIKeysClientIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put account API keys client ID params
func (o *PutAccountAPIKeysClientIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put account API keys client ID params
func (o *PutAccountAPIKeysClientIDParams) WithBody(body PutAccountAPIKeysClientIDBody) *PutAccountAPIKeysClientIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put account API keys client ID params
func (o *PutAccountAPIKeysClientIDParams) SetBody(body PutAccountAPIKeysClientIDBody) {
	o.Body = body
}

// WithClientID adds the clientID to the put account API keys client ID params
func (o *PutAccountAPIKeysClientIDParams) WithClientID(clientID string) *PutAccountAPIKeysClientIDParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the put account API keys client ID params
func (o *PutAccountAPIKeysClientIDParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WriteToRequest writes these params to a swagger request
func (o *PutAccountAPIKeysClientIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param client_id
	if err := r.SetPathParam("client_id", o.ClientID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
