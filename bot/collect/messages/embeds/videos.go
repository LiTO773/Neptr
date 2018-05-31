package embeds

import (
	"database/sql"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

func InitEmbedVideosTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS embedVideos (
		id text,
		url text,
		proxyURL text,
		width integer,
		height integer
	)`)
	stmt.Exec()
}

func AddEmbedVideo(ev *discordgo.MessageEmbedVideo, db *sql.DB) string {
	id := strconv.Itoa(rand.Int())
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare(`INSERT INTO embedVideos
		(id, url, proxyURL, width, height)
	values (?, ?, ?, ?, ?)`)
	stmt.Exec(
		id,
		ev.URL,
		ev.ProxyURL,
		ev.Width,
		ev.Height)
	tx.Commit()

	return id
}
