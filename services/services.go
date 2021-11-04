package services

import (
	"github.com/getground/tech-tasks/backend/database"
	"github.com/getground/tech-tasks/backend/dto"
	e "github.com/getground/tech-tasks/backend/errors"
	"github.com/getground/tech-tasks/backend/logger"
)

func AddGuestService(r dto.AddGuestRequest, res dto.GuestResponse) (int64, dto.GuestResponse, e.ErrorInfo) {
	if ok, err := dto.Validate(r.Name); !ok {
		return 400, res, err
	}
	result, err := insertData(r.Table, r.Name, r.Accompanying_guests)
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

func insertData(table int, name string, partner int) (string, error) {
	db := database.DbConn()
	_, err := db.Exec("insert into guest_info(table_number, space) values(?, ?)", table, name, partner)
	if err != nil {
		logger.Error.Printf("[insertData] insert data error: %v\n", err)
		return "", err
	}
	return name, nil

}

func deleteData(name string) error {
	db := database.DbConn()
	_, err := db.Exec("delete from guest_info where name=?", name)
	if err != nil {
		logger.Error.Printf("[deleteData] delete db error: %v\n", err)
		return err
	}
	return nil
}
