package dbConnector

import (
    "fmt"
	"os"
    "database/sql"
)

type DB struct {
	connection *sql.DB
}

func InitDB() DB {
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