// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrinterStatus printer status
//
// swagger:model printer_status
type PrinterStatus struct {
	MessageCommand

	// bottom exposure num
	// Required: true
	BottomExposureNum *string `json:"BottomExposureNum"`

	// cur slice layer
	// Required: true
	CurSliceLayer *string `json:"CurSliceLayer"`

	// delay light
	// Required: true
	DelayLight *string `json:"DelayLight"`

	// ele speed
	// Required: true
	EleSpeed *string `json:"EleSpeed"`

	// filename
	// Required: true
	Filename *string `json:"Filename"`

	// init exposure
	// Required: true
	InitExposure *string `json:"InitExposure"`

	// layer thickness
	// Required: true
	LayerThickness *string `json:"LayerThickness"`

	// print exposure
	// Required: true
	PrintExposure *string `json:"PrintExposure"`

	// print height
	// Required: true
	PrintHeight *string `json:"PrintHeight"`

	// print remain time
	// Required: true
	PrintRemainTime *string `json:"PrintRemainTime"`

	// print status
	// Required: true
	PrintStatus *string `json:"PrintStatus"`

	// resin
	// Required: true
	Resin *string `json:"Resin"`

	// slice layer count
	// Required: true
	SliceLayerCount *string `json:"SliceLayerCount"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PrinterStatus) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MessageCommand
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MessageCommand = aO0

	// AO1
	var dataAO1 struct {
		BottomExposureNum *string `json:"BottomExposureNum"`

		CurSliceLayer *string `json:"CurSliceLayer"`

		DelayLight *string `json:"DelayLight"`

		EleSpeed *string `json:"EleSpeed"`

		Filename *string `json:"Filename"`

		InitExposure *string `json:"InitExposure"`

		LayerThickness *string `json:"LayerThickness"`

		PrintExposure *string `json:"PrintExposure"`

		PrintHeight *string `json:"PrintHeight"`

		PrintRemainTime *string `json:"PrintRemainTime"`

		PrintStatus *string `json:"PrintStatus"`

		Resin *string `json:"Resin"`

		SliceLayerCount *string `json:"SliceLayerCount"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BottomExposureNum = dataAO1.BottomExposureNum

	m.CurSliceLayer = dataAO1.CurSliceLayer

	m.DelayLight = dataAO1.DelayLight

	m.EleSpeed = dataAO1.EleSpeed

	m.Filename = dataAO1.Filename

	m.InitExposure = dataAO1.InitExposure

	m.LayerThickness = dataAO1.LayerThickness

	m.PrintExposure = dataAO1.PrintExposure

	m.PrintHeight = dataAO1.PrintHeight

	m.PrintRemainTime = dataAO1.PrintRemainTime

	m.PrintStatus = dataAO1.PrintStatus

	m.Resin = dataAO1.Resin

	m.SliceLayerCount = dataAO1.SliceLayerCount

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PrinterStatus) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MessageCommand)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		BottomExposureNum *string `json:"BottomExposureNum"`

		CurSliceLayer *string `json:"CurSliceLayer"`

		DelayLight *string `json:"DelayLight"`

		EleSpeed *string `json:"EleSpeed"`

		Filename *string `json:"Filename"`

		InitExposure *string `json:"InitExposure"`

		LayerThickness *string `json:"LayerThickness"`

		PrintExposure *string `json:"PrintExposure"`

		PrintHeight *string `json:"PrintHeight"`

		PrintRemainTime *string `json:"PrintRemainTime"`

		PrintStatus *string `json:"PrintStatus"`

		Resin *string `json:"Resin"`

		SliceLayerCount *string `json:"SliceLayerCount"`
	}

	dataAO1.BottomExposureNum = m.BottomExposureNum

	dataAO1.CurSliceLayer = m.CurSliceLayer

	dataAO1.DelayLight = m.DelayLight

	dataAO1.EleSpeed = m.EleSpeed

	dataAO1.Filename = m.Filename

	dataAO1.InitExposure = m.InitExposure

	dataAO1.LayerThickness = m.LayerThickness

	dataAO1.PrintExposure = m.PrintExposure

	dataAO1.PrintHeight = m.PrintHeight

	dataAO1.PrintRemainTime = m.PrintRemainTime

	dataAO1.PrintStatus = m.PrintStatus

	dataAO1.Resin = m.Resin

	dataAO1.SliceLayerCount = m.SliceLayerCount

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this printer status
func (m *PrinterStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MessageCommand
	if err := m.MessageCommand.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBottomExposureNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurSliceLayer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelayLight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEleSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitExposure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLayerThickness(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrintExposure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrintHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrintRemainTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrintStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSliceLayerCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrinterStatus) validateBottomExposureNum(formats strfmt.Registry) error {

	if err := validate.Required("BottomExposureNum", "body", m.BottomExposureNum); err != nil {
		return err
	}

	return nil
}

func (m *PrinterStatus) validateCurSliceLayer(formats strfmt.Registry) error {

	if err := validate.Required("CurSliceLayer", "body", m.CurSliceLayer); err != nil {
		return err
	}

	return nil
}

func (m *PrinterStatus) validateDelayLight(formats strfmt.Registry) error {

	if err := validate.Required("DelayLight", "body", m.DelayLight); err != nil {
		return err
	}

	return nil
}

func (m *PrinterStatus) validateEleSpeed(formats strfmt.Registry) error {

	if err := validate.Required("EleSpeed", "body", m.EleSpeed); err != nil {
		return err
	}

	return nil
}

func (m *PrinterStatus) validateFilename(formats strfmt.Registry) error {

	if err := validate.Required("Filename", "body", m.Filename); err != nil {
		return err
	}

	return nil
}

func (m *PrinterStatus) validateInitExposure(formats strfmt.Registry) error {

	if err := validate.Required("InitExposure", "body", m.InitExposure); err != nil {
		return err
	}

	return nil
}

func (m *PrinterStatus) validateLayerThickness(formats strfmt.Registry) error {

	if err := validate.Required("LayerThickness", "body", m.LayerThickness); err != nil {
		return err
	}

	return nil
}

func (m *PrinterStatus) validatePrintExposure(formats strfmt.Registry) error {

	if err := validate.Required("PrintExposure", "body", m.PrintExposure); err != nil {
		return err
	}

	return nil
}

func (m *PrinterStatus) validatePrintHeight(formats strfmt.Registry) error {

	if err := validate.Required("PrintHeight", "body", m.PrintHeight); err != nil {
		return err
	}

	return nil
}

func (m *PrinterStatus) validatePrintRemainTime(formats strfmt.Registry) error {

	if err := validate.Required("PrintRemainTime", "body", m.PrintRemainTime); err != nil {
		return err
	}

	return nil
}

func (m *PrinterStatus) validatePrintStatus(formats strfmt.Registry) error {

	if err := validate.Required("PrintStatus", "body", m.PrintStatus); err != nil {
		return err
	}

	return nil
}

func (m *PrinterStatus) validateResin(formats strfmt.Registry) error {

	if err := validate.Required("Resin", "body", m.Resin); err != nil {
		return err
	}

	return nil
}

func (m *PrinterStatus) validateSliceLayerCount(formats strfmt.Registry) error {

	if err := validate.Required("SliceLayerCount", "body", m.SliceLayerCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this printer status based on the context it is used
func (m *PrinterStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MessageCommand
	if err := m.MessageCommand.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PrinterStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrinterStatus) UnmarshalBinary(b []byte) error {
	var res PrinterStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
