package collect

import (
	"database/sql"

	"../../config"
	"github.com/bwmarrin/discordgo"
)

// InitRPTable Creates the richPresencesByUser table
func InitRPTable(db *sql.DB) {
	tx, _ := db.Begin()
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS richPresencesByUser (
		user text,
		name text
	)`)
	stmt.Exec()
	tx.Commit()
}

// Probably will change the structure of the table to be name and a list of users
// Who used the RP (TODO)
func GetRP(presences []*discordgo.Presence) {
	db, _ := sql.Open("sqlite3", config.DB)
	InitRPTable(db)

	for _, rp := range presences {
		if rp.Game != nil {
			tx, _ := db.Begin()

			// Add the rich presence to the db
			stmt, _ := db.Prepare(`INSERT INTO richPresencesByUser(user,name) 
				SELECT ?, ?
				WHERE NOT EXISTS(SELECT 1 FROM richPresencesByUser WHERE user = ? AND name = ?)`)
			stmt.Exec(rp.User.ID, rp.Game.Name, rp.User.ID, rp.Game.Name)

			tx.Commit()
		}
	}
}
