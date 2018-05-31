package messages

import (
	"database/sql"
	"strings"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"

	"../../../config"
)

// InitAttachmentsTable Creates the attachments table
func InitAttachmentsTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS attachments (
		id text,
		url text,
		proxyURL text,
		filename text,
		width integer,
		height integer,
		size integer
	)`)
	stmt.Exec()
}

// AddAttachments Inserts a new entry in the attachments table
func AddAttachments(attachments []*discordgo.MessageAttachment) string {
	db, _ := sql.Open("sqlite3", config.DB)
	var attachmentSlice []string

	for _, attachment := range attachments {
		tx, _ := db.Begin()

		stmt, _ := tx.Prepare(`INSERT INTO attachments
			(id, url, proxyURL, filename, width, height, size)
		values (?, ?, ?, ?, ?, ?, ?)`)
		stmt.Exec(
			attachment.ID,
			attachment.URL,
			attachment.ProxyURL,
			attachment.Filename,
			attachment.Width,
			attachment.Height,
			attachment.Size)
		tx.Commit()

		attachmentSlice = append(attachmentSlice, attachment.ID)
	}

	return strings.Join(attachmentSlice, ",")
}
