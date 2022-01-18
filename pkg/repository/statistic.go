package redirect

import (
    "database/sql"
)

type Statistic struct {
	ID   int    `db:"id"`
	Ip string `db:"ip"`
	UserAgent string `db:"user_agent"`
	CreatedAt string `db:"created_at"`
}

type StatisticRepository struct {
	Connection *sql.DB
}

func (repository *StatisticRepository) add(statistic Statistic) (int, error) {
    var link LinkRedirect

    id := 0
    sqlStatement := `INSERT INTO statistics (ip, user_agent, created_at) VALUE ($1, $2, $3) RETURNING id`
    err := repository.Connection.QueryRow(sqlStatement, statistic.Ip, statistic.UserAgent, statistic.CreatedAt).Scan(&id)

    return id, err
}