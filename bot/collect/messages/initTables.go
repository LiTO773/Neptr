package messages

import (
	"database/sql"

	"./count"
	"./embeds"
)

// InitTables Creates all tables necessary to process messages
func InitTables(db *sql.DB) {
	InitMessagesTable(db)
	InitAttachmentsTable(db)
	InitMessagesTable(db)
	InitEmbedsTable(db)
	embeds.InitEmbedAuthorsTable(db)
	embeds.InitEmbedFieldsTable(db)
	embeds.InitEmbedFootersTable(db)
	embeds.InitEmbedImagesTable(db)
	embeds.InitEmbedProvidersTable(db)
	embeds.InitEmbedThumbnailsTable(db)
	embeds.InitEmbedVideosTable(db)
	count.InitCharactersTable(db)
	count.InitCharactersByUserTable(db)
	count.InitCharactersByChannel(db)
}
