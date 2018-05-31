package embeds

import (
	"database/sql"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// InitEmbedAuthorsTable Creates the embedAuthors table
func InitEmbedAuthorsTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS embedAuthors (
		id text,
		name text,
		iconURL text,
		proxyIconURL text
	)`)
	stmt.Exec()
}

// AddEmbedAuthor Inserts a new entry in the embedAuthors table
func AddEmbedAuthor(ea *discordgo.MessageEmbedAuthor, db *sql.DB) string {
	id := strconv.Itoa(rand.Int())
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`INSERT INTO embedAuthors
		(id, name, iconURL, proxyIconURL)
	values (?, ?, ?, ?)`)
	stmt.Exec(
		id,
		ea.Name,
		ea.IconURL,
		ea.ProxyIconURL)
	tx.Commit()

	return id
}
