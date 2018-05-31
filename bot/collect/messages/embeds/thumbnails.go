package embeds

import (
	"database/sql"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

func InitEmbedThumbnailsTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS embedThumbnails (
		id text,
		url text,
		proxyURL text,
		width integer,
		height integer
	)`)
	stmt.Exec()
}

func AddEmbedThumbnail(et *discordgo.MessageEmbedThumbnail, db *sql.DB) string {
	id := strconv.Itoa(rand.Int())
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`INSERT INTO embedThumbnails
		(id, url, proxyURL, width, height)
	values (?, ?, ?, ?, ?)`)
	stmt.Exec(
		id,
		et.URL,
		et.ProxyURL,
		et.Width,
		et.Height)
	tx.Commit()

	return id
}
