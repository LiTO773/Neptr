package embeds

import (
	"database/sql"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

func InitEmbedFootersTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS embedFooters (
		id text,
		text text,
		iconURL text,
		proxyIconURL text
	)`)
	stmt.Exec()
}

func AddEmbedFooter(ef *discordgo.MessageEmbedFooter, db *sql.DB) string {
	id := strconv.Itoa(rand.Int())
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`INSERT INTO embedFooters
		(id, name, iconURL, proxyIconURL)
	values (?, ?, ?, ?)`)
	stmt.Exec(
		id,
		ef.Text,
		ef.IconURL,
		ef.ProxyIconURL)
	tx.Commit()

	return id
}
