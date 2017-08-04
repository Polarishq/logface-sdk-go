package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SearchContext search context
// swagger:model SearchContext
type SearchContext struct {

	// entities filter
	EntitiesFilter *SimpleFilter `json:"entities_filter,omitempty"`

	// sourcetypes filter
	SourcetypesFilter *SimpleFilter `json:"sourcetypes_filter,omitempty"`

	// time range
	TimeRange *TimeRange `json:"time_range,omitempty"`
}

// Validate validates this search context
func (m *SearchContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntitiesFilter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSourcetypesFilter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeRange(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchContext) validateEntitiesFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.EntitiesFilter) { // not required
		return nil
	}

	if m.EntitiesFilter != nil {

		if err := m.EntitiesFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entities_filter")
			}
			return err
		}
	}

	return nil
}

func (m *SearchContext) validateSourcetypesFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.SourcetypesFilter) { // not required
		return nil
	}

	if m.SourcetypesFilter != nil {

		if err := m.SourcetypesFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourcetypes_filter")
			}
			return err
		}
	}

	return nil
}

func (m *SearchContext) validateTimeRange(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeRange) { // not required
		return nil
	}

	if m.TimeRange != nil {

		if err := m.TimeRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time_range")
			}
			return err
		}
	}

	return nil
}
