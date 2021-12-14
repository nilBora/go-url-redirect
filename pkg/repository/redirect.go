package redirect

type LinkRedirect struct {
	ID   int    `db:"id"`
	Host string `db:"host"`
	Link string `db:"link"`
	Code string `db:"code"`
}

type RedirectRepository struct {
	connection *sql.DB
}

func (repository *RedirectRepository) getByCode(code string) (LinkRedirect, error) {
    var link LinkRedirect

    sqlStatement := `SELECT id, host, link, code FROM redirects WHERE code = $1`
    row := repository.connection.QueryRow(sqlStatement, code)
    queryErr := row.Scan(&link.ID, &link.Host, &link.Link, &link.Code)

    return link, queryErr
}