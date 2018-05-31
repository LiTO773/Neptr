package channels

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

func InitCategoriesTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS channelsCategories (
		id text,
		position integer,
		name text,
		nsfw boolean
	)`)
	stmt.Exec()
}

func AddCategory(channel *discordgo.Channel, db *sql.DB) {
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`INSERT INTO channelsCategories
		(id, position, name, nsfw)
	values (?, ?, ?, ?)`)
	stmt.Exec(
		channel.ID,
		channel.Position,
		channel.Name,
		channel.NSFW)
	tx.Commit()
}
