package messages

import (
	"database/sql"
	"strings"

	"../../../config"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

func UpdateMemberMentions(mentions []*discordgo.User) string {
	db, _ := sql.Open("sqlite3", config.DB)
	var mentionSlice []string

	for _, mention := range mentions {
		tx, _ := db.Begin()

		stmt, _ := tx.Prepare(`UPDATE members SET mentions = mentions + 1 WHERE id = ?`)
		stmt.Exec(mention.ID)
		tx.Commit()

		mentionSlice = append(mentionSlice, mention.ID)
	}

	return strings.Join(mentionSlice, ",")
}

func UpdateRoleMentions(roles []string) string {
	db, _ := sql.Open("sqlite3", config.DB)

	for _, role := range roles {
		tx, _ := db.Begin()

		stmt, _ := tx.Prepare(`UPDATE roles SET mentions = mentions + 1 WHERE id = ?`)
		stmt.Exec(role)
		tx.Commit()
	}

	return strings.Join(roles, ",")
}

func UpdateEveryoneMentions(mention bool) bool {
	if mention {
		db, _ := sql.Open("sqlite3", config.DB)

		tx, _ := db.Begin()
		stmt, _ := tx.Prepare(`UPDATE roles SET mentions = mentions + 1 WHERE name = "@everyone"`)
		stmt.Exec()
		tx.Commit()
	}

	return mention
}
