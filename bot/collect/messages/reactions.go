package messages

import (
	"database/sql"
	"encoding/hex"
	"strconv"
	"strings"

	"../../../config"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// UpdateReactions +1 to an emoji when used as a reaction
// Returns a string as follows: "<reaction id> <uses>"
func UpdateReactions(reactions []*discordgo.MessageReactions) string {
	db, _ := sql.Open("sqlite3", config.DB)
	var reactionSlice []string

	for _, reaction := range reactions {
		tx, _ := db.Begin()

		id := idGenerator(reaction.Emoji)

		// Check if it's not a server only emoji
		if !CheckIfEmojiExists(reaction.Emoji, db) {
			AddEmoji(reaction.Emoji, db, reaction.Count)
		} else {
			stmt, _ := tx.Prepare(`UPDATE emojis SET reactions = reactions + ? WHERE id = ?`)
			stmt.Exec(
				reaction.Count,
				id)
			tx.Commit()
		}
		reactionSlice = append(reactionSlice, id+" "+strconv.Itoa(reaction.Count))
	}

	return strings.Join(reactionSlice, ",")
}

// idGenerator Returns the id or a generated id in case it's a normal emoji (like this one 👌)
func idGenerator(emoji *discordgo.Emoji) string {
	id := emoji.ID
	if emoji.Name[0] == 240 {
		idSlice := []byte(emoji.Name)
		id = hex.EncodeToString(idSlice)
	}

	return id
}

// CheckIfEmojiExists Checks if an emoji is already in the 'emojis' table
// Returns a matching boolean
func CheckIfEmojiExists(emoji *discordgo.Emoji, db *sql.DB) bool {
	id := idGenerator(emoji)
	rows, _ := db.Query("SELECT id,name FROM emojis WHERE (id=? AND name=?)", id, emoji.Name)
	defer rows.Close()

	// Check if it exists or not
	found := false
	for rows.Next() {
		found = true
	}

	return found
}

// AddEmoji Inserts a new entry in the emojis table
func AddEmoji(emoji *discordgo.Emoji, db *sql.DB, reactions int) {
	tx, _ := db.Begin()

	// Generate an ID if it isn't an Unicode emoji
	id := idGenerator(emoji)

	stmt, _ := tx.Prepare(`INSERT INTO emojis
		(id, name, roles, managed, requireColons, animated, reactions)
	values (?, ?, ?, ?, ?, ?, ?)`)
	stmt.Exec(
		id,
		emoji.Name,
		strings.Join(emoji.Roles, ","),
		emoji.Managed,
		emoji.RequireColons,
		emoji.Animated,
		reactions)
	tx.Commit()
}
