package handlers

import (
	"fmt"
	"net/http"

	db "github.com/getground/tech-tasks/backend/database"
	"github.com/getground/tech-tasks/backend/dto"
	"github.com/getground/tech-tasks/backend/errors"
	"github.com/getground/tech-tasks/backend/logger"
	"github.com/getground/tech-tasks/backend/services"
	"github.com/gin-gonic/gin"
)

func GuestHandler(c *gin.Context) {
	db := db.DbConn()

	_, err := db.Exec("insert INTO guest_table(table_number, space) values(?,?)", 3, 10)
	if err != nil {
		fmt.Printf("Insert data failed,err:%v", err)
		return
	}
	// db.InsertData()
	// db := database.DbConn()
	// insDB, err := db.Prepare("INSERT INTO guest_table(table_number, space) VALUES(?,?)")
	// if err != nil {
	// 	fmt.Printf("Insert data failed,err:%v", err)
	// 	return
	// }
	// insDB.Exec(2, 10)
	// defer db.Close()
}

func AddGuest(c *gin.Context) {
	request := dto.AddGuestRequest{}
	// response := dto.GuestResponse{}
	request.Name = c.Param("name")
	c.BindJSON(&request)
	// fmt.Println(request.Name)
	// fmt.Println(request.)
	logger.Info.Printf("[AddGuest Handler] request=%+v\n", request)
	c.IndentedJSON(http.StatusCreated, request)
	// statusCode, result, err := services.AddGuestService(request, response)
	// switch statusCode {
	// case 200:
	// 	c.JSON(http.StatusOK, result)
	// case 400:
	// 	c.JSON(http.StatusBadRequest, err)
	// case 500:
	// 	c.JSON(http.StatusInternalServerError, err)
	// default:
	// 	c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	// }

}

func DeleteGuest(c *gin.Context) {
	request := dto.DeleteGuestRequest{}
	request.Name = c.Param("name")
	logger.Info.Printf("[DeleteGuest Handler] request=%+v\n", request)
	statusCode, err := services.DeleteGuestService(request)
	switch statusCode {
	case 204:
		c.JSON(http.StatusNoContent, nil)
	case 400:
		c.JSON(http.StatusBadRequest, err)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}
}
