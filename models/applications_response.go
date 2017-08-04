package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ApplicationsResponse applications response
// swagger:model ApplicationsResponse
type ApplicationsResponse struct {

	// items
	Items []*Application `json:"items"`

	// meta data
	MetaData *ListMetaData `json:"meta_data,omitempty"`
}

// Validate validates this applications response
func (m *ApplicationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetaData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationsResponse) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {

		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {

			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationsResponse) validateMetaData(formats strfmt.Registry) error {

	if swag.IsZero(m.MetaData) { // not required
		return nil
	}

	if m.MetaData != nil {

		if err := m.MetaData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta_data")
			}
			return err
		}
	}

	return nil
}
