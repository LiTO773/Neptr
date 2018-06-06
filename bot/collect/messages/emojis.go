package messages

import (
	"database/sql"
	"regexp"
	"strconv"
	"strings"

	"../../../config"
)

func UpdateEmojis(content string) string {
	r := regexp.MustCompile(`<:.*?:([0-9]{18})>`)
	match := r.FindAllStringSubmatch(content, -1)
	uses := make(map[string]int)

	// Checks if there are no repeated emojis
	for _, id := range match {
		_, exists := uses[id[1]]
		if exists {
			uses[id[1]]++
		} else {
			uses[id[1]] = 1
		}
	}

	// Creates a string of emojis and adds to the DB
	db, _ := sql.Open("sqlite3", config.DB)
	var emojiSlice []string

	for key, val := range uses {
		// Adds to the DB
		tx, _ := db.Begin()

		stmt, _ := tx.Prepare(`UPDATE emojis SET uses = uses + ? WHERE id = ?`)
		stmt.Exec(val, key)
		tx.Commit()

		// Adds to the string
		emojiSlice = append(emojiSlice, key+" "+strconv.Itoa(val))
	}

	return strings.Join(emojiSlice, ",")
}
