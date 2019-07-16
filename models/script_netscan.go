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

// ScriptNetscan script netscan
// swagger:model ScriptNetscan
type ScriptNetscan struct {
	collectorField *int32

	collectorDescriptionField string

	collectorGroupField int32

	collectorGroupNameField string

	creatorField string

	descriptionField string

	duplicateField *ExcludeDuplicateIps

	groupField string

	idField int32

	nameField *string

	nextStartField string

	nextStartEpochField int64

	nsgIdField int32

	scheduleField *RestSchedule

	versionField int32

	// The ID of the default group to add discovered devices to
	DefaultGroup int32 `json:"defaultGroup,omitempty"`

	// The full path of the default group to add discovered devices to
	DefaultGroupFullPath string `json:"defaultGroupFullPath,omitempty"`

	// For embedded script scans, the groovy script contents
	GroovyScript string `json:"groovyScript,omitempty"`

	// For embedded script scans, the groovy script parameters
	GroovyScriptParams string `json:"groovyScriptParams,omitempty"`

	// linux script
	LinuxScript string `json:"linuxScript,omitempty"`

	// linux script params
	LinuxScriptParams string `json:"linuxScriptParams,omitempty"`

	// The parameters for an external script
	ScriptParams string `json:"scriptParams,omitempty"`

	// The script path for an external script
	ScriptPath string `json:"scriptPath,omitempty"`

	// For script scans, the type of script. Options are embeded and external
	// Required: true
	ScriptType *string `json:"scriptType"`

	// windows script
	WindowsScript string `json:"windowsScript,omitempty"`

	// windows script params
	WindowsScriptParams string `json:"windowsScriptParams,omitempty"`
}

// Collector gets the collector of this subtype
func (m *ScriptNetscan) Collector() *int32 {
	return m.collectorField
}

// SetCollector sets the collector of this subtype
func (m *ScriptNetscan) SetCollector(val *int32) {
	m.collectorField = val
}

// CollectorDescription gets the collector description of this subtype
func (m *ScriptNetscan) CollectorDescription() string {
	return m.collectorDescriptionField
}

// SetCollectorDescription sets the collector description of this subtype
func (m *ScriptNetscan) SetCollectorDescription(val string) {
	m.collectorDescriptionField = val
}

// CollectorGroup gets the collector group of this subtype
func (m *ScriptNetscan) CollectorGroup() int32 {
	return m.collectorGroupField
}

// SetCollectorGroup sets the collector group of this subtype
func (m *ScriptNetscan) SetCollectorGroup(val int32) {
	m.collectorGroupField = val
}

// CollectorGroupName gets the collector group name of this subtype
func (m *ScriptNetscan) CollectorGroupName() string {
	return m.collectorGroupNameField
}

// SetCollectorGroupName sets the collector group name of this subtype
func (m *ScriptNetscan) SetCollectorGroupName(val string) {
	m.collectorGroupNameField = val
}

// Creator gets the creator of this subtype
func (m *ScriptNetscan) Creator() string {
	return m.creatorField
}

// SetCreator sets the creator of this subtype
func (m *ScriptNetscan) SetCreator(val string) {
	m.creatorField = val
}

// Description gets the description of this subtype
func (m *ScriptNetscan) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *ScriptNetscan) SetDescription(val string) {
	m.descriptionField = val
}

// Duplicate gets the duplicate of this subtype
func (m *ScriptNetscan) Duplicate() *ExcludeDuplicateIps {
	return m.duplicateField
}

// SetDuplicate sets the duplicate of this subtype
func (m *ScriptNetscan) SetDuplicate(val *ExcludeDuplicateIps) {
	m.duplicateField = val
}

// Group gets the group of this subtype
func (m *ScriptNetscan) Group() string {
	return m.groupField
}

// SetGroup sets the group of this subtype
func (m *ScriptNetscan) SetGroup(val string) {
	m.groupField = val
}

// ID gets the id of this subtype
func (m *ScriptNetscan) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *ScriptNetscan) SetID(val int32) {
	m.idField = val
}

// Method gets the method of this subtype
func (m *ScriptNetscan) Method() string {
	return "script"
}

// SetMethod sets the method of this subtype
func (m *ScriptNetscan) SetMethod(val string) {

}

// Name gets the name of this subtype
func (m *ScriptNetscan) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *ScriptNetscan) SetName(val *string) {
	m.nameField = val
}

// NextStart gets the next start of this subtype
func (m *ScriptNetscan) NextStart() string {
	return m.nextStartField
}

// SetNextStart sets the next start of this subtype
func (m *ScriptNetscan) SetNextStart(val string) {
	m.nextStartField = val
}

