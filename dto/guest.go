package dto

import (
	e "github.com/getground/tech-tasks/backend/errors"
)

/**
 * @param Name                 the name of the guest
 * @param Table                the table number of guest
 * @param Accompanying_guests  the accompanying number of guest
 */
type AddGuestRequest struct {
	Name                string `param:"name" `
	Table               int    `json:"table" binding:"required"`
	Accompanying_guests int    `json:"accompanying_guests" binding:"required"`
}

type DeleteGuestRequest struct {
	Name string `param:"name" binding:"required"`
}

type UpdateGuestRequest struct {
	Name                string `param:"name" binding:"required"`
	Accompanying_guests int    `json:"accompanying_guests" binding:"required"`
}

type GuestResponse struct {
	Name string `json:"name" binding:"required"`
}

type EmptySeatsResponse struct {
	Seats_empty int `json:"seats_empty" binding:"required" `
}

type Guest struct {
	Name                string `json:"name" binding:"required"`
	Table               int    `json:"table,omitempty" `
	Accompanying_guests int    `json:"accompanying_guests" binding:"required"`
	Time_arrived        string `json:"time_arrived,omitempty"`
}

type GuestListResponse struct {
	Guests []Guest `json:"guests"`
}

type ArrivedGuest struct {
	Name                string `json:"name" binding:"required"`
	Accompanying_guests int    `json:"accompanying_guests" binding:"required"`
	Time_arrived        string `json:"time_arrived" binding:"required"`
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
