package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Polarishq/logface-sdk-go/models"
)

// PostAccountAPIKeysReader is a Reader for the PostAccountAPIKeys structure.
type PostAccountAPIKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccountAPIKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostAccountAPIKeysCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostAccountAPIKeysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAccountAPIKeysCreated creates a PostAccountAPIKeysCreated with default headers values
func NewPostAccountAPIKeysCreated() *PostAccountAPIKeysCreated {
	return &PostAccountAPIKeysCreated{}
}

/*PostAccountAPIKeysCreated handles this case with default header values.

API Key Created
*/
type PostAccountAPIKeysCreated struct {
	Payload *models.APIKeyResponse
}

func (o *PostAccountAPIKeysCreated) Error() string {
	return fmt.Sprintf("[POST /account/apiKeys][%d] postAccountApiKeysCreated  %+v", 201, o.Payload)
}

func (o *PostAccountAPIKeysCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountAPIKeysDefault creates a PostAccountAPIKeysDefault with default headers values
func NewPostAccountAPIKeysDefault(code int) *PostAccountAPIKeysDefault {
	return &PostAccountAPIKeysDefault{
		_statusCode: code,
	}
}

/*PostAccountAPIKeysDefault handles this case with default header values.

Unexpected error
*/
type PostAccountAPIKeysDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post account API keys default response
func (o *PostAccountAPIKeysDefault) Code() int {
	return o._statusCode
}

func (o *PostAccountAPIKeysDefault) Error() string {
	return fmt.Sprintf("[POST /account/apiKeys][%d] PostAccountAPIKeys default  %+v", o._statusCode, o.Payload)
}

func (o *PostAccountAPIKeysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostAccountAPIKeysBody post account API keys body
swagger:model PostAccountAPIKeysBody
*/
type PostAccountAPIKeysBody struct {

	// plain name of api key
	Label string `json:"label,omitempty"`
}
