package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Event Event payload with metadata.
// swagger:model Event
type Event struct {

	// An identifier that specified where this event originated from. Typically this is a hostname, IP address, or an alias of the host that generated this event.
	//
	Entity string `json:"entity,omitempty"`

	// This identifier helps to classify events of this type based of what generated them.
	//
	Source string `json:"source,omitempty"`

	// Epoch-formatted time of event generation.
	//
	Time string `json:"time,omitempty"`

	// event
	Event map[string]string `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *Event) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// An identifier that specified where this event originated from. Typically this is a hostname, IP address, or an alias of the host that generated this event.
		//
		Entity string `json:"entity,omitempty"`

		// This identifier helps to classify events of this type based of what generated them.
		//
		Source string `json:"source,omitempty"`

		// Epoch-formatted time of event generation.
		//
		Time string `json:"time,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv Event

	rcv.Entity = stage1.Entity

	rcv.Source = stage1.Source

	rcv.Time = stage1.Time

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "entity")

	delete(stage2, "source")

	delete(stage2, "time")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]string)
		for k, v := range stage2 {
			var toadd string
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.Event = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m Event) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// An identifier that specified where this event originated from. Typically this is a hostname, IP address, or an alias of the host that generated this event.
		//
		Entity string `json:"entity,omitempty"`

		// This identifier helps to classify events of this type based of what generated them.
		//
		Source string `json:"source,omitempty"`

		// Epoch-formatted time of event generation.
		//
		Time string `json:"time,omitempty"`
	}

	stage1.Entity = m.Entity

	stage1.Source = m.Source

	stage1.Time = m.Time

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.Event) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.Event)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}