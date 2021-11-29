package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

    e.GET("/*", doRedirect)

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "3000"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func doRedirect(c echo.Context) error {
    uri := c.Request().URL.String()
    return c.HTML(http.StatusOK, "Url: "+uri)
}