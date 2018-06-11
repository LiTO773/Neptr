package timestamp

import (
	"database/sql"

	"../../../../config"
)

// UpdateTimestamp Updates the hour and day in the DB
func UpdateTimestamp(timestamp string) {
	day, hour := GetDayAndHour(timestamp)

	db, _ := sql.Open("sqlite3", config.DB)
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`UPDATE timestamps SET messages = messages + 1 WHERE day = ? AND hour = ?`)
	res, _ := stmt.Exec(day, hour)
	rows, _ := res.RowsAffected()

	if rows == 0 {
		stmt, _ := tx.Prepare(`INSERT INTO timestamps (day, hour, messages) values (?, ?, 1)`)
		stmt.Exec(day, hour)
	}

	tx.Commit()
}
