package timestamp

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"

	"../../../../config"
)

// UpdateTimestamp Updates the hour and day in the DB
func UpdateTimestamp(message *discordgo.Message) {
	day, hour := GetDayAndHour(string(message.Timestamp))
	reactions := len(message.Reactions)
	embeds := len(message.Embeds)

	db, _ := sql.Open("sqlite3", config.DB)
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`UPDATE timestamps SET 
		messages = messages + 1,
		reactions = reactions + ?,
		embeds = embeds + ? 
	WHERE day = ? AND hour = ? AND channel = ?`)
	res, _ := stmt.Exec(reactions, embeds, day, hour, message.ChannelID)
	rows, _ := res.RowsAffected()

	if rows == 0 {
		stmt, _ := tx.Prepare(`INSERT INTO timestamps 
			(day, hour, channel, messages, reactions, embeds) 
		VALUES (?, ?, ?, 1, ?, ?)`)
		stmt.Exec(day, hour, message.ChannelID, reactions, embeds)
	}

	tx.Commit()
}
