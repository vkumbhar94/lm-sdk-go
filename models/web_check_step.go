// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
)

// WebCheckStep web check step
// swagger:model WebCheckStep
type WebCheckStep struct {

	// HTTP body
	HTTPBody string `json:"HTTPBody,omitempty"`

	// HTTP headers
	HTTPHeaders string `json:"HTTPHeaders,omitempty"`

	// HTTP method
	HTTPMethod string `json:"HTTPMethod,omitempty"`

	// HTTP version
	HTTPVersion string `json:"HTTPVersion,omitempty"`

	authField Authentication

	// description
	Description string `json:"description,omitempty"`

	// enable
	Enable interface{} `json:"enable,omitempty"`

	// follow redirection
	FollowRedirection interface{} `json:"followRedirection,omitempty"`

	// fullpage load
	FullpageLoad bool `json:"fullpageLoad,omitempty"`

	// invert match
	InvertMatch bool `json:"invertMatch,omitempty"`

	// keyword
	Keyword string `json:"keyword,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// match type
	MatchType string `json:"matchType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// post data edit type
	PostDataEditType string `json:"postDataEditType,omitempty"`

	// req script
	ReqScript string `json:"reqScript,omitempty"`

	// req type
	ReqType string `json:"reqType,omitempty"`

	// require auth
	RequireAuth bool `json:"requireAuth,omitempty"`

	// resp script
	RespScript string `json:"respScript,omitempty"`

	// resp type
	RespType string `json:"respType,omitempty"`

	// schema
	Schema string `json:"schema,omitempty"`

	// status code
	StatusCode string `json:"statusCode,omitempty"`

	// timeout
	Timeout int32 `json:"timeout,omitempty"`

	// type
	// Read Only: true
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// use default root
	UseDefaultRoot interface{} `json:"useDefaultRoot,omitempty"`
}

// Auth gets the auth of this base type
func (m *WebCheckStep) Auth() Authentication {
	return m.authField
}

