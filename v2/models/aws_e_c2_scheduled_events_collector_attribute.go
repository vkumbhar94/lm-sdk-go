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

// AwsEC2ScheduledEventsCollectorAttribute aws e c2 scheduled events collector attribute
// swagger:model AwsEC2ScheduledEventsCollectorAttribute
type AwsEC2ScheduledEventsCollectorAttribute struct {

	// period
	Period int32 `json:"period,omitempty"`
}

// Name gets the name of this subtype
func (m *AwsEC2ScheduledEventsCollectorAttribute) Name() string {
	return "awsec2scheduledevents"
}

// SetName sets the name of this subtype
func (m *AwsEC2ScheduledEventsCollectorAttribute) SetName(val string) {

}

// Period gets the period of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AwsEC2ScheduledEventsCollectorAttribute) UnmarshalJSON(raw []byte) error {
	var data struct {

		// period
		Period int32 `json:"period,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Name string `json:"name"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result AwsEC2ScheduledEventsCollectorAttribute

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.Period = data.Period

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AwsEC2ScheduledEventsCollectorAttribute) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// period
		Period int32 `json:"period,omitempty"`
	}{

		Period: m.Period,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Name string `json:"name"`
	}{

		Name: m.Name(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this aws e c2 scheduled events collector attribute
func (m *AwsEC2ScheduledEventsCollectorAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AwsEC2ScheduledEventsCollectorAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsEC2ScheduledEventsCollectorAttribute) UnmarshalBinary(b []byte) error {
	var res AwsEC2ScheduledEventsCollectorAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
