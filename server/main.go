package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/log", logCheckin)
	e.Logger.Fatal(e.Start(":5000"))
}

func logCheckin(c echo.Context) error {
	return c.String(http.StatusOK, "hi there :)")
}