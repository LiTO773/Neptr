package embeds

import (
	"database/sql"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

func InitEmbedImagesTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS embedImages (
		id text,
		url text,
		proxyURL text,
		width integer,
		height integer
	)`)
	stmt.Exec()
}

func AddEmbedImage(ei *discordgo.MessageEmbedImage, db *sql.DB) string {
	id := strconv.Itoa(rand.Int())
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`INSERT INTO embedImages
		(id, url, proxyURL, width, height)
	values (?, ?, ?, ?, ?)`)
	stmt.Exec(
		id,
		ei.URL,
		ei.ProxyURL,
		ei.Width,
		ei.Height)
	tx.Commit()

	return id
}
