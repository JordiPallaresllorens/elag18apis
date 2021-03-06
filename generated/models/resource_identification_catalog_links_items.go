// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResourceIdentificationCatalogLinksItems resource identification catalog links items
// swagger:model resourceIdentificationCatalogLinksItems
type ResourceIdentificationCatalogLinksItems struct {

	// Catalog that is the source of the linked record.
	// Required: true
	Catalog *string `json:"catalog"`

	// Record identifier that is unique within the context of the linked record's catalog.
	// Required: true
	CatalogRecordID *string `json:"catalogRecordId"`

	// If the linked record should be automatically updated when the resource descriptive metadata changes.
	DeliverMetadata bool `json:"deliverMetadata,omitempty"`

	// If the resource descriptive metadata should be automatically updated when the linked record changes.
	DeriveMetadata bool `json:"deriveMetadata,omitempty"`

	// Metadata schema of the linked record.
	RecordSchema string `json:"recordSchema,omitempty"`

	// Whether the linked record describes a resource that is broader than, equivalent to, or narrower than the resource the resource represents.
	RecordScope string `json:"recordScope,omitempty"`
}

// Validate validates this resource identification catalog links items
func (m *ResourceIdentificationCatalogLinksItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalog(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCatalogRecordID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecordScope(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceIdentificationCatalogLinksItems) validateCatalog(formats strfmt.Registry) error {

	if err := validate.Required("catalog", "body", m.Catalog); err != nil {
		return err
	}

	return nil
}

func (m *ResourceIdentificationCatalogLinksItems) validateCatalogRecordID(formats strfmt.Registry) error {

	if err := validate.Required("catalogRecordId", "body", m.CatalogRecordID); err != nil {
		return err
	}

	return nil
}

var resourceIdentificationCatalogLinksItemsTypeRecordScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["broader","equivalent","narrower"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceIdentificationCatalogLinksItemsTypeRecordScopePropEnum = append(resourceIdentificationCatalogLinksItemsTypeRecordScopePropEnum, v)
	}
}

const (
	// ResourceIdentificationCatalogLinksItemsRecordScopeBroader captures enum value "broader"
	ResourceIdentificationCatalogLinksItemsRecordScopeBroader string = "broader"
	// ResourceIdentificationCatalogLinksItemsRecordScopeEquivalent captures enum value "equivalent"
	ResourceIdentificationCatalogLinksItemsRecordScopeEquivalent string = "equivalent"
	// ResourceIdentificationCatalogLinksItemsRecordScopeNarrower captures enum value "narrower"
	ResourceIdentificationCatalogLinksItemsRecordScopeNarrower string = "narrower"
)

// prop value enum
func (m *ResourceIdentificationCatalogLinksItems) validateRecordScopeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceIdentificationCatalogLinksItemsTypeRecordScopePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResourceIdentificationCatalogLinksItems) validateRecordScope(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecordScopeEnum("recordScope", "body", m.RecordScope); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceIdentificationCatalogLinksItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceIdentificationCatalogLinksItems) UnmarshalBinary(b []byte) error {
	var res ResourceIdentificationCatalogLinksItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
