package embeds

import (
	"database/sql"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// InitEmbedProvidersTable Creates the embedProviders table
func InitEmbedProvidersTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS embedProviders (
		id text,
		name text,
		url text
	)`)
	stmt.Exec()
}

// AddEmbedProvider Inserts a new entry in the embedProviders table
func AddEmbedProvider(ep *discordgo.MessageEmbedProvider, db *sql.DB) string {
	id := strconv.Itoa(rand.Int())
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`INSERT INTO embedProviders
		(id, name, url)
	values (?, ?, ?)`)
	stmt.Exec(
		id,
		ep.Name,
		ep.URL)
	tx.Commit()

	return id
}
