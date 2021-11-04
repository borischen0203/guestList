package handlers

import (
	"net/http"

	"github.com/getground/tech-tasks/backend/dto"
	"github.com/getground/tech-tasks/backend/errors"
	"github.com/getground/tech-tasks/backend/logger"
	"github.com/getground/tech-tasks/backend/services"
	"github.com/gin-gonic/gin"
)

func AddGuest(c *gin.Context) {
	request := dto.AddGuestRequest{}
	response := dto.GuestResponse{}
	request.Name = c.Param("name")
	c.BindJSON(&request)
	logger.Info.Printf("[AddGuest Handler] request=%+v\n", request)
	statusCode, result, err := services.AddGuestService(request, response)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 400:
		c.JSON(http.StatusBadRequest, err)
	case 403:
		c.JSON(http.StatusForbidden, err)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}

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

func UpdateAccomGuest(c *gin.Context) {
	request := dto.UpdateGuestRequest{}
	response := dto.GuestResponse{}
	request.Name = c.Param("name")
	c.BindJSON(&request)
	logger.Info.Printf("[UpdateAccomGuest Handler] request=%+v\n", request)
	statusCode, result, err := services.UpdateGuestService(request, response)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 400:
		c.JSON(http.StatusBadRequest, err)
	case 403:
		c.JSON(http.StatusForbidden, err)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}

}
func GetEmptySeats(c *gin.Context) {
	response := dto.EmptySeatsResponse{}
	statusCode, result, err := services.GetEmptySeatsNumberService(response)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}
}
