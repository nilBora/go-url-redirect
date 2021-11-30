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
    loadEnv()
	initDB()
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

    e.GET("/:code", doRedirect)

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

func doRedirect(c echo.Context) error {

    code := c.Param("code")
    var link LinkRedirect

    sqlStatement := `SELECT id, host, link, code FROM redirects WHERE code = $1`
    row := db.QueryRow(sqlStatement, code)
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


    return c.HTML(http.StatusOK, "Url: "+code+" Res: "+result)
    //return c.Redirect(301, link.Link)
    //return nil
    //
}

func loadEnv() {
    if err := godotenv.Load(); err != nil {
        //fmt.Println("No .env file found");
    }
}

func initDB() {
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

    err = db.Ping()
    if err != nil {
        panic(err)
    }
}