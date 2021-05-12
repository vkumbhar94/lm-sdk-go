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
	"github.com/go-openapi/validate"
)

// NetflowGroupWidget netflow group widget
// swagger:model NetflowGroupWidget
type NetflowGroupWidget struct {
	dashboardIdField *int32

	descriptionField string

	idField int32

	intervalField int32

	lastUpdatedByField string

	lastUpdatedOnField int64

	nameField *string

	themeField string

	timescaleField string

	userPermissionField string

	// data type
	DataType string `json:"dataType,omitempty"`

	// device group Id
	// Required: true
	DeviceGroupID *int32 `json:"deviceGroupId"`

	// device group name
	// Read Only: true
	DeviceGroupName string `json:"deviceGroupName,omitempty"`

	// qos type
	QosType string `json:"qosType,omitempty"`

	// row filters
	RowFilters string `json:"rowFilters,omitempty"`
}

// DashboardID gets the dashboard Id of this subtype
func (m *NetflowGroupWidget) DashboardID() *int32 {
	return m.dashboardIdField
}

// SetDashboardID sets the dashboard Id of this subtype
func (m *NetflowGroupWidget) SetDashboardID(val *int32) {
	m.dashboardIdField = val
}

// Description gets the description of this subtype
func (m *NetflowGroupWidget) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *NetflowGroupWidget) SetDescription(val string) {
	m.descriptionField = val
}

// ID gets the id of this subtype
func (m *NetflowGroupWidget) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *NetflowGroupWidget) SetID(val int32) {
	m.idField = val
}

// Interval gets the interval of this subtype
func (m *NetflowGroupWidget) Interval() int32 {
	return m.intervalField
}

// SetInterval sets the interval of this subtype
func (m *NetflowGroupWidget) SetInterval(val int32) {
	m.intervalField = val
}

// LastUpdatedBy gets the last updated by of this subtype
func (m *NetflowGroupWidget) LastUpdatedBy() string {
	return m.lastUpdatedByField
}

// SetLastUpdatedBy sets the last updated by of this subtype
func (m *NetflowGroupWidget) SetLastUpdatedBy(val string) {
	m.lastUpdatedByField = val
}

// LastUpdatedOn gets the last updated on of this subtype
func (m *NetflowGroupWidget) LastUpdatedOn() int64 {
	return m.lastUpdatedOnField
}

// SetLastUpdatedOn sets the last updated on of this subtype
func (m *NetflowGroupWidget) SetLastUpdatedOn(val int64) {
	m.lastUpdatedOnField = val
}

// Name gets the name of this subtype
func (m *NetflowGroupWidget) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *NetflowGroupWidget) SetName(val *string) {
	m.nameField = val
}

// Theme gets the theme of this subtype
func (m *NetflowGroupWidget) Theme() string {
	return m.themeField
}

// SetTheme sets the theme of this subtype
func (m *NetflowGroupWidget) SetTheme(val string) {
	m.themeField = val
}

// Timescale gets the timescale of this subtype
func (m *NetflowGroupWidget) Timescale() string {
	return m.timescaleField
}

// SetTimescale sets the timescale of this subtype
func (m *NetflowGroupWidget) SetTimescale(val string) {
	m.timescaleField = val
}

// Type gets the type of this subtype
func (m *NetflowGroupWidget) Type() string {
	return "groupNetflow"
}

// SetType sets the type of this subtype
func (m *NetflowGroupWidget) SetType(val string) {

}

// UserPermission gets the user permission of this subtype
func (m *NetflowGroupWidget) UserPermission() string {
	return m.userPermissionField
}

// SetUserPermission sets the user permission of this subtype
func (m *NetflowGroupWidget) SetUserPermission(val string) {
	m.userPermissionField = val
}

// DataType gets the data type of this subtype

// DeviceGroupID gets the device group Id of this subtype

// DeviceGroupName gets the device group name of this subtype

// QosType gets the qos type of this subtype

// RowFilters gets the row filters of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *NetflowGroupWidget) UnmarshalJSON(raw []byte) error {
	var data struct {

		// data type
		DataType string `json:"dataType,omitempty"`

		// device group Id
		// Required: true
		DeviceGroupID *int32 `json:"deviceGroupId"`

		// device group name
		// Read Only: true
		DeviceGroupName string `json:"deviceGroupName,omitempty"`

		// qos type
		QosType string `json:"qosType,omitempty"`

		// row filters
		RowFilters string `json:"rowFilters,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		DashboardID *int32 `json:"dashboardId"`

		Description string `json:"description,omitempty"`

		ID int32 `json:"id,omitempty"`

		Interval int32 `json:"interval,omitempty"`

		LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

		LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

		Name *string `json:"name"`

		Theme string `json:"theme,omitempty"`

		Timescale string `json:"timescale,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result NetflowGroupWidget

	result.dashboardIdField = base.DashboardID

	result.descriptionField = base.Description

	result.idField = base.ID

	result.intervalField = base.Interval

	result.lastUpdatedByField = base.LastUpdatedBy

	result.lastUpdatedOnField = base.LastUpdatedOn

	result.nameField = base.Name

	result.themeField = base.Theme

	result.timescaleField = base.Timescale

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.userPermissionField = base.UserPermission

	result.DataType = data.DataType

	result.DeviceGroupID = data.DeviceGroupID

	result.DeviceGroupName = data.DeviceGroupName

	result.QosType = data.QosType

	result.RowFilters = data.RowFilters

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m NetflowGroupWidget) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// data type
		DataType string `json:"dataType,omitempty"`

		// device group Id
		// Required: true
		DeviceGroupID *int32 `json:"deviceGroupId"`

		// device group name
		// Read Only: true
		DeviceGroupName string `json:"deviceGroupName,omitempty"`

		// qos type
		QosType string `json:"qosType,omitempty"`

		// row filters
		RowFilters string `json:"rowFilters,omitempty"`
	}{

		DataType: m.DataType,

		DeviceGroupID: m.DeviceGroupID,

		DeviceGroupName: m.DeviceGroupName,

		QosType: m.QosType,

		RowFilters: m.RowFilters,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		DashboardID *int32 `json:"dashboardId"`

		Description string `json:"description,omitempty"`

		ID int32 `json:"id,omitempty"`

		Interval int32 `json:"interval,omitempty"`

		LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

		LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

		Name *string `json:"name"`

		Theme string `json:"theme,omitempty"`

		Timescale string `json:"timescale,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}{

		DashboardID: m.DashboardID(),

		Description: m.Description(),

		ID: m.ID(),

		Interval: m.Interval(),

		LastUpdatedBy: m.LastUpdatedBy(),

		LastUpdatedOn: m.LastUpdatedOn(),

		Name: m.Name(),

		Theme: m.Theme(),

		Timescale: m.Timescale(),

		Type: m.Type(),

		UserPermission: m.UserPermission(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this netflow group widget
func (m *NetflowGroupWidget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboardID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceGroupID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetflowGroupWidget) validateDashboardID(formats strfmt.Registry) error {

	if err := validate.Required("dashboardId", "body", m.DashboardID()); err != nil {
		return err
	}

	return nil
}

func (m *NetflowGroupWidget) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *NetflowGroupWidget) validateDeviceGroupID(formats strfmt.Registry) error {

	if err := validate.Required("deviceGroupId", "body", m.DeviceGroupID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetflowGroupWidget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetflowGroupWidget) UnmarshalBinary(b []byte) error {
	var res NetflowGroupWidget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
