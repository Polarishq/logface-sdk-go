package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ApplicationRequest application request
// swagger:model ApplicationRequest
type ApplicationRequest struct {

	// app url
	// Required: true
	AppURL *string `json:"app_url"`

	// description
	// Required: true
	Description *string `json:"description"`

	// image url
	// Required: true
	ImageURL *string `json:"image_url"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this application request
func (m *ApplicationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImageURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationRequest) validateAppURL(formats strfmt.Registry) error {

	if err := validate.Required("app_url", "body", m.AppURL); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationRequest) validateImageURL(formats strfmt.Registry) error {

	if err := validate.Required("image_url", "body", m.ImageURL); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
