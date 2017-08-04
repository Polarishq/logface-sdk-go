// Code generated by go-swagger; DO NOT EDIT.

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

// NewDeleteAccountAPIKeysClientIDParams creates a new DeleteAccountAPIKeysClientIDParams object
// with the default values initialized.
func NewDeleteAccountAPIKeysClientIDParams() *DeleteAccountAPIKeysClientIDParams {
	var ()
	return &DeleteAccountAPIKeysClientIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAccountAPIKeysClientIDParamsWithTimeout creates a new DeleteAccountAPIKeysClientIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAccountAPIKeysClientIDParamsWithTimeout(timeout time.Duration) *DeleteAccountAPIKeysClientIDParams {
	var ()
	return &DeleteAccountAPIKeysClientIDParams{

		timeout: timeout,
	}
}

// NewDeleteAccountAPIKeysClientIDParamsWithContext creates a new DeleteAccountAPIKeysClientIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAccountAPIKeysClientIDParamsWithContext(ctx context.Context) *DeleteAccountAPIKeysClientIDParams {
	var ()
	return &DeleteAccountAPIKeysClientIDParams{

		Context: ctx,
	}
}

// NewDeleteAccountAPIKeysClientIDParamsWithHTTPClient creates a new DeleteAccountAPIKeysClientIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAccountAPIKeysClientIDParamsWithHTTPClient(client *http.Client) *DeleteAccountAPIKeysClientIDParams {
	var ()
	return &DeleteAccountAPIKeysClientIDParams{
		HTTPClient: client,
	}
}

/*DeleteAccountAPIKeysClientIDParams contains all the parameters to send to the API endpoint
for the delete account API keys client ID operation typically these are written to a http.Request
*/
type DeleteAccountAPIKeysClientIDParams struct {

	/*ClientID
	  the client ID of the API Key to delete

	*/
	ClientID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete account API keys client ID params
func (o *DeleteAccountAPIKeysClientIDParams) WithTimeout(timeout time.Duration) *DeleteAccountAPIKeysClientIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete account API keys client ID params
func (o *DeleteAccountAPIKeysClientIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete account API keys client ID params
func (o *DeleteAccountAPIKeysClientIDParams) WithContext(ctx context.Context) *DeleteAccountAPIKeysClientIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete account API keys client ID params
func (o *DeleteAccountAPIKeysClientIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete account API keys client ID params
func (o *DeleteAccountAPIKeysClientIDParams) WithHTTPClient(client *http.Client) *DeleteAccountAPIKeysClientIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete account API keys client ID params
func (o *DeleteAccountAPIKeysClientIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the delete account API keys client ID params
func (o *DeleteAccountAPIKeysClientIDParams) WithClientID(clientID string) *DeleteAccountAPIKeysClientIDParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the delete account API keys client ID params
func (o *DeleteAccountAPIKeysClientIDParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAccountAPIKeysClientIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param client_id
	if err := r.SetPathParam("client_id", o.ClientID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
