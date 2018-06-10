package collect

// Sets up the emoji table and populates it with guild specific emojis
// Unicode emojis are handled in the messages/emoji package
import (
	"database/sql"
	"fmt"
	"strings"

	"../../config"
	"github.com/bwmarrin/discordgo"
)

// initEmojisTable Creates the emojis table
func initEmojisTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS emojis (
		id text,
		name text,
		roles text DEFAULT '',
		managed integer DEFAULT 0,
		requireColons integer DEFAULT 0,
		animated integer DEFAULT 0,
		reactions integer DEFAULT 0,
		uses integer DEFAULT 0
	)`)
	stmt.Exec()
}

// GetEmojisData Inserts a new entry in the emojis table
func GetEmojisData(emojis []*discordgo.Emoji) {
	fmt.Println("GetEmojisData")
	db, _ := sql.Open("sqlite3", config.DB)
	initEmojisTable(db)

	for _, emoji := range emojis {
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
}
