package emoji

import "database/sql"

// InitEmojisByUserTable Creates the emojisByUser table
func InitEmojisByUserTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS emojisByUser (
		userID text,
		emojiID text,
		reactions integer DEFAULT 0,
		uses integer DEFAULT 0
	)`)
	stmt.Exec()
}

// InitEmojisByChannel Creates the emojisByChannel table
func InitEmojisByChannelTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS emojisByChannel (
		channel text,
		emojiID text,
		reactions integer DEFAULT 0,
		uses integer DEFAULT 0
	)`)
	stmt.Exec()
}
