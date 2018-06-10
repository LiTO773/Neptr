package count

import (
	"database/sql"

	"../../../../config"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// UpdateCounters Updates the counted characters
func UpdateCounters(message *discordgo.Message) {
	charMap := ProccessCharacters(message)
	db, _ := sql.Open("sqlite3", config.DB)

	for char, count := range charMap {
		tx, _ := db.Begin()

		updateCharacters(tx, char, count)
		updateCharactersByUser(tx, char, count, message.Author.ID)
		updateCharactersByChannel(tx, char, count, message.ChannelID)

		tx.Commit()
	}
}

func updateCharacters(tx *sql.Tx, char string, count int) {
	stmt, _ := tx.Prepare(`UPDATE characters SET count = count + ? WHERE character = ?`)
	res, _ := stmt.Exec(count, char)
	rows, _ := res.RowsAffected()

	if rows == 0 {
		stmt, _ := tx.Prepare(`INSERT INTO characters
			(character, count)
		values (?, ?)`)
		stmt.Exec(char, count)
	}
}

func updateCharactersByUser(tx *sql.Tx, char string, count int, user string) {
	stmt, _ := tx.Prepare(`UPDATE charactersByUser SET count = count + ? WHERE character = ? AND user = ?`)
	res, _ := stmt.Exec(count, char, user)
	rows, _ := res.RowsAffected()

	if rows == 0 {
		stmt, _ := tx.Prepare(`INSERT INTO charactersByUser
			(user, character, count)
		values (?, ?, ?)`)
		stmt.Exec(user, char, count)
	}
}

func updateCharactersByChannel(tx *sql.Tx, char string, count int, channel string) {
	stmt, _ := tx.Prepare(`UPDATE charactersByChannel SET count = count + ? WHERE character = ? AND channel = ?`)
	res, _ := stmt.Exec(count, char, channel)
	rows, _ := res.RowsAffected()

	if rows == 0 {
		stmt, _ := tx.Prepare(`INSERT INTO charactersByChannel
			(channel, character, count)
		values (?, ?, ?)`)
		stmt.Exec(channel, char, count)
	}
}
