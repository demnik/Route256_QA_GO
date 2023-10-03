// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1RemoveDeviceV1Response v1 remove device v1 response
//
// swagger:model v1RemoveDeviceV1Response
type V1RemoveDeviceV1Response struct {

	// found
	Found bool `json:"found,omitempty"`
}

// Validate validates this v1 remove device v1 response
func (m *V1RemoveDeviceV1Response) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 remove device v1 response based on context it is used
func (m *V1RemoveDeviceV1Response) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1RemoveDeviceV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RemoveDeviceV1Response) UnmarshalBinary(b []byte) error {
	var res V1RemoveDeviceV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
