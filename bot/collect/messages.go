package collect

import (
	"database/sql"
	"fmt"

	"./messages"

	"./messages/count"
	"./messages/embeds"
	"./messages/emoji"
	"./messages/timestamp"

	"../../config"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// InitMessagesTables Creates all tables necessary to process messages
func InitMessagesTables() {
	db, _ := sql.Open("sqlite3", config.DB)
	tx, _ := db.Begin()
	messages.InitMessagesTable(db)
	messages.InitAttachmentsTable(db)
	messages.InitMessagesTable(db)
	messages.InitEmbedsTable(db)
	embeds.InitEmbedAuthorsTable(db)
	embeds.InitEmbedFieldsTable(db)
	embeds.InitEmbedFootersTable(db)
	embeds.InitEmbedImagesTable(db)
	embeds.InitEmbedProvidersTable(db)
	embeds.InitEmbedThumbnailsTable(db)
	embeds.InitEmbedVideosTable(db)
	count.InitCharactersTable(db)
	count.InitCharactersByUserTable(db)
	count.InitCharactersByChannelTable(db)
	emoji.InitEmojisByUserTable(db)
	emoji.InitEmojisByChannelTable(db)
	timestamp.InitTimestampsTable(db)
	tx.Commit()
}

// GetMessagesData Inserts a new entry in the messages table
func GetMessagesData(msgs []*discordgo.Message) {
	fmt.Println("GetMessagesData")
	db, _ := sql.Open("sqlite3", config.DB)

	for _, message := range msgs {
		messages.AddMessage(message, db)
	}
}
