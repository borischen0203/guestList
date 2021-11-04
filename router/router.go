package router

import (
	"log"

	"github.com/getground/tech-tasks/backend/handlers"
	"github.com/getground/tech-tasks/backend/logger"

	// _ "url-shortener/docs"

	"github.com/gin-gonic/gin"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
)

var Router *gin.Engine

func errorHandlingMiddleWare(log *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		log.Printf("unexpected error: %s\n", err.Error())
	}
}

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger(), errorHandlingMiddleWare(logger.Error))

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/health", handlers.HealthHandler)
	router.GET("/version", handlers.VersionHandler)

	router.POST("/guest_list/:name", handlers.AddGuest)

	router.GET("/guest_list")
	router.GET("/guests")
	router.GET("/seats_empty", handlers.GetEmptySeats)

	router.PUT("/guests/:name", handlers.UpdateAccomGuest)

	router.DELETE("/guests/:name", handlers.DeleteGuest)

	// router.GET(("/:id"), handlers.GetLongUrl)
	// router.POST("/api/url-shortener/v1/url", handlers.GenerateShortUrl)
	return router
}

func Setup() {
	Router = SetupRouter()
}