// NextStartEpoch gets the next start epoch of this subtype
func (m *ScriptNetscan) NextStartEpoch() int64 {
	return m.nextStartEpochField
}

// SetNextStartEpoch sets the next start epoch of this subtype
func (m *ScriptNetscan) SetNextStartEpoch(val int64) {
	m.nextStartEpochField = val
}

// NsgID gets the nsg Id of this subtype
func (m *ScriptNetscan) NsgID() int32 {
	return m.nsgIdField
}

// SetNsgID sets the nsg Id of this subtype
func (m *ScriptNetscan) SetNsgID(val int32) {
	m.nsgIdField = val
}

// Schedule gets the schedule of this subtype
func (m *ScriptNetscan) Schedule() *RestSchedule {
	return m.scheduleField
}

// SetSchedule sets the schedule of this subtype
func (m *ScriptNetscan) SetSchedule(val *RestSchedule) {
	m.scheduleField = val
}

// Version gets the version of this subtype
func (m *ScriptNetscan) Version() int32 {
	return m.versionField
}

// SetVersion sets the version of this subtype
func (m *ScriptNetscan) SetVersion(val int32) {
	m.versionField = val
}

// DefaultGroup gets the default group of this subtype

// DefaultGroupFullPath gets the default group full path of this subtype

// GroovyScript gets the groovy script of this subtype

// GroovyScriptParams gets the groovy script params of this subtype

// LinuxScript gets the linux script of this subtype

// LinuxScriptParams gets the linux script params of this subtype

// ScriptParams gets the script params of this subtype

// ScriptPath gets the script path of this subtype

// ScriptType gets the script type of this subtype

// WindowsScript gets the windows script of this subtype

