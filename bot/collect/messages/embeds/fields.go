package embeds

import (
	"database/sql"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// InitEmbedFieldsTable Creates the embedFields table
func InitEmbedFieldsTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS embedFields (
		id text,
		name text,
		value text,
		inline integer
	)`)
	stmt.Exec()
}

// AddEmbedField Inserts a new entry in the embedFields table
func AddEmbedField(ef []*discordgo.MessageEmbedField, db *sql.DB) string {
	var fieldsSlice []string
	for _, field := range ef {
		id := strconv.Itoa(rand.Int())
		fieldsSlice = append(fieldsSlice, id)

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

	return strings.Join(fieldsSlice, ",")
}
