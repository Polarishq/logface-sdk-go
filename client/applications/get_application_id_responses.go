package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Polarishq/api-client-go/models"
)

// GetApplicationIDReader is a Reader for the GetApplicationID structure.
type GetApplicationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetApplicationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetApplicationIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplicationIDOK creates a GetApplicationIDOK with default headers values
func NewGetApplicationIDOK() *GetApplicationIDOK {
	return &GetApplicationIDOK{}
}

/*GetApplicationIDOK handles this case with default header values.

successful operation
*/
type GetApplicationIDOK struct {
	Payload *models.Application
}

func (o *GetApplicationIDOK) Error() string {
	return fmt.Sprintf("[GET /application/{id}][%d] getApplicationIdOK  %+v", 200, o.Payload)
}

func (o *GetApplicationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationIDDefault creates a GetApplicationIDDefault with default headers values
func NewGetApplicationIDDefault(code int) *GetApplicationIDDefault {
	return &GetApplicationIDDefault{
		_statusCode: code,
	}
}

/*GetApplicationIDDefault handles this case with default header values.

Unexpected error
*/
type GetApplicationIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get application ID default response
func (o *GetApplicationIDDefault) Code() int {
	return o._statusCode
}

func (o *GetApplicationIDDefault) Error() string {
	return fmt.Sprintf("[GET /application/{id}][%d] GetApplicationID default  %+v", o._statusCode, o.Payload)
}

func (o *GetApplicationIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}