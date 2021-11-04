package dto

import (
	e "github.com/getground/tech-tasks/backend/errors"
)

type AddGuestRequest struct {
	Name                string `param:"name" `
	Table               int    `json:"table" binding:"required"`
	Accompanying_guests int    `json:"accompanying_guests" binding:"required"`
}

type DeleteGuestRequest struct {
	Name string `param:"name" `
}

type GuestResponse struct {
	Name string `json:"name" binding:"required"`
}

func Validate(s string) (bool, e.ErrorInfo) {
	if s == "" {
		return false, e.InvalidInputError
	}
	return true, e.NoError
}

func (r DeleteGuestRequest) Validate() (bool, e.ErrorInfo) {
	if r.IsEmpty() {
		return false, e.InvalidInputError
	}
	return true, e.NoError
}

func (r DeleteGuestRequest) IsEmpty() bool {
	return r.Name == ""
}