// WindowsScriptParams gets the windows script params of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ScriptNetscan) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The ID of the default group to add discovered devices to
		DefaultGroup int32 `json:"defaultGroup,omitempty"`

		// The full path of the default group to add discovered devices to
		DefaultGroupFullPath string `json:"defaultGroupFullPath,omitempty"`

		// For embedded script scans, the groovy script contents
		GroovyScript string `json:"groovyScript,omitempty"`

		// For embedded script scans, the groovy script parameters
		GroovyScriptParams string `json:"groovyScriptParams,omitempty"`

		// linux script
		LinuxScript string `json:"linuxScript,omitempty"`

		// linux script params
		LinuxScriptParams string `json:"linuxScriptParams,omitempty"`

		// The parameters for an external script
		ScriptParams string `json:"scriptParams,omitempty"`

		// The script path for an external script
		ScriptPath string `json:"scriptPath,omitempty"`

		// For script scans, the type of script. Options are embeded and external
		// Required: true
		ScriptType *string `json:"scriptType"`

		// windows script
		WindowsScript string `json:"windowsScript,omitempty"`

		// windows script params
		WindowsScriptParams string `json:"windowsScriptParams,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Collector *int32 `json:"collector"`

		CollectorDescription string `json:"collectorDescription,omitempty"`

		CollectorGroup int32 `json:"collectorGroup,omitempty"`

		CollectorGroupName string `json:"collectorGroupName,omitempty"`

		Creator string `json:"creator,omitempty"`

		Description string `json:"description,omitempty"`

		Duplicate *ExcludeDuplicateIps `json:"duplicate"`

		Group string `json:"group,omitempty"`

		ID int32 `json:"id,omitempty"`

		Method string `json:"method"`

		Name *string `json:"name"`

		NextStart string `json:"nextStart,omitempty"`

		NextStartEpoch int64 `json:"nextStartEpoch,omitempty"`

		NsgID int32 `json:"nsgId,omitempty"`

		Schedule *RestSchedule `json:"schedule,omitempty"`

		Version int32 `json:"version,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result ScriptNetscan

	result.collectorField = base.Collector

	result.collectorDescriptionField = base.CollectorDescription

	result.collectorGroupField = base.CollectorGroup

	result.collectorGroupNameField = base.CollectorGroupName

	result.creatorField = base.Creator

	result.descriptionField = base.Description

	result.duplicateField = base.Duplicate

	result.groupField = base.Group

	result.idField = base.ID

	if base.Method != result.Method() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid method value: %q", base.Method)
	}

	result.nameField = base.Name

	result.nextStartField = base.NextStart

	result.nextStartEpochField = base.NextStartEpoch

	result.nsgIdField = base.NsgID

	result.scheduleField = base.Schedule

	result.versionField = base.Version

	result.DefaultGroup = data.DefaultGroup

	result.DefaultGroupFullPath = data.DefaultGroupFullPath

	result.GroovyScript = data.GroovyScript

	result.GroovyScriptParams = data.GroovyScriptParams

	result.LinuxScript = data.LinuxScript

	result.LinuxScriptParams = data.LinuxScriptParams

	result.ScriptParams = data.ScriptParams

	result.ScriptPath = data.ScriptPath

	result.ScriptType = data.ScriptType

	result.WindowsScript = data.WindowsScript

	result.WindowsScriptParams = data.WindowsScriptParams

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ScriptNetscan) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The ID of the default group to add discovered devices to
		DefaultGroup int32 `json:"defaultGroup,omitempty"`

		// The full path of the default group to add discovered devices to
		DefaultGroupFullPath string `json:"defaultGroupFullPath,omitempty"`

		// For embedded script scans, the groovy script contents
		GroovyScript string `json:"groovyScript,omitempty"`

		// For embedded script scans, the groovy script parameters
		GroovyScriptParams string `json:"groovyScriptParams,omitempty"`

		// linux script
		LinuxScript string `json:"linuxScript,omitempty"`

		// linux script params
		LinuxScriptParams string `json:"linuxScriptParams,omitempty"`

		// The parameters for an external script
		ScriptParams string `json:"scriptParams,omitempty"`

		// The script path for an external script
		ScriptPath string `json:"scriptPath,omitempty"`

		// For script scans, the type of script. Options are embeded and external
		// Required: true
		ScriptType *string `json:"scriptType"`

		// windows script
		WindowsScript string `json:"windowsScript,omitempty"`

		// windows script params
		WindowsScriptParams string `json:"windowsScriptParams,omitempty"`
	}{

		DefaultGroup: m.DefaultGroup,

		DefaultGroupFullPath: m.DefaultGroupFullPath,

		GroovyScript: m.GroovyScript,

		GroovyScriptParams: m.GroovyScriptParams,

		LinuxScript: m.LinuxScript,

		LinuxScriptParams: m.LinuxScriptParams,

		ScriptParams: m.ScriptParams,

		ScriptPath: m.ScriptPath,

		ScriptType: m.ScriptType,

		WindowsScript: m.WindowsScript,

		WindowsScriptParams: m.WindowsScriptParams,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Collector *int32 `json:"collector"`

		CollectorDescription string `json:"collectorDescription,omitempty"`

		CollectorGroup int32 `json:"collectorGroup,omitempty"`

		CollectorGroupName string `json:"collectorGroupName,omitempty"`

		Creator string `json:"creator,omitempty"`

		Description string `json:"description,omitempty"`

		Duplicate *ExcludeDuplicateIps `json:"duplicate"`

		Group string `json:"group,omitempty"`

		ID int32 `json:"id,omitempty"`

		Method string `json:"method"`

		Name *string `json:"name"`

		NextStart string `json:"nextStart,omitempty"`

		NextStartEpoch int64 `json:"nextStartEpoch,omitempty"`

		NsgID int32 `json:"nsgId,omitempty"`

		Schedule *RestSchedule `json:"schedule,omitempty"`

		Version int32 `json:"version,omitempty"`
	}{

		Collector: m.Collector(),

		CollectorDescription: m.CollectorDescription(),

		CollectorGroup: m.CollectorGroup(),

		CollectorGroupName: m.CollectorGroupName(),

		Creator: m.Creator(),

		Description: m.Description(),

		Duplicate: m.Duplicate(),

		Group: m.Group(),

		ID: m.ID(),

		Method: m.Method(),

		Name: m.Name(),

		NextStart: m.NextStart(),

		NextStartEpoch: m.NextStartEpoch(),

		NsgID: m.NsgID(),

		Schedule: m.Schedule(),

		Version: m.Version(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this script netscan
func (m *ScriptNetscan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuplicate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScriptType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScriptNetscan) validateCollector(formats strfmt.Registry) error {

	if err := validate.Required("collector", "body", m.Collector()); err != nil {
		return err
	}

	return nil
}

func (m *ScriptNetscan) validateDuplicate(formats strfmt.Registry) error {

	if err := validate.Required("duplicate", "body", m.Duplicate()); err != nil {
		return err
	}

	if m.Duplicate() != nil {
		if err := m.Duplicate().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duplicate")
			}
			return err
		}
	}

	return nil
}

func (m *ScriptNetscan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *ScriptNetscan) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule()) { // not required
		return nil
	}

	if m.Schedule() != nil {
		if err := m.Schedule().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *ScriptNetscan) validateScriptType(formats strfmt.Registry) error {

	if err := validate.Required("scriptType", "body", m.ScriptType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScriptNetscan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScriptNetscan) UnmarshalBinary(b []byte) error {
	var res ScriptNetscan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
