package dbConnector

import (
    "fmt"
	"os"
    "database/sql"
)

type DB struct {
    host string
    port string
    user string
    pass string
    db_name string
	Connection *sql.DB
}

func Init() DB {
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
                 "password=%s dbname=%s sslmode=disable",
                 os.Getenv("APP_POSTGRESQL_HOST"),
                 os.Getenv("APP_POSTGRESQL_PORT"),
                 os.Getenv("APP_POSTGRESQL_USER"),
                 os.Getenv("APP_POSTGRESQL_PASS"),
                 os.Getenv("POSTGRESQL_DB"))

    var err error
    var dbContainer DB

    dbContainer.Connection, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }

    err = dbContainer.Connection.Ping()
    if err != nil {
        panic(err)
    }
    return dbContainer
}