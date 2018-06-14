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

// GetMessagesData Inserts a new entries in the messages table
func GetMessagesData(msgs []*discordgo.Message) {
	fmt.Println("GetMessagesData")
	db, _ := sql.Open("sqlite3", config.DB)

	for _, message := range msgs {
		messages.AddMessage(message, db)
	}
}

// GetMessageData Inserts a new entries in the messages table
func GetMessageData(msg *discordgo.MessageCreate) {
	fmt.Println("GetMessageData")
	db, _ := sql.Open("sqlite3", config.DB)

	standardMsg := discordgo.Message{
		msg.ID,
		msg.ChannelID,
		msg.Content,
		msg.Timestamp,
		msg.EditedTimestamp,
		msg.MentionRoles,
		msg.Tts,
		msg.MentionEveryone,
		msg.Author,
		msg.Attachments,
		msg.Embeds,
		msg.Mentions,
		msg.Reactions,
		msg.Type}

	standardMsgPointer := &standardMsg
	fmt.Println(standardMsg)
	messages.AddMessage(standardMsgPointer, db)
}
