package channels

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// InitTextChannelsTable Creates the channelsText table
func InitTextChannelsTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS channelsText (
		id text,
		position integer,
		name text,
		topic text,
		nsfw boolean,
		parentID text
	)`)
	stmt.Exec()
}

// AddTextChannel Inserts a new entry in the channelsText table
func AddTextChannel(channel *discordgo.Channel, db *sql.DB) {
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`INSERT INTO channelsText
		(id, position, name, topic, nsfw, parentID)
	values (?, ?, ?, ?, ?, ?)`)
	stmt.Exec(
		channel.ID,
		channel.Position,
		channel.Name,
		channel.Topic,
		channel.NSFW,
		channel.ParentID)
	tx.Commit()
}
