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

// EventsReader is a Reader for the Events structure.
type EventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEventsOK creates a EventsOK with default headers values
func NewEventsOK() *EventsOK {
	return &EventsOK{}
}

/*EventsOK handles this case with default header values.

Upon successful indexing of event.
*/
type EventsOK struct {
	Payload *models.EventsReturn
}

func (o *EventsOK) Error() string {
	return fmt.Sprintf("[POST /events][%d] eventsOK  %+v", 200, o.Payload)
}

func (o *EventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventsReturn)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEventsDefault creates a EventsDefault with default headers values
func NewEventsDefault(code int) *EventsDefault {
	return &EventsDefault{
		_statusCode: code,
	}
}

/*EventsDefault handles this case with default header values.

Unexpected error
*/
type EventsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the events default response
func (o *EventsDefault) Code() int {
	return o._statusCode
}

func (o *EventsDefault) Error() string {
	return fmt.Sprintf("[POST /events][%d] events default  %+v", o._statusCode, o.Payload)
}

func (o *EventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
