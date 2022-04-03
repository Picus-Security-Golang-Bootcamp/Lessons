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

// Book book
//
// swagger:model Book
type Book struct {

	// author ID
	AuthorID int64 `json:"authorID,omitempty"`

	// genre
	// Required: true
	Genre *string `json:"genre"`

	// id
	ID int64 `json:"id,omitempty"`

	// isbn
	// Required: true
	Isbn *string `json:"isbn"`

	// name
	// Required: true
	Name *string `json:"name"`

	// page number
	// Required: true
	PageNumber *int64 `json:"pageNumber"`

	// release date
	// Required: true
	// Format: date
	ReleaseDate *strfmt.Date `json:"releaseDate"`
}

// Validate validates this book
func (m *Book) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGenre(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsbn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Book) validateGenre(formats strfmt.Registry) error {

	if err := validate.Required("genre", "body", m.Genre); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateIsbn(formats strfmt.Registry) error {

	if err := validate.Required("isbn", "body", m.Isbn); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Book) validatePageNumber(formats strfmt.Registry) error {

	if err := validate.Required("pageNumber", "body", m.PageNumber); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateReleaseDate(formats strfmt.Registry) error {

	if err := validate.Required("releaseDate", "body", m.ReleaseDate); err != nil {
		return err
	}

	if err := validate.FormatOf("releaseDate", "body", "date", m.ReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this book based on context it is used
func (m *Book) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Book) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Book) UnmarshalBinary(b []byte) error {
	var res Book
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
