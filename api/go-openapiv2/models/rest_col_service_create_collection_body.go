// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestColServiceCreateCollectionBody rest col service create collection body
//
// swagger:model RestColServiceCreateCollectionBody
type RestColServiceCreateCollectionBody struct {

	// collection Id
	CollectionID string `json:"collectionId,omitempty"`

	// collection type
	CollectionType *APICollectionType `json:"collectionType,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// schemas
	Schemas []*APISchemaField `json:"schemas"`
}

// Validate validates this rest col service create collection body
func (m *RestColServiceCreateCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollectionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestColServiceCreateCollectionBody) validateCollectionType(formats strfmt.Registry) error {
	if swag.IsZero(m.CollectionType) { // not required
		return nil
	}

	if m.CollectionType != nil {
		if err := m.CollectionType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collectionType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collectionType")
			}
			return err
		}
	}

	return nil
}

func (m *RestColServiceCreateCollectionBody) validateSchemas(formats strfmt.Registry) error {
	if swag.IsZero(m.Schemas) { // not required
		return nil
	}

	for i := 0; i < len(m.Schemas); i++ {
		if swag.IsZero(m.Schemas[i]) { // not required
			continue
		}

		if m.Schemas[i] != nil {
			if err := m.Schemas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schemas" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("schemas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this rest col service create collection body based on the context it is used
func (m *RestColServiceCreateCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollectionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchemas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestColServiceCreateCollectionBody) contextValidateCollectionType(ctx context.Context, formats strfmt.Registry) error {

	if m.CollectionType != nil {

		if swag.IsZero(m.CollectionType) { // not required
			return nil
		}

		if err := m.CollectionType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collectionType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collectionType")
			}
			return err
		}
	}

	return nil
}

func (m *RestColServiceCreateCollectionBody) contextValidateSchemas(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Schemas); i++ {

		if m.Schemas[i] != nil {

			if swag.IsZero(m.Schemas[i]) { // not required
				return nil
			}

			if err := m.Schemas[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schemas" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("schemas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestColServiceCreateCollectionBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestColServiceCreateCollectionBody) UnmarshalBinary(b []byte) error {
	var res RestColServiceCreateCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}