package handlers

import (
	"net/http"

	"github.com/getground/tech-tasks/backend/dto"
	"github.com/getground/tech-tasks/backend/errors"
	"github.com/getground/tech-tasks/backend/logger"
	"github.com/getground/tech-tasks/backend/services"
	"github.com/gin-gonic/gin"
)

// @Summary Add guest
// @Description Add guest to sql db
// @Tags Add
// @Accept json
// @Produce json
// @Param body body dto.AddGuestRequest true "body"
// @Success 200 {object} dto.GuestResponse "name"
// @Failure 400 {object} errors.ErrorInfo "Invalid input"
// @Failure 403 {object} errors.ErrorInfo "Forbiden to add"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
// @Router /guest_list/:name [POST]
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

// @Summary Delete guest
// @Description Delete guest from db
// @Tags Delete
// @Success 204 Delete successful
// @Failure 400 {object} errors.ErrorInfo "Invalid input"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
// @Router /guests/:name [DELETE]
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

// @Summary Update guest
// @Description Update guest's Accompanying guest number
// @Tags UPDATE
// @Accept json
// @Produce json
// @Param body body dto.UpdateGuestRequest true "body"
// @Success 200 {object} dto.GuestResponse "name"
// @Failure 400 {object} errors.ErrorInfo "Invalid input"
// @Failure 403 {object} errors.ErrorInfo "Forbiden to update"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
// @Router /guests/:name [PUT]
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

// @Summary Get empty seats
// @Description get all empty seats number
// @Tags GET
// @Produce json
// @Success 200 {object} dto.EmptySeatsResponse "seats_empty"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
// @Router /seats_empty [GET]
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

// @Summary Get guest lists
// @Description get all guest list
// @Tags GET
// @Produce json
// @Success 200 {object} dto.GuestListResponse "guests[]"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
// @Router /guest_list [GET]
func GetGuestLists(c *gin.Context) {
	response := dto.GuestListResponse{}
	statusCode, result, err := services.GetGuestListService(response)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}
}

// @Summary Get guest lists with arrived time
// @Description get all guest list with arrived time
// @Tags GET
// @Produce json
// @Success 200 {object} dto.GuestListResponse "guests[]"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
// @Router /guests [GET]
func GetArrivedGuestLists(c *gin.Context) {
	response := dto.GuestListResponse{}
	statusCode, result, err := services.GetArrivedGuestListService(response)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}
}
