package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Polarishq/api-client-go/models"
)

// GetEventsReader is a Reader for the GetEvents structure.
type GetEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEventsOK creates a GetEventsOK with default headers values
func NewGetEventsOK() *GetEventsOK {
	return &GetEventsOK{}
}

/*GetEventsOK handles this case with default header values.

List of events that matched
*/
type GetEventsOK struct {
	Payload *models.QueryResponse
}

func (o *GetEventsOK) Error() string {
	return fmt.Sprintf("[GET /events][%d] getEventsOK  %+v", 200, o.Payload)
}

func (o *GetEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsDefault creates a GetEventsDefault with default headers values
func NewGetEventsDefault(code int) *GetEventsDefault {
	return &GetEventsDefault{
		_statusCode: code,
	}
}

/*GetEventsDefault handles this case with default header values.

Unexpected error
*/
type GetEventsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get events default response
func (o *GetEventsDefault) Code() int {
	return o._statusCode
}

func (o *GetEventsDefault) Error() string {
	return fmt.Sprintf("[GET /events][%d] GetEvents default  %+v", o._statusCode, o.Payload)
}

func (o *GetEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
