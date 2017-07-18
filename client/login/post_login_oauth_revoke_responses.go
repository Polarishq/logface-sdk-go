package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Polarishq/api-client-go/models"
)

// PostLoginOauthRevokeReader is a Reader for the PostLoginOauthRevoke structure.
type PostLoginOauthRevokeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLoginOauthRevokeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostLoginOauthRevokeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostLoginOauthRevokeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostLoginOauthRevokeOK creates a PostLoginOauthRevokeOK with default headers values
func NewPostLoginOauthRevokeOK() *PostLoginOauthRevokeOK {
	return &PostLoginOauthRevokeOK{}
}

/*PostLoginOauthRevokeOK handles this case with default header values.

Token Revoked
*/
type PostLoginOauthRevokeOK struct {
}

func (o *PostLoginOauthRevokeOK) Error() string {
	return fmt.Sprintf("[POST /login/oauth/revoke][%d] postLoginOauthRevokeOK ", 200)
}

func (o *PostLoginOauthRevokeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLoginOauthRevokeDefault creates a PostLoginOauthRevokeDefault with default headers values
func NewPostLoginOauthRevokeDefault(code int) *PostLoginOauthRevokeDefault {
	return &PostLoginOauthRevokeDefault{
		_statusCode: code,
	}
}

/*PostLoginOauthRevokeDefault handles this case with default header values.

Unexpected error
*/
type PostLoginOauthRevokeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post login oauth revoke default response
func (o *PostLoginOauthRevokeDefault) Code() int {
	return o._statusCode
}

func (o *PostLoginOauthRevokeDefault) Error() string {
	return fmt.Sprintf("[POST /login/oauth/revoke][%d] PostLoginOauthRevoke default  %+v", o._statusCode, o.Payload)
}

func (o *PostLoginOauthRevokeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLoginOauthRevokeBody post login oauth revoke body
swagger:model PostLoginOauthRevokeBody
*/
type PostLoginOauthRevokeBody struct {

	// refresh token
	// Required: true
	RefreshToken *string `json:"refresh_token"`
}

// MarshalBinary interface implementation
func (o *PostLoginOauthRevokeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLoginOauthRevokeBody) UnmarshalBinary(b []byte) error {
	var res PostLoginOauthRevokeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