// SetAuth sets the auth of this base type
func (m *WebCheckStep) SetAuth(val Authentication) {
	m.authField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *WebCheckStep) UnmarshalJSON(raw []byte) error {
	var data struct {
		HTTPBody string `json:"HTTPBody,omitempty"`

		HTTPHeaders string `json:"HTTPHeaders,omitempty"`

		HTTPMethod string `json:"HTTPMethod,omitempty"`

		HTTPVersion string `json:"HTTPVersion,omitempty"`

		Auth json.RawMessage `json:"auth,omitempty"`

		Description string `json:"description,omitempty"`

		Enable interface{} `json:"enable,omitempty"`

		FollowRedirection interface{} `json:"followRedirection,omitempty"`

		FullpageLoad bool `json:"fullpageLoad,omitempty"`

		InvertMatch bool `json:"invertMatch,omitempty"`

		Keyword string `json:"keyword,omitempty"`

		Label string `json:"label,omitempty"`

		MatchType string `json:"matchType,omitempty"`

		Name string `json:"name,omitempty"`

		Path string `json:"path,omitempty"`

		PostDataEditType string `json:"postDataEditType,omitempty"`

		ReqScript string `json:"reqScript,omitempty"`

		ReqType string `json:"reqType,omitempty"`

		RequireAuth bool `json:"requireAuth,omitempty"`

		RespScript string `json:"respScript,omitempty"`

		RespType string `json:"respType,omitempty"`

		Schema string `json:"schema,omitempty"`

		StatusCode string `json:"statusCode,omitempty"`

		Timeout int32 `json:"timeout,omitempty"`

		Type string `json:"type,omitempty"`

		URL string `json:"url,omitempty"`

		UseDefaultRoot interface{} `json:"useDefaultRoot,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propAuth Authentication
	if string(data.Auth) != "null" {
		auth, err := UnmarshalAuthentication(bytes.NewBuffer(data.Auth), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propAuth = auth
	}

	var result WebCheckStep

	// HTTPBody
	result.HTTPBody = data.HTTPBody

	// HTTPHeaders
	result.HTTPHeaders = data.HTTPHeaders

	// HTTPMethod
	result.HTTPMethod = data.HTTPMethod

	// HTTPVersion
	result.HTTPVersion = data.HTTPVersion

	// auth
	result.authField = propAuth

	// description
	result.Description = data.Description

	// enable
	result.Enable = data.Enable

	// followRedirection
	result.FollowRedirection = data.FollowRedirection

	// fullpageLoad
	result.FullpageLoad = data.FullpageLoad

	// invertMatch
	result.InvertMatch = data.InvertMatch

	// keyword
	result.Keyword = data.Keyword

	// label
	result.Label = data.Label

	// matchType
	result.MatchType = data.MatchType

	// name
	result.Name = data.Name

	// path
	result.Path = data.Path

	// postDataEditType
	result.PostDataEditType = data.PostDataEditType

	// reqScript
	result.ReqScript = data.ReqScript

	// reqType
	result.ReqType = data.ReqType

	// requireAuth
	result.RequireAuth = data.RequireAuth

	// respScript
	result.RespScript = data.RespScript

	// respType
	result.RespType = data.RespType

	// schema
	result.Schema = data.Schema

	// statusCode
	result.StatusCode = data.StatusCode

	// timeout
	result.Timeout = data.Timeout

	// type
	result.Type = data.Type

	// url
	result.URL = data.URL

	// useDefaultRoot
	result.UseDefaultRoot = data.UseDefaultRoot

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m WebCheckStep) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		HTTPBody string `json:"HTTPBody,omitempty"`

		HTTPHeaders string `json:"HTTPHeaders,omitempty"`

		HTTPMethod string `json:"HTTPMethod,omitempty"`

		HTTPVersion string `json:"HTTPVersion,omitempty"`

		Description string `json:"description,omitempty"`

		Enable interface{} `json:"enable,omitempty"`

		FollowRedirection interface{} `json:"followRedirection,omitempty"`

		FullpageLoad bool `json:"fullpageLoad,omitempty"`

		InvertMatch bool `json:"invertMatch,omitempty"`

		Keyword string `json:"keyword,omitempty"`

		Label string `json:"label,omitempty"`

		MatchType string `json:"matchType,omitempty"`

		Name string `json:"name,omitempty"`

		Path string `json:"path,omitempty"`

		PostDataEditType string `json:"postDataEditType,omitempty"`

		ReqScript string `json:"reqScript,omitempty"`

		ReqType string `json:"reqType,omitempty"`

		RequireAuth bool `json:"requireAuth,omitempty"`

		RespScript string `json:"respScript,omitempty"`

		RespType string `json:"respType,omitempty"`

		Schema string `json:"schema,omitempty"`

		StatusCode string `json:"statusCode,omitempty"`

		Timeout int32 `json:"timeout,omitempty"`

		Type string `json:"type,omitempty"`

		URL string `json:"url,omitempty"`

		UseDefaultRoot interface{} `json:"useDefaultRoot,omitempty"`
	}{

		HTTPBody: m.HTTPBody,

		HTTPHeaders: m.HTTPHeaders,

		HTTPMethod: m.HTTPMethod,

		HTTPVersion: m.HTTPVersion,

		Description: m.Description,

		Enable: m.Enable,

		FollowRedirection: m.FollowRedirection,

		FullpageLoad: m.FullpageLoad,

		InvertMatch: m.InvertMatch,

		Keyword: m.Keyword,

		Label: m.Label,

		MatchType: m.MatchType,

		Name: m.Name,

		Path: m.Path,

		PostDataEditType: m.PostDataEditType,

		ReqScript: m.ReqScript,

		ReqType: m.ReqType,

		RequireAuth: m.RequireAuth,

		RespScript: m.RespScript,

		RespType: m.RespType,

		Schema: m.Schema,

		StatusCode: m.StatusCode,

		Timeout: m.Timeout,

		Type: m.Type,

		URL: m.URL,

		UseDefaultRoot: m.UseDefaultRoot,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Auth Authentication `json:"auth,omitempty"`
	}{

		Auth: m.authField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this web check step
func (m *WebCheckStep) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebCheckStep) validateAuth(formats strfmt.Registry) error {

	if swag.IsZero(m.Auth()) { // not required
		return nil
	}

	if err := m.Auth().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("auth")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebCheckStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebCheckStep) UnmarshalBinary(b []byte) error {
	var res WebCheckStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
