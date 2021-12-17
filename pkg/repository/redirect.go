package redirect

import (
    "database/sql"
)

type LinkRedirect struct {
	ID   int    `db:"id"`
	Host string `db:"host"`
	Link string `db:"link"`
	Code string `db:"code"`
}

type RedirectRepository struct {
	Connection *sql.DB
}

func (repository *RedirectRepository) GetByCode(code string) (LinkRedirect, error) {
    var link LinkRedirect

    sqlStatement := `SELECT id, host, link, code FROM redirects WHERE code = $1`
    row := repository.Connection.QueryRow(sqlStatement, code)
    queryErr := row.Scan(&link.ID, &link.Host, &link.Link, &link.Code)

    return link, queryErr
}