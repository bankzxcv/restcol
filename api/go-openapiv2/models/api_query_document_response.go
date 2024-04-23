// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIQueryDocumentResponse api query document response
//
// swagger:model apiQueryDocumentResponse
type APIQueryDocumentResponse struct {

	// docs
	Docs *APIGetDocumentResponse `json:"docs,omitempty"`
}

// Validate validates this api query document response
func (m *APIQueryDocumentResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIQueryDocumentResponse) validateDocs(formats strfmt.Registry) error {
	if swag.IsZero(m.Docs) { // not required
		return nil
	}

	if m.Docs != nil {
		if err := m.Docs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("docs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("docs")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this api query document response based on the context it is used
func (m *APIQueryDocumentResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDocs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIQueryDocumentResponse) contextValidateDocs(ctx context.Context, formats strfmt.Registry) error {

	if m.Docs != nil {

		if swag.IsZero(m.Docs) { // not required
			return nil
		}

		if err := m.Docs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("docs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("docs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIQueryDocumentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIQueryDocumentResponse) UnmarshalBinary(b []byte) error {
	var res APIQueryDocumentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}