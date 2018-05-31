package embeds

import (
	"database/sql"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

func InitEmbedProvidersTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS embedProviders (
		id text,
		name text,
		url text
	)`)
	stmt.Exec()
}

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
