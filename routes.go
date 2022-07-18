package main

import (
	"github.com/labstack/echo/v4"
)

func routes(e *echo.Echo) {

	e.GET("/", hello)
	e.GET("/getfilm", Get)
	e.GET("/getfilm2", Create2)
	e.POST("/createfilm", Create)
	e.DELETE("/deletefilm", Delete)

}
