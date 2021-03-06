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

// Resource Domain-defined abstraction of a deposited 'work'. Resource's abstractions are describable for our domain’s purposes, i.e. for management needs within our system.
// swagger:model Resource
type Resource struct {

	// URI for the JSON-LD context definitions.
	// Required: true
	AtContext *string `json:"@context"`

	// The content type of the resource. Selected from an established set of values.
	// Required: true
	AtType *string `json:"@type"`

	// access
	// Required: true
	Access *ResourceAccess `json:"access"`

	// administrative
	// Required: true
	Administrative *ResourceAdministrative `json:"administrative"`

	// Identifier retrieved from identification.sourceId that stands for analog or source identifier that this resource is a digital representation of.
	DedupeIdentifier string `json:"dedupeIdentifier,omitempty"`

	// The Agent (User, Group, Application, Department, other) that deposited the resource into TAQUITO.
	Depositor *Agent `json:"depositor,omitempty"`

	// Following version for the resource within TAQUITO.
	FollowingVersion string `json:"followingVersion,omitempty"`

	// identification
	// Required: true
	Identification *ResourceIdentification `json:"identification"`

	// Identifier for the resource within TAQUITO. UUID, unique for each new version of a TAQUITO resource.
	// Required: true
	Identifier *string `json:"identifier"`

	// Primary processing label (can be same as title) for a resource.
	// Required: true
	Label *string `json:"label"`

	// permissions
	Permissions *ResourcePermissions `json:"permissions,omitempty"`

	// Preceding version for the resource within TAQUITO.
	PrecedingVersion string `json:"precedingVersion,omitempty"`

	// structural
	// Required: true
	Structural *ResourceStructural `json:"structural"`

	// Version for the resource within TAQUITO.
	// Required: true
	Version *int64 `json:"version"`
}

// Validate validates this resource
func (m *Resource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtContext(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAtType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAccess(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAdministrative(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDepositor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIdentifier(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStructural(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resource) validateAtContext(formats strfmt.Registry) error {

	if err := validate.Required("@context", "body", m.AtContext); err != nil {
		return err
	}

	return nil
}

var resourceTypeAtTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http://sdr.sul.stanford.edu/models/sdr3-object.jsonld","http://sdr.sul.stanford.edu/models/sdr3-3d.jsonld","http://sdr.sul.stanford.edu/models/sdr3-agreement.jsonld","http://sdr.sul.stanford.edu/models/sdr3-book.jsonld","http://sdr.sul.stanford.edu/models/sdr3-document.jsonld","http://sdr.sul.stanford.edu/models/sdr3-geo.jsonld","http://sdr.sul.stanford.edu/models/sdr3-image.jsonld","http://sdr.sul.stanford.edu/models/sdr3-page.jsonld","http://sdr.sul.stanford.edu/models/sdr3-photograph.jsonld","http://sdr.sul.stanford.edu/models/sdr3-manuscript.jsonld","http://sdr.sul.stanford.edu/models/sdr3-map.jsonld","http://sdr.sul.stanford.edu/models/sdr3-media.jsonld","http://sdr.sul.stanford.edu/models/sdr3-track.jsonld","http://sdr.sul.stanford.edu/models/sdr3-webarchive-binary.jsonld","http://sdr.sul.stanford.edu/models/sdr3-webarchive-seed.jsonld"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceTypeAtTypePropEnum = append(resourceTypeAtTypePropEnum, v)
	}
}

const (
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3ObjectJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-object.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3ObjectJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-object.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr33dJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-3d.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr33dJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-3d.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3AgreementJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-agreement.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3AgreementJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-agreement.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3BookJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-book.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3BookJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-book.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3DocumentJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-document.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3DocumentJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-document.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3GeoJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-geo.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3GeoJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-geo.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3ImageJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-image.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3ImageJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-image.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3PageJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-page.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3PageJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-page.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3PhotographJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-photograph.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3PhotographJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-photograph.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3ManuscriptJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-manuscript.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3ManuscriptJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-manuscript.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3MapJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-map.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3MapJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-map.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3MediaJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-media.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3MediaJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-media.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3TrackJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-track.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3TrackJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-track.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3WebarchiveBinaryJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-webarchive-binary.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3WebarchiveBinaryJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-webarchive-binary.jsonld"
	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3WebarchiveSeedJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-webarchive-seed.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3WebarchiveSeedJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-webarchive-seed.jsonld"
)

// prop value enum
func (m *Resource) validateAtTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceTypeAtTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Resource) validateAtType(formats strfmt.Registry) error {

	if err := validate.Required("@type", "body", m.AtType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAtTypeEnum("@type", "body", *m.AtType); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access", "body", m.Access); err != nil {
		return err
	}

	if m.Access != nil {

		if err := m.Access.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("access")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateAdministrative(formats strfmt.Registry) error {

	if err := validate.Required("administrative", "body", m.Administrative); err != nil {
		return err
	}

	if m.Administrative != nil {

		if err := m.Administrative.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("administrative")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateDepositor(formats strfmt.Registry) error {

	if swag.IsZero(m.Depositor) { // not required
		return nil
	}

	if m.Depositor != nil {

		if err := m.Depositor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("depositor")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateIdentification(formats strfmt.Registry) error {

	if err := validate.Required("identification", "body", m.Identification); err != nil {
		return err
	}

	if m.Identification != nil {

		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("identifier", "body", m.Identifier); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if m.Permissions != nil {

		if err := m.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateStructural(formats strfmt.Registry) error {

	if err := validate.Required("structural", "body", m.Structural); err != nil {
		return err
	}

	if m.Structural != nil {

		if err := m.Structural.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("structural")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Resource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resource) UnmarshalBinary(b []byte) error {
	var res Resource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
