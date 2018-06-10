package emoji

import (
	"database/sql"

	"../../../../config"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// TODO Current string concatenation is a bit performance heavy

// UpdateOrCreateEmoji Updates or creates a used emoji in the emojis table
// Returns a string to be used in the messages table with all the emojis used
func UpdateOrCreateEmoji(message *discordgo.Message) string {
	emojiMap := ProccessEmojis(message.Content)
	operation(emojiMap, message, "uses")
	return TransformToString(emojiMap)
}

// UpdateOrCreateReaction Updates or creates a used reaction in the emojis table
// Returns a string to be used in the messages table with all the reactions used
func UpdateOrCreateReaction(message *discordgo.Message) string {
	emojiMap := ProccessReactions(message.Reactions)
	operation(emojiMap, message, "reactions")
	return TransformToString(emojiMap)
}

// operation Inserts into the db an emoji/reaction
func operation(emojiMap map[string]int, message *discordgo.Message, field string) {
	db, _ := sql.Open("sqlite3", config.DB)

	for emoji, uses := range emojiMap {
		tx, _ := db.Begin()

		updateEmojis(tx, emoji, uses, field)
		updateEmojisByUser(tx, emoji, uses, message.Author.ID, field)
		updateEmojisByChannel(tx, emoji, uses, message.ChannelID, field)

		tx.Commit()
	}
}

func updateEmojis(tx *sql.Tx, emoji string, uses int, field string) {
	operation := "UPDATE emojis SET " + field + " = " + field + " + ? WHERE id = ?"
	stmt, _ := tx.Prepare(operation)
	res, _ := stmt.Exec(uses, emoji)
	rows, _ := res.RowsAffected()

	// This will only run if it's a unicode emoji
	if rows == 0 {
		operation = "INSERT INTO emojis (id, name, " + field + ") values (?, ?, ?)"
		stmt, _ := tx.Prepare(operation)
		stmt.Exec(emoji, emoji, uses)
	}
}

func updateEmojisByUser(tx *sql.Tx, emoji string, uses int, user string, field string) {
	operation := "UPDATE emojisByUser SET " + field + " = " + field + " + ? WHERE emojiID = ? AND userID = ?"
	stmt, _ := tx.Prepare(operation)
	res, _ := stmt.Exec(uses, emoji, user)
	rows, _ := res.RowsAffected()

	if rows == 0 {
		operation = "INSERT INTO emojisByUser (userID, emojiID, " + field + ") values (?, ?, ?)"
		stmt, _ := tx.Prepare(operation)
		stmt.Exec(user, emoji, uses)
	}
}

func updateEmojisByChannel(tx *sql.Tx, emoji string, uses int, channel string, field string) {
	operation := "UPDATE emojisByChannel SET " + field + " = " + field + " + ? WHERE emojiID = ? AND channel = ?"
	stmt, _ := tx.Prepare(operation)
	res, _ := stmt.Exec(uses, emoji, channel)
	rows, _ := res.RowsAffected()

	if rows == 0 {
		operation = "INSERT INTO emojisByChannel (channel, emojiID, " + field + ") values (?, ?, ?)"
		stmt, _ := tx.Prepare(operation)
		stmt.Exec(channel, emoji, uses)
	}
}
