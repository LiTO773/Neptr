package count

import "database/sql"

// InitCharactersTable Creates the characters table
func InitCharactersTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS characters (
		character text,
		count integer
	)`)
	stmt.Exec()
}

// InitCharactersByUserTable Creates the charactersByUser table
func InitCharactersByUserTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS charactersByUser (
		user text,
		character text,
		count integer
	)`)
	stmt.Exec()
}

// InitCharactersByChannel Creates the charactersByChannel table
func InitCharactersByChannel(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS charactersByChannel (
		channel text,
		character text,
		count integer
	)`)
	stmt.Exec()
}
