package count

import (
	"database/sql"

	"../../../../config"
	"github.com/bwmarrin/discordgo"
)

// UpdateCounters Updates the counted characters
func UpdateCounters(message *discordgo.Message) {
	charMap := proccessCharacters(message)
	db, _ := sql.Open("sqlite3", config.DB)

	for char, count := range charMap {
		tx, _ := db.Begin()

		stmt, _ := tx.Prepare(`UPDATE characters SET count = count + ? WHERE character = ?`)
		res, _ := stmt.Exec(count, char)
		rows, _ := res.RowsAffected()

		if rows == 0 {
			stmt, _ := tx.Prepare(`INSERT INTO characters
				(character, count)
			values (?, ?)`)
			stmt.Exec(char, count)
		}

		stmt, _ = tx.Prepare(`UPDATE charactersByUser SET count = count + ? WHERE character = ? AND user = ?`)
		res, _ = stmt.Exec(count, char, message.Author.ID)
		rows, _ = res.RowsAffected()

		if rows == 0 {
			stmt, _ := tx.Prepare(`INSERT INTO charactersByUser
				(user, character, count)
			values (?, ?, ?)`)
			stmt.Exec(message.Author.ID, char, count)
		}

		stmt, _ = tx.Prepare(`UPDATE charactersByChannel SET count = count + ? WHERE character = ? AND channel = ?`)
		res, _ = stmt.Exec(count, char, message.ChannelID)
		rows, _ = res.RowsAffected()

		if rows == 0 {
			stmt, _ := tx.Prepare(`INSERT INTO charactersByChannel
				(channel, character, count)
			values (?, ?, ?)`)
			stmt.Exec(message.ChannelID, char, count)
		}

		tx.Commit()
	}
}
