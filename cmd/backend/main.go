package main

import (
    "fmt"
	"net/http"
	"os"
    "database/sql"

    "utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_"github.com/lib/pq"
)

type LinkRedirect struct {
	ID   int    `db:"id"`
	Host string `db:"host"`
	Link string `db:"link"`
	Code string `db:"code"`
}

type DB struct {
	connection *sql.DB
}

type RedirectRepository struct {
	connection *sql.DB
}

// type IRedirectRepository interface {
//     getByCode(code string)(LinkRedirect, error)
// }

func (repository *RedirectRepository) getByCode(code string) (LinkRedirect, error) {
    var link LinkRedirect

    sqlStatement := `SELECT id, host, link, code FROM redirects WHERE code = $1`
    row := repository.connection.QueryRow(sqlStatement, code)
    queryErr := row.Scan(&link.ID, &link.Host, &link.Link, &link.Code)

    return link, queryErr
}

func main() {
    utils.LoadEnv()
	dbContainer := initDB()
	defer dbContainer.connection.Close()

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

func doRedirect(dbContainer DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    code := c.Param("code")
    var link LinkRedirect

    var redirectRepository RedirectRepository
    redirectRepository.connection = dbContainer.connection
    link, queryErr := redirectRepository.getByCode(code)

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
  }
}

func initDB() DB {
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
                 "password=%s dbname=%s sslmode=disable",
                 os.Getenv("APP_POSTGRESQL_HOST"),
                 os.Getenv("APP_POSTGRESQL_PORT"),
                 os.Getenv("APP_POSTGRESQL_USER"),
                 os.Getenv("APP_POSTGRESQL_PASS"),
                 os.Getenv("POSTGRESQL_DB"))

    var err error
    var dbContainer DB

    dbContainer.connection, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }

    err = dbContainer.connection.Ping()
    if err != nil {
        panic(err)
    }
    return dbContainer
}