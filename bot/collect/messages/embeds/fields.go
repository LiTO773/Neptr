package embeds

import (
	"database/sql"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

func InitEmbedFieldsTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS embedFields (
		id text,
		name text,
		value text,
		inline boolean
	)`)
	stmt.Exec()
}

func AddEmbedField(ef []*discordgo.MessageEmbedField, db *sql.DB) string {
	var sb strings.Builder
	for _, field := range ef {
		id := strconv.Itoa(rand.Int())
		sb.WriteString(id + ",")

		tx, _ := db.Begin()
		stmt, _ := tx.Prepare(`INSERT INTO embedFields
			(id, name, value, inline)
		values (?, ?, ?, ?)`)
		stmt.Exec(
			id,
			field.Name,
			field.Value,
			field.Inline)
		tx.Commit()
	}

	return sb.String()
}
