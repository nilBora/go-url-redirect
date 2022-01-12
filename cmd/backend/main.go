package main

import (
    //"fmt"
	"net/http"
	"os"
    "database/sql"

    "utils"
    "dbConnector"
    "repository/redirect"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_"github.com/lib/pq"
)

func main() {
    utils.LoadEnv()
	dbContainer := dbConnector.Init()
	//defer dbContainer.connection.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

    e.GET("/:code", doRedirect(dbContainer))

	httpPort := getHttpPort()

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func getHttpPort() string {
    httpPort := os.Getenv("HTTP_PORT")
    if httpPort == "" {
        httpPort = "3000"
    }
    return httpPort
}

func doRedirect(dbContainer dbConnector.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    code := c.Param("code")

    redirectRepository := new(redirect.RedirectRepository)

    redirectRepository.Connection = dbContainer.Connection
    link, err := redirectRepository.GetByCode(code)
    if (err == nil) {
        addStatistics(c)
        return c.HTML(http.StatusOK, "Url: "+code+" Res: "+link.Link)
    }

    if (err == sql.ErrNoRows) {
        return c.HTML(http.StatusOK, "No rows were returned!")
    }

    panic(err)
  }
}

func addStatistics(c echo.Context) {
    //
}