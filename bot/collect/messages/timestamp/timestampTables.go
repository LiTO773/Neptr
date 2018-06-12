package timestamp

import (
	"database/sql"
)

// InitTimestampsTable Creates the timestamps table
func InitTimestampsTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS timestamps (
		day text,
		hour text,
		channel text,
		messages integer DEFAULT 0,
		reactions integer DEFAULT 0,
		embeds integer DEFAULT 0
	)`)
	stmt.Exec()
}
