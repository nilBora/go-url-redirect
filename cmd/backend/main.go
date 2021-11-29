package main

import (
    "fmt"
	"net/http"
	"os"
    "database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/joho/godotenv"
	_"github.com/lib/pq"
)

type LinkRedirect struct {
	ID   int    `db:"id"`
	Host string `db:"host"`
	Link string `db:"link"`
	Code string `db:"code"`
}

var db *sql.DB

func main() {

    if err := godotenv.Load(); err != nil {
		//fmt.Println("No .env file found");
	}

     psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
             "password=%s dbname=%s sslmode=disable",
             os.Getenv("APP_POSTGRESQL_HOST"),
             os.Getenv("APP_POSTGRESQL_PORT"),
             os.Getenv("APP_POSTGRESQL_USER"),
             os.Getenv("APP_POSTGRESQL_PASS"),
             os.Getenv("POSTGRESQL_DB"))

    var err error

     db, err = sql.Open("postgres", psqlInfo)
     if err != nil {
       panic(err)
     }
     defer db.Close()

     err = db.Ping()
     if err != nil {
        panic(err)
     }

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

     var link LinkRedirect

    sqlStatement := `SELECT id, host, link, code FROM redirects WHERE id = $1`
    row := db.QueryRow(sqlStatement, 1)
    queryErr := row.Scan(&link.ID, &link.Host, &link.Link, &link.Code)

    var result string

    switch queryErr {
        case sql.ErrNoRows:
          result = "No rows were returned!"
          return c.HTML(http.StatusOK, "No rows were returned!")
        case nil:
          result = link.Link
        default:
          panic(queryErr)
    }

    uri := c.Request().URL.String()
    return c.HTML(http.StatusOK, "Url: "+uri+" Res: "+result)
    //c.Redirect(301, "https://google.com")
    //return nil
    //
}