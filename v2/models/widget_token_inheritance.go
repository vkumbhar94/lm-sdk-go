// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// WidgetTokenInheritance widget token inheritance
// swagger:model WidgetTokenInheritance
type WidgetTokenInheritance struct {

	// fullpath
	// Read Only: true
	Fullpath string `json:"fullpath,omitempty"`

	// value
	// Read Only: true
	Value string `json:"value,omitempty"`
}

// Validate validates this widget token inheritance
func (m *WidgetTokenInheritance) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WidgetTokenInheritance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WidgetTokenInheritance) UnmarshalBinary(b []byte) error {
	var res WidgetTokenInheritance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
