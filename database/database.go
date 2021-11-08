package database

/**
 * This database mainly set up mysql db and handle data
 *
 * @author: Boris
 * @version: 2021-11-04
 *
 */

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/getground/tech-tasks/backend/config"
	"github.com/getground/tech-tasks/backend/dto"
	"github.com/getground/tech-tasks/backend/logger"
	_ "github.com/go-sql-driver/mysql"
)

//Connect to sql db
func DbConn() (db *sql.DB) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Env.DBUser, config.Env.DBPassword, config.Env.DBHost, config.Env.DBPort, config.Env.DBName)
	db, err := sql.Open("mysql", source)
	if err != nil {
		logger.Error.Fatalf("Setup MySQL connect error %+v\n", err)
	}
	return db
}

//Mysql db setup
func Setup() {
	// init mysql.
	db := DbConn()

	defer db.Close()
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("MySQL connection successful")
}

//Insert guest data
func InsertGuest(table int, name string, partner int) (string, error) {
	db := DbConn()
	_, err := db.Exec("INSERT INTO guest_info(table_number, name, accompanying_guests) VALUES(?, ?,?)", table, name, partner)
	if err != nil {
		logger.Error.Printf("[InsertGuest] insert guest db error: %v\n", err)
		return "", err
	}
	return name, nil

}

//Delete guest data
func DeleteGuest(name string) error {
	db := DbConn()
	_, err := db.Exec("DELETE FROM guest_info WHERE name=?", name)
	if err != nil {
		logger.Error.Printf("[DeleteGuest] delete guest db error: %v\n", err)
		return err
	}
	return nil
}

//Update guest data
func UpdateGuest(name string, partner int) (bool, error) {
	db := DbConn()
	_, err := db.Exec("UPDATE guest_info SET accompanying_guests=? WHERE name=?", partner, name)
	if err != nil {
		logger.Error.Printf("[UpdateGuest] update guest data error: %v\n", err)
		return false, err
	}
	return true, nil
}

//Query guest data by name
func QueryTableByName(name string) (int, error) {
	db := DbConn()
	tableNumber := 0
	var n sql.NullInt32
	err := db.QueryRow("SELECT table_number FROM guest_info WHERE name = ?", name).Scan(&n)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error.Printf("[QueryTableByName] Name not found error: %v\n", err)
			return tableNumber, err
		}
		logger.Error.Printf("[QueryTableByName] query Table number db error: %v\n", err)
		return -1, err
	}
	if n.Valid {
		tableNumber = int(n.Int32)
	}
	return tableNumber, nil

}

//Query table max seats by table number
func QueryMaxSpaceByTable(table int) (int, error) {
	db := DbConn()
	maxNumber := 0
	var n sql.NullInt32
	err := db.QueryRow("SELECT space FROM guest_table WHERE table_number = ?", table).Scan(&n)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error.Printf("[QueryMaxSpaceByTable] Table not found error: %v\n", err)
			return maxNumber, err
		}
		logger.Error.Printf("[QueryMaxSpaceByTable] check max space process error: %v\n", err)
		return -1, err
	}
	if n.Valid {
		maxNumber = int(n.Int32)
	}
	return maxNumber, nil
}

//Count all max seat number
func CountMaxSpace() (int, error) {
	db := DbConn()
	maxNumber := 0
	var n sql.NullInt32
	err := db.QueryRow("SELECT SUM(space) FROM guest_table").Scan(&n)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Info.Printf("[CountMaxSpace] Max space: %v\n", 0)
			return maxNumber, nil
		}
		logger.Error.Printf("[CountMaxSpace] count max space process error: %v\n", err)
		return -1, err
	}
	if n.Valid {
		maxNumber = int(n.Int32)
	}
	return maxNumber, nil
}

//Query current table max guest number
func CountGuestNumberByTable(name string, table int) (int, error) {
	db := DbConn()
	currentRestNumber := 0
	var n sql.NullInt32
	switch name {
	case "": //For POST
		err := db.QueryRow("SELECT COUNT(id)+SUM(accompanying_guests) FROM guest_info WHERE table_number = ?", table).Scan(&n)
		if err != nil {
			logger.Error.Printf("[CountGuestNumberByTable] count current Guest Number process error: %v\n", err)
			return -1, err
		}
	default: // For PUT
		err := db.QueryRow("SELECT COUNT(id)+SUM(accompanying_guests) FROM guest_info WHERE table_number = ? AND name!=?", table, name).Scan(&n)
		if err != nil {
			logger.Error.Printf("[CountGuestNumberByTable] count current Rest Guest Number process error: %v\n", err)
			return -1, err
		}
	}

	if n.Valid {
		currentRestNumber = int(n.Int32)
	}
	return currentRestNumber, nil
}

//Count all guest number
func CountAllGuestNumber() (int, error) {
	db := DbConn()
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

//Count all empty seats
func CountEmptySeats() (int, error) {
	allGuestNumber, err := CountAllGuestNumber()
	if err != nil {
		logger.Error.Printf("[CountEmptySeats] count all Guest Number process error: %v\n", err)
		return -1, err
	}
	maxSpace, err := CountMaxSpace()
	if err != nil {
		logger.Error.Printf("[CountEmptySeats] count max space process error: %v\n", err)
		return -1, err
	}
	result := maxSpace - allGuestNumber
	return result, nil

}

//Query all guests
func QueryGuests() ([]dto.Guest, error) {
	db := DbConn()
	rows, err := db.Query("select name,table_number,accompanying_guests from guest_info")
	if err != nil {
		logger.Error.Printf("[QueryGuest] Query guest error: %v\n", err)
		return []dto.Guest{}, err
	}
	guest := dto.Guest{}        //object
	guest_list := []dto.Guest{} // object slice

	for rows.Next() {
		err = rows.Scan(&guest.Name, &guest.Table, &guest.Accompanying_guests)
		if err != nil {
			logger.Error.Printf("[QueryGuests] Query next guest error: %v\n", err)
			return []dto.Guest{}, err
		}
		guest_list = append(guest_list, guest)
	}
	defer db.Close()
	return guest_list, nil
}

//Query all guests with arrived time
func QueryArrivedGuests() ([]dto.Guest, error) {
	db := DbConn()
	rows, err := db.Query("SELECT name,accompanying_guests,time_arrived FROM guest_info")
	if err != nil {
		logger.Error.Printf("[QueryArrivedGuests] Query guest error: %v\n", err)
		return []dto.Guest{}, err
	}
	guest := dto.Guest{}        //object
	guest_list := []dto.Guest{} // object slice

	for rows.Next() {
		err = rows.Scan(&guest.Name, &guest.Accompanying_guests, &guest.Time_arrived)
		if err != nil {
			logger.Error.Printf("[QueryArrivedGuests] Query next guest error: %v\n", err)
			return []dto.Guest{}, err
		}
		guest_list = append(guest_list, guest)
	}
	defer db.Close()
	return guest_list, nil
}
