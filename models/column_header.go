// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ColumnHeader column header
// swagger:model ColumnHeader
type ColumnHeader struct {

	// display type
	// Read Only: true
	DisplayType string `json:"displayType,omitempty"`

	// name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// need forecast
	// Read Only: true
	NeedForecast *bool `json:"needForecast,omitempty"`
}

// Validate validates this column header
func (m *ColumnHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ColumnHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ColumnHeader) UnmarshalBinary(b []byte) error {
	var res ColumnHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}