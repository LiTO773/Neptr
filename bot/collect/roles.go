package collect

import (
	"database/sql"
	"fmt"

	"../../config"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// initRolesTable Creates the roles table
func initRolesTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS roles (
		id text,
		position integer,
		name text,
		managed integer,
		mentionable integer,
		hoist integer,
		color integer,
		permissions integer,
		mentions integer DEFAULT 0
	)`)
	stmt.Exec()
}

func addRole(role *discordgo.Role, db *sql.DB) {
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`INSERT INTO roles
		(id, position, name, managed, mentionable, hoist, color, permissions)
	values (?, ?, ?, ?, ?, ?, ?, ?)`)
	stmt.Exec(
		role.ID,
		role.Position,
		role.Name,
		role.Managed,
		role.Mentionable,
		role.Hoist,
		role.Color,
		role.Permissions)
	tx.Commit()
}

// GetRolesData Inserts a new entry in the roles table
func GetRolesData(roles []*discordgo.Role) {
	fmt.Println("GetRolesData")
	db, _ := sql.Open("sqlite3", config.DB)
	initRolesTable(db)

	for _, role := range roles {
		addRole(role, db)
	}
}
