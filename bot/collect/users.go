package collect

import (
	"database/sql"
	"fmt"

	"../../config"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// initUsersTable Creates the users table
func initUsersTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS users (
		id text,
		username text,
		discriminator text,
		avatar text,
		bot integer
	)`)
	stmt.Exec()
}

func addUser(user *discordgo.User, db *sql.DB) {
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`INSERT INTO users
		(id, username, discriminator, avatar, bot)
	values (?, ?, ?, ?, ?)`)
	stmt.Exec(
		user.ID,
		user.Username,
		user.Discriminator,
		user.AvatarURL(""),
		user.Bot)
	tx.Commit()
}

func getUsers(members []*discordgo.Member) []*discordgo.User {
	var userSlice []*discordgo.User
	for _, member := range members {
		userSlice = append(userSlice, member.User)
	}
	return userSlice
}

// GetUsersData Inserts a new entry in the users table
func GetUsersData(members []*discordgo.Member) {
	fmt.Println("GetUsersData")
	db, _ := sql.Open("sqlite3", config.DB)
	initUsersTable(db)

	users := getUsers(members)

	for _, user := range users {
		addUser(user, db)
	}
}
