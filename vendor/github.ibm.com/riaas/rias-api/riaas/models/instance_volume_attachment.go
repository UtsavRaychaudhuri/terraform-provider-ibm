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

// InstanceVolumeAttachment VolumeAttachment
// swagger:model instance_volume_attachment
type InstanceVolumeAttachment struct {

	// The date and time that the volume was attached
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// If set to true, this volume will be automatically deleted if the only server it is attached to is deleted
	DeleteVolumeOnInstanceDelete *bool `json:"delete_volume_on_instance_delete,omitempty"`

	// The URL for this volume interface
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The unique identifier for this volume interface
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The user-defined name for this volume interface
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// status
	// Enum: [attaching attached detaching]
	Status string `json:"status,omitempty"`

	// A collection of tags for this resource
	Tags []string `json:"tags,omitempty"`

	// type
	// Enum: [boot data]
	Type string `json:"type,omitempty"`

	// volume
	Volume *ResourceReference `json:"volume,omitempty"`
}

// Validate validates this instance volume attachment
func (m *InstanceVolumeAttachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceVolumeAttachment) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InstanceVolumeAttachment) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *InstanceVolumeAttachment) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InstanceVolumeAttachment) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

var instanceVolumeAttachmentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["attaching","attached","detaching"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceVolumeAttachmentTypeStatusPropEnum = append(instanceVolumeAttachmentTypeStatusPropEnum, v)
	}
}

const (

	// InstanceVolumeAttachmentStatusAttaching captures enum value "attaching"
	InstanceVolumeAttachmentStatusAttaching string = "attaching"

	// InstanceVolumeAttachmentStatusAttached captures enum value "attached"
	InstanceVolumeAttachmentStatusAttached string = "attached"

	// InstanceVolumeAttachmentStatusDetaching captures enum value "detaching"
	InstanceVolumeAttachmentStatusDetaching string = "detaching"
)

// prop value enum
func (m *InstanceVolumeAttachment) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceVolumeAttachmentTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceVolumeAttachment) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var instanceVolumeAttachmentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["boot","data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceVolumeAttachmentTypeTypePropEnum = append(instanceVolumeAttachmentTypeTypePropEnum, v)
	}
}

const (

	// InstanceVolumeAttachmentTypeBoot captures enum value "boot"
	InstanceVolumeAttachmentTypeBoot string = "boot"

	// InstanceVolumeAttachmentTypeData captures enum value "data"
	InstanceVolumeAttachmentTypeData string = "data"
)

// prop value enum
func (m *InstanceVolumeAttachment) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceVolumeAttachmentTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceVolumeAttachment) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *InstanceVolumeAttachment) validateVolume(formats strfmt.Registry) error {

	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceVolumeAttachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceVolumeAttachment) UnmarshalBinary(b []byte) error {
	var res InstanceVolumeAttachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
