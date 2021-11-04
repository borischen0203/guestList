package services

import (
	"database/sql"

	"github.com/getground/tech-tasks/backend/database"
	"github.com/getground/tech-tasks/backend/dto"
	e "github.com/getground/tech-tasks/backend/errors"
	"github.com/getground/tech-tasks/backend/logger"
)

func AddGuestService(r dto.AddGuestRequest, res dto.GuestResponse) (int64, dto.GuestResponse, e.ErrorInfo) {
	if ok, err := dto.Validate(r.Name); !ok {
		return 400, res, err
	}

	//check space
	sufficient, err := isSufficient(r.Name, r.Table, r.Accompanying_guests)
	if err != nil {
		logger.Error.Printf("[AddGuestService] check-space db error: %v\n", err)
		return 500, res, e.InternalServerError
	}
	if !sufficient {
		logger.Info.Printf("[AddGuestService] Not sufficient seats: %v\n", r)
		return 403, res, e.InsufficientSpaceError
	}

	//Insert guest
	result, err := insertGuest(r.Table, r.Name, r.Accompanying_guests)
	if err != nil {
		logger.Error.Printf("[AddGuestService] insert db error: %v\n", err)
		return 500, res, e.InternalServerError
	}
	res.Name = result
	return 200, res, e.NoError
}

func DeleteGuestService(request dto.DeleteGuestRequest) (int64, e.ErrorInfo) {
	if ok, err := dto.Validate(request.Name); !ok {
		return 400, err
	}
	err := deleteData(request.Name)
	if err != nil {
		logger.Error.Printf("[DeleteGuestService] delete db error: %v\n", err)
		return 500, e.InternalServerError
	}
	return 204, e.NoError
}

func GetEmptySeatsNumberService(res dto.EmptySeatsResponse) (int64, dto.EmptySeatsResponse, e.ErrorInfo) {
	result, err := countEmptySeats()
	if err != nil {
		logger.Error.Printf("[GetEmptySeatsNumberService] query db error: %v\n", err)
		return 500, res, e.InternalServerError
	}
	res.Seats_empty = result
	return 200, res, e.NoError
}

func UpdateGuestService(r dto.UpdateGuestRequest, res dto.GuestResponse) (int64, dto.GuestResponse, e.ErrorInfo) {
	sufficient, err := allowUpdate(r.Name, r.Accompanying_guests)
	if err != nil {
		if err == sql.ErrNoRows {
			return 400, res, e.InvalidNameError
		}
		logger.Error.Printf("[GetEmptySeatsNumberService] query db error: %v\n", err)
		return 500, res, e.InternalServerError
	}
	if !sufficient {
		logger.Info.Printf("[UpdateGuestService] Not sufficient seats: %v\n", r)
		return 403, res, e.InsufficientSpaceError
	}

	res.Name = r.Name
	return 200, res, e.NoError
}

//Check table space whether available
func isSufficient(name string, table int, guestPartner int) (bool, error) {
	maxSpace, err := queryMaxSpaceByTable(table)
	if err != nil {
		logger.Error.Printf("[isSufficient] check max space process error: %v\n", err)
		return false, err
	}

	currentRestGuestNumber, err := countRestGuestNumberByTable(name, table)
	if err != nil {
		logger.Error.Printf("[isSufficient] check current Guest Number process error: %v\n", err)
		return false, err
	}

	guestSelf := 1
	if (currentRestGuestNumber + guestPartner + guestSelf) > maxSpace {
		return false, nil
	}
	return true, nil
}

func insertGuest(table int, name string, partner int) (string, error) {
	db := database.DbConn()
	_, err := db.Exec("INSERT INTO guest_info(table_number, name, accompanying_guests) VALUES(?, ?,?)", table, name, partner)
	if err != nil {
		logger.Error.Printf("[insertData] insert data error: %v\n", err)
		return "", err
	}
	return name, nil

}

func deleteData(name string) error {
	db := database.DbConn()
	_, err := db.Exec("DELETE FROM guest_info WHERE name=?", name)
	if err != nil {
		logger.Error.Printf("[deleteData] delete db error: %v\n", err)
		return err
	}
	return nil
}

