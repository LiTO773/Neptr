package collect

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"../../config"
	"github.com/bwmarrin/discordgo"
)

// initMembersTable Creates the members table
func initMembersTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS members (
		id text,
		username text,
		nickname text,
		joinedAt text,
		messages integer,
		mentions integer
	)`)
	stmt.Exec()
}

func addMember(member *discordgo.Member, db *sql.DB) {
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`INSERT INTO members
		(id, username, nickname, joinedAt, messages, mentions) 
	values (?, ?, ?, ?, 0, 0)`)
	stmt.Exec(member.User.ID, member.User.Username, member.Nick, member.JoinedAt)
	tx.Commit()
}

// GetMembersData Inserts a new entry in the members table
func GetMembersData(members []*discordgo.Member) {
	fmt.Println("GetMembersData")
	db, _ := sql.Open("sqlite3", config.DB)
	initMembersTable(db)

	for _, member := range members {
		addMember(member, db)
	}
}
