package emoji

import (
	"database/sql"

	"../../../../config"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// TODO Make this file also handle reactions

// UpdateOrCreateEmoji Updates or creates a used emoji in the emojis table
// Returns a string to be used in the messages table with all the emojis used
func UpdateOrCreateEmoji(message *discordgo.Message) string {
	emojiMap := ProccessEmojis(message.Content)
	db, _ := sql.Open("sqlite3", config.DB)

	for emoji, uses := range emojiMap {
		tx, _ := db.Begin()

		updateEmojis(tx, emoji, uses)
		updateEmojisByUser(tx, emoji, uses, message.Author.ID)
		updateEmojisByChannel(tx, emoji, uses, message.ChannelID)

		tx.Commit()
	}

	return TransformToString(emojiMap)
}

func updateEmojis(tx *sql.Tx, emoji string, uses int) {
	stmt, _ := tx.Prepare(`UPDATE emojis SET uses = uses + ? WHERE id = ?`)
	res, _ := stmt.Exec(uses, emoji)
	rows, _ := res.RowsAffected()

	// This will only run if it's a unicode emoji
	if rows == 0 {
		stmt, _ := tx.Prepare(`INSERT INTO emojis
			(id, name, uses)
		values (?, ?, ?)`)
		stmt.Exec(emoji, emoji, uses)
	}
}

func updateEmojisByUser(tx *sql.Tx, emoji string, uses int, user string) {
	stmt, _ := tx.Prepare(`UPDATE emojisByUser SET uses = uses + ? WHERE emojiID = ? AND userID = ?`)
	res, _ := stmt.Exec(uses, emoji, user)
	rows, _ := res.RowsAffected()

	if rows == 0 {
		stmt, _ := tx.Prepare(`INSERT INTO emojisByUser
			(userID, emojiID, uses)
		values (?, ?, ?)`)
		stmt.Exec(user, emoji, uses)
	}
}

func updateEmojisByChannel(tx *sql.Tx, emoji string, uses int, channel string) {
	stmt, _ := tx.Prepare(`UPDATE emojisByChannel SET uses = uses + ? WHERE emojiID = ? AND channel = ?`)
	res, _ := stmt.Exec(uses, emoji, channel)
	rows, _ := res.RowsAffected()

	if rows == 0 {
		stmt, _ := tx.Prepare(`INSERT INTO emojisByChannel
			(channel, emojiID, uses)
		values (?, ?, ?)`)
		stmt.Exec(channel, emoji, uses)
	}
}