func updateGuest(name string, partner int) (bool, error) {
	db := database.DbConn()
	_, err := db.Exec("UPDATE guest_info set accompanying_guests=? where name=?", partner, name)
	if err != nil {
		logger.Error.Printf("[updateGuest] update data error: %v\n", err)
		return false, err
	}
	return true, nil
}

func allowUpdate(name string, partner int) (bool, error) {
	tableNumber, err := queryTableByName(name)
	if err != nil {
		logger.Error.Printf("[allowUpdate] query table number error: %v\n", err)
		return false, err
	}
	ok, err := isSufficient(name, tableNumber, partner)
	if err != nil {
		logger.Error.Printf("[allowUpdate] check seats error: %v\n", err)
		return false, err
	}
	if !ok {
		logger.Info.Printf("[allowUpdate] Not sufficient seats: %v\n", name)
		return false, nil
	}
	_, err = updateGuest(name, partner)
	if err != nil {
		logger.Error.Printf("[allowUpdate] Update accompanying guest error: %v\n", err)
		return false, err
	}
	return true, nil
}

func queryTableByName(name string) (int, error) {
	db := database.DbConn()
	tableNumber := 0
	var n sql.NullInt32
	err := db.QueryRow("SELECT table_number FROM guest_info WHERE name = ?", name).Scan(&n)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error.Printf("[queryTableByName] Name not found error: %v\n", err)
			return tableNumber, err
		}
		logger.Error.Printf("[queryTableByName] query Table number error: %v\n", err)
		return -1, err
	}
	if n.Valid {
		tableNumber = int(n.Int32)
	}
	return tableNumber, nil

}

//Query current table max space
func queryMaxSpaceByTable(table int) (int, error) {
	db := database.DbConn()
	maxNumber := 0
	var n sql.NullInt32
	err := db.QueryRow("SELECT space FROM guest_table WHERE table_number = ?", table).Scan(&n)

	if err != nil {
		if err == sql.ErrNoRows {
			return maxNumber, nil
		}
		logger.Error.Printf("[queryMaxSpaceByTable] check max space process error: %v\n", err)
		return -1, err
	}
	if n.Valid {
		maxNumber = int(n.Int32)
	}
	return maxNumber, nil
}

//Count max seat number
func countMaxSpace() (int, error) {
	db := database.DbConn()
	maxNumber := 0
	var n sql.NullInt32
	err := db.QueryRow("SELECT SUM(space) FROM guest_table").Scan(&n)
	if err != nil {
		if err == sql.ErrNoRows {
			return maxNumber, nil
		}
		logger.Error.Printf("[countMaxSpace] count max space process error: %v\n", err)
		return -1, err
	}
	if n.Valid {
		maxNumber = int(n.Int32)
	}
	return maxNumber, nil
}

//Query current table max guest number
func countRestGuestNumberByTable(name string, table int) (int, error) {
	db := database.DbConn()
	currentRestNumber := 0
	var n sql.NullInt32
	err := db.QueryRow("SELECT COUNT(id)+SUM(accompanying_guests) FROM guest_info WHERE table_number = ? AND name!=?", table, name).Scan(&n)
	if err != nil {
		logger.Error.Printf("[countGuestNumberByTable] count current Guest Number process error: %v\n", err)
		return -1, err
	}
	if n.Valid {
		currentRestNumber = int(n.Int32)
	}
	return currentRestNumber, nil
}

//Count all guest number
func countAllGuestNumber() (int, error) {
	db := database.DbConn()
	guestNumber := 0
	var n sql.NullInt32
	err := db.QueryRow("SELECT COUNT(id)+SUM(accompanying_guests) FROM guest_info").Scan(&n)
	if err != nil {
		logger.Error.Printf("[countAllGuestNumber] count all Guest Number process error: %v\n", err)
		return -1, err
	}
	if n.Valid {
		guestNumber = int(n.Int32)
	}
	return guestNumber, nil
}

func countEmptySeats() (int, error) {
	allGuestNumber, err := countAllGuestNumber()
	if err != nil {
		logger.Error.Printf("[countEmptySeats] count all Guest Number process error: %v\n", err)
		return -1, err
	}
	maxSpace, err := countMaxSpace()
	if err != nil {
		logger.Error.Printf("[countEmptySeats] count max space process error: %v\n", err)
		return -1, err
	}
	result := maxSpace - allGuestNumber
	return result, nil

}
