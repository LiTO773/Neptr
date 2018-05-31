package channels

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// InitVoiceChannelsTable Creates the channelsVoice table
func InitVoiceChannelsTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS channelsVoice (
		id text,
		position integer,
		name text,
		nsfw boolean,
		bitrate integer,
		parentID text
	)`)
	stmt.Exec()
}

// AddVoiceChannel Inserts a new entry in the channelsVoice table
func AddVoiceChannel(channel *discordgo.Channel, db *sql.DB) {
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`INSERT INTO channelsVoice
		(id, position, name, nsfw, bitrate, parentID)
	values (?, ?, ?, ?, ?, ?)`)
	stmt.Exec(
		channel.ID,
		channel.Position,
		channel.Name,
		channel.NSFW,
		channel.Bitrate,
		channel.ParentID)
	tx.Commit()
}
