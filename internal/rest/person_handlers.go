package rest

import (
	"crud/internal/logic"

	"github.com/labstack/echo/v4"
)

func RegisterPersonRoutes(e *echo.Echo, personHandler *logic.PersonHandler) {
	e.POST("/person/", personHandler.PostPersonHandler)
	e.GET("/person/", personHandler.GetPersonHandler)
	e.GET("/person/:id", personHandler.GetPersonIdHandler)
	e.PUT("/person/:id", personHandler.PutPersonIdHandler)
	e.DELETE("/person/:id", personHandler.DelPersonIdHandler)
}
