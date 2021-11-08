package services

import (
	"database/sql"

	"github.com/getground/tech-tasks/backend/database"
	"github.com/getground/tech-tasks/backend/dto"
	e "github.com/getground/tech-tasks/backend/errors"
	"github.com/getground/tech-tasks/backend/logger"
)

//Add guest service
func AddGuestService(request dto.AddGuestRequest, res dto.GuestResponse) (int64, dto.GuestResponse, e.ErrorInfo) {
	if ok, err := dto.Validate(request.Name); !ok {
		logger.Error.Printf("[AddGuestService] Invalid input error: %v\n", err)
		return 400, res, e.InvalidInputError
	}

	//Check seats is sufficient
	sufficient, err := isSufficient("", request.Table, request.Accompanying_guests)
	if err != nil {
		logger.Error.Printf("[AddGuestService] check-space db error: %v\n", err)
		return 500, res, e.InternalServerError
	}
	//Seats are not sufficient
	if !sufficient {
		logger.Info.Printf("[AddGuestService] Not sufficient seats: %v\n", request)
		return 403, res, e.InsufficientSpaceError
	}

	//Insert guest
	result, err := database.InsertGuest(request.Table, request.Name, request.Accompanying_guests)
	if err != nil {
		logger.Error.Printf("[AddGuestService] Insert gurest db error: %v\n", err)
		return 500, res, e.InternalServerError
	}
	res.Name = result
	return 200, res, e.NoError
}

//Delete guest service
func DeleteGuestService(request dto.DeleteGuestRequest) (int64, e.ErrorInfo) {
	if ok, err := dto.Validate(request.Name); !ok {
		logger.Error.Printf("[DeleteGuestService] Invalid input error: %v\n", err)
		return 400, e.InvalidInputError
	}
	err := database.DeleteGuest(request.Name)
	if err != nil {
		logger.Error.Printf("[DeleteGuestService] delete guest db error: %v\n", err)
		return 500, e.InternalServerError
	}
	return 204, e.NoError
}

//Get empty seats number service
func GetEmptySeatsNumberService(response dto.EmptySeatsResponse) (int64, dto.EmptySeatsResponse, e.ErrorInfo) {
	result, err := database.CountEmptySeats()
	if err != nil {
		logger.Error.Printf("[GetEmptySeatsNumberService] query seats data db error: %v\n", err)
		return 500, response, e.InternalServerError
	}
	response.Seats_empty = result
	return 200, response, e.NoError
}

//Update accompanying guests number service
func UpdateGuestService(request dto.UpdateGuestRequest, response dto.GuestResponse) (int64, dto.GuestResponse, e.ErrorInfo) {
	sufficient, err := AllowUpdate(request.Name, request.Accompanying_guests)
	if err != nil {
		if err == sql.ErrNoRows {
			return 400, response, e.InvalidNameError
		}
		logger.Error.Printf("[UpdateGuestService] query db error: %v\n", err)
		return 500, response, e.InternalServerError
	}
	if !sufficient {
		logger.Info.Printf("[UpdateGuestService] Not sufficient seats: %v\n", request)
		return 403, response, e.InsufficientSpaceError
	}

	response.Name = request.Name
	return 200, response, e.NoError
}

//Get guest lists service
func GetGuestListService(response dto.GuestListResponse) (int64, dto.GuestListResponse, e.ErrorInfo) {
	result, err := database.QueryGuests()
	if err != nil {
		logger.Error.Printf("[GetGuestListService] query db error: %v\n", err)
		return 500, response, e.InternalServerError
	}
	response.Guests = result
	return 200, response, e.NoError
}

//Get guest lists service with arrived time
func GetArrivedGuestListService(response dto.GuestListResponse) (int64, dto.GuestListResponse, e.ErrorInfo) {
	result, err := database.QueryArrivedGuests()
	if err != nil {
		logger.Error.Printf("[GetArrivedGuestListService] query db error: %v\n", err)
		return 500, response, e.InternalServerError
	}
	response.Guests = result
	return 200, response, e.NoError
}

//Check guest accompanying guest number allow update
func AllowUpdate(name string, partner int) (bool, error) {
	tableNumber, err := database.QueryTableByName(name)
	if err != nil {
		logger.Error.Printf("[AllowUpdate] query table number error: %v\n", err)
		return false, err
	}
	ok, err := isSufficient(name, tableNumber, partner)
	if err != nil {
		logger.Error.Printf("[AllowUpdate] check seats error: %v\n", err)
		return false, err
	}
	if !ok {
		logger.Info.Printf("[AllowUpdate] Not sufficient seats: %v\n", name)
		return false, nil
	}
	_, err = database.UpdateGuest(name, partner)
	if err != nil {
		logger.Error.Printf("[AllowUpdate] Update accompanying guest error: %v\n", err)
		return false, err
	}
	return true, nil
}

//Check table seat is Sufficient
func isSufficient(name string, table int, guestPartner int) (bool, error) {
	maxSpace, err := database.QueryMaxSpaceByTable(table)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error.Printf("[isSufficient] Not found table: %v\n", err)
			return false, err
		}
		logger.Error.Printf("[isSufficient] check max space process error: %v\n", err)
		return false, err
	}

	currentGuestNumber, err := database.CountGuestNumberByTable(name, table)
	if err != nil {
		logger.Error.Printf("[isSufficient] check current Guest Number process error: %v\n", err)
		return false, err
	}

	guestSelf := 1
	if (currentGuestNumber + guestPartner + guestSelf) > maxSpace {
		logger.Info.Printf("[isSufficient] Not sufficient seats: %v\n%v", currentGuestNumber+guestPartner+guestSelf, maxSpace)
		return false, nil
	}
	return true, nil
}
