package messages

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"../../../config"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// Returns a string as follows: "<reaction id> <uses>"
func UpdateReactions(reactions []*discordgo.MessageReactions) string {
	db, _ := sql.Open("sqlite3", config.DB)
	var reactionSlice []string

	for _, reaction := range reactions {
		tx, _ := db.Begin()

		// Check if it's not a server only emoji
		if !CheckIfEmojiExists(reaction.Emoji, db) {
			AddEmoji(reaction.Emoji, db, reaction.Count)
		} else {
			stmt, _ := tx.Prepare(`UPDATE emojis SET reactions = reactions + ? WHERE id = ?`)
			stmt.Exec(
				reaction.Count,
				reaction.Emoji.ID)
			tx.Commit()
		}

		reactionSlice = append(reactionSlice, reaction.Emoji.ID+" "+strconv.Itoa(reaction.Count))
	}

	return strings.Join(reactionSlice, ",")
}

func CheckIfEmojiExists(emoji *discordgo.Emoji, db *sql.DB) bool {
	rows, _ := db.Query("SELECT id,name FROM emojis WHERE (id=? AND name=?)", emoji.ID, emoji.Name)
	defer rows.Close()

	// Check if it exists or not
	found := false
	for rows.Next() {
		found = true
	}

	return found
}

// This might be moved to somewhere elese
func AddEmoji(emoji *discordgo.Emoji, db *sql.DB, reactions int) {
	tx, _ := db.Begin()

	// Generate an ID if it isn't an Unicode emoji
	id := emoji.ID
	if emoji.Name[0] == 240 {
		idSlice := []byte(emoji.Name)
		id = hex.EncodeToString(idSlice)
	}

	fmt.Println(id)

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
