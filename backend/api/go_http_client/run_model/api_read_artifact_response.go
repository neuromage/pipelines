// Code generated by go-swagger; DO NOT EDIT.

package run_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APIReadArtifactResponse api read artifact response
// swagger:model apiReadArtifactResponse
type APIReadArtifactResponse struct {

	// The bytes of the artifact content.
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`
}

// Validate validates this api read artifact response
func (m *APIReadArtifactResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIReadArtifactResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

// MarshalBinary interface implementation
func (m *APIReadArtifactResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIReadArtifactResponse) UnmarshalBinary(b []byte) error {
	var res APIReadArtifactResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
