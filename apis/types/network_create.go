// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NetworkCreate is the expected body of the "create network" http request message
// swagger:model NetworkCreate

type NetworkCreate struct {

	// Attachable means the network can be attached or not.
	Attachable bool `json:"Attachable,omitempty"`

	// CheckDuplicate is used to check the network is duplicate or not.
	CheckDuplicate bool `json:"CheckDuplicate,omitempty"`

	// config from
	ConfigFrom *ConfigReference `json:"ConfigFrom,omitempty"`

	// config only
	ConfigOnly bool `json:"ConfigOnly,omitempty"`

	// Driver means the network's driver.
	Driver string `json:"Driver,omitempty"`

	// enable ipv6
	EnableIPV6 bool `json:"EnableIPv6,omitempty"`

	// IP a m
	IPAM *IPAM `json:"IPAM,omitempty"`

	// Ingress checks the network is ingress network or not.
	Ingress bool `json:"Ingress,omitempty"`

	// Internal checks the network is internal network or not.
	Internal bool `json:"Internal,omitempty"`

	// labels
	Labels map[string]string `json:"Labels,omitempty"`

	// options
	Options map[string]string `json:"Options,omitempty"`

	// Scope means the network's scope.
	Scope string `json:"Scope,omitempty"`
}

/* polymorph NetworkCreate Attachable false */

/* polymorph NetworkCreate CheckDuplicate false */

/* polymorph NetworkCreate ConfigFrom false */

/* polymorph NetworkCreate ConfigOnly false */

/* polymorph NetworkCreate Driver false */

/* polymorph NetworkCreate EnableIPv6 false */

/* polymorph NetworkCreate IPAM false */

/* polymorph NetworkCreate Ingress false */

/* polymorph NetworkCreate Internal false */

/* polymorph NetworkCreate Labels false */

/* polymorph NetworkCreate Options false */

/* polymorph NetworkCreate Scope false */

// Validate validates this network create
func (m *NetworkCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigFrom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPAM(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkCreate) validateConfigFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigFrom) { // not required
		return nil
	}

	if m.ConfigFrom != nil {

		if err := m.ConfigFrom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ConfigFrom")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkCreate) validateIPAM(formats strfmt.Registry) error {

	if swag.IsZero(m.IPAM) { // not required
		return nil
	}

	if m.IPAM != nil {

		if err := m.IPAM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IPAM")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkCreate) UnmarshalBinary(b []byte) error {
	var res NetworkCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
