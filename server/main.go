package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tempor1s/msconsole/server/db"
)

type response struct {
	Counter int `json:"counter"`
}

func main() {
	e := echo.New()

	d := db.New()
	d.AutoMigrate(db.CheckinCounter{Counter: 0})

	cs := db.NewCheckinCounterStore(d)

	// Just basic home page welcome message
	e.GET("/", homeRoute)
	// Add one to amount of times checked in.
	e.POST("/log", func(c echo.Context) error {
		counter, err := cs.UpdateCounter()
		if err != nil {
			log.Fatal(err)
		}

		resp := &response{Counter: counter.Counter}

		return c.JSON(http.StatusOK, resp)
	})
	// Get the amount of times checked in
	e.GET("/log", func(c echo.Context) error {
		counter, err := cs.GetCounter()
		if err != nil {
			log.Fatal(err)
		}

		resp := &response{Counter: counter.Counter}

		return c.JSON(http.StatusOK, resp)
	})

	e.Logger.Fatal(e.Start(":5000"))
}

// Just a basic hello message :)
func homeRoute(c echo.Context) error {
	return c.String(http.StatusOK,"Hello there. This is the basic MSConsole server :)")
}