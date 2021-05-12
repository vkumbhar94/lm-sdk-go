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

// OpsNoteGroupAllScope ops note group all scope
// swagger:model OpsNoteGroupAllScope
type OpsNoteGroupAllScope struct {

	// group Id
	GroupID int32 `json:"groupId,omitempty"`
}

// Type gets the type of this subtype
func (m *OpsNoteGroupAllScope) Type() string {
	return "groupAll"
}

// SetType sets the type of this subtype
func (m *OpsNoteGroupAllScope) SetType(val string) {

}

// GroupID gets the group Id of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *OpsNoteGroupAllScope) UnmarshalJSON(raw []byte) error {
	var data struct {

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

	var result OpsNoteGroupAllScope

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.GroupID = data.GroupID

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m OpsNoteGroupAllScope) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// group Id
		GroupID int32 `json:"groupId,omitempty"`
	}{

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

// Validate validates this ops note group all scope
func (m *OpsNoteGroupAllScope) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpsNoteGroupAllScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpsNoteGroupAllScope) UnmarshalBinary(b []byte) error {
	var res OpsNoteGroupAllScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
