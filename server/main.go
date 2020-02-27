package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Just basic home page welcome message
	e.GET("/", homeRoute)
	// Add one to amount of times checked in.
	e.POST("/log", logCheckinRoute)
	// Get the amount of times checked in
	e.GET("/log", getCheckinCount)

	e.Logger.Fatal(e.Start(":5000"))
}

// Just a basic hello message :)
func homeRoute(c echo.Context) error {
	return c.String(http.StatusOK,"Hello there. This is the basic MSConsole server :)")
}

// TODO: Add one to the checkin count
func logCheckinRoute(c echo.Context) error {
	return c.String(http.StatusOK, "hi there :)")
}

// TODO: Add checkin count
func getCheckinCount(c echo.Context) error {
	return c.String(http.StatusOK, "Checkin Count: 0")
}