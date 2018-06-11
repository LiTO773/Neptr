package collect

import (
	"database/sql"

	"../../config"
)

// InitTimestampsTable Creates the timestamps table
func InitTimestampsTable() {
	// TODO There is still a lot to be done, like emojis, embeds, reactions, channels, etc.
	db, _ := sql.Open("sqlite3", config.DB)
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS timestamps (
		day text,
		hour text,
		messages integer DEFAULT 0,
		emojis integer DEFAULT 0,
		embeds integer DEFAULT 0
	)`)
	stmt.Exec()
}
