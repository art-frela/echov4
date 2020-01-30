package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/assets/:id", assetID)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// assetID mock of the real handler
func assetID(c echo.Context) error {
	return c.String(http.StatusOK, "assetID")
}
