package collect

import (
	"database/sql"
	"fmt"

	"../../config"
	"./messages"
	"github.com/bwmarrin/discordgo"
)

// initEmojisTable Creates the emojis table
func initEmojisTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS emojis (
		id text,
		name text,
		roles text,
		managed boolean,
		requireColons boolean,
		animated boolean,
		reactions integer
	)`)
	stmt.Exec()
}

// GetChannelsData Inserts a new entry in the emojis table
func GetEmojisData(emojis []*discordgo.Emoji) {
	fmt.Println("GetEmojisData")
	db, _ := sql.Open("sqlite3", config.DB)
	initEmojisTable(db)

	for _, emoji := range emojis {
		messages.AddEmoji(emoji, db, 0)
	}
}
