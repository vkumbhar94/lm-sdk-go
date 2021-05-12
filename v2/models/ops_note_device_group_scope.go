// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpsNoteDeviceGroupScope ops note device group scope
// swagger:model OpsNoteDeviceGroupScope
type OpsNoteDeviceGroupScope struct {

	// full path
	FullPath string `json:"fullPath,omitempty"`

	// group Id
	GroupID int32 `json:"groupId,omitempty"`
}

// Type gets the type of this subtype
func (m *OpsNoteDeviceGroupScope) Type() string {
	return "deviceGroup"
}

// SetType sets the type of this subtype
func (m *OpsNoteDeviceGroupScope) SetType(val string) {

}

// FullPath gets the full path of this subtype

// GroupID gets the group Id of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *OpsNoteDeviceGroupScope) UnmarshalJSON(raw []byte) error {
	var data struct {

		// full path
		FullPath string `json:"fullPath,omitempty"`

		// group Id
		GroupID int32 `json:"groupId,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result OpsNoteDeviceGroupScope

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.FullPath = data.FullPath

	result.GroupID = data.GroupID

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m OpsNoteDeviceGroupScope) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// full path
		FullPath string `json:"fullPath,omitempty"`

		// group Id
		GroupID int32 `json:"groupId,omitempty"`
	}{

		FullPath: m.FullPath,

		GroupID: m.GroupID,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`
	}{

		Type: m.Type(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this ops note device group scope
func (m *OpsNoteDeviceGroupScope) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpsNoteDeviceGroupScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpsNoteDeviceGroupScope) UnmarshalBinary(b []byte) error {
	var res OpsNoteDeviceGroupScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
