package collect

import (
	"database/sql"
	"fmt"
	"strings"

	"../../config"
	"github.com/bwmarrin/discordgo"
)

func initEmojisTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS emojis (
		id text,
		name text,
		roles text,
		managed boolean,
		requireColons boolean,
		animated boolean,
		uses integer
	)`)
	stmt.Exec()
}

func addEmoji(emoji *discordgo.Emoji, db *sql.DB) {
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`INSERT INTO emojis
		(id, name, roles, managed, requireColons, animated)
	values (?, ?, ?, ?, ?, ?)`)
	stmt.Exec(
		emoji.ID,
		emoji.Name,
		strings.Join(emoji.Roles, ","),
		emoji.Managed,
		emoji.RequireColons,
		emoji.Animated)
	tx.Commit()
}

func GetEmojisData(emojis []*discordgo.Emoji) {
	fmt.Println("GetEmojisData")
	db, _ := sql.Open("sqlite3", config.DB)
	initEmojisTable(db)

	for _, emoji := range emojis {
		addEmoji(emoji, db)
	}
}
