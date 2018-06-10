package messages

import (
	"database/sql"

	"./count"
	"./emoji"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// SQLMessage Database friendly message struct
type SQLMessage struct {
	ID              string
	Type            int
	ChannelID       string
	Content         string
	Characters      int
	Timestamp       string
	EditedTimestamp string
	MentionRoles    string
	Tts             bool
	MentionEveryone bool
	Author          string
	Attachments     string
	Embeds          string
	Mentions        string
	Reactions       string
	Emojis          string
}

// InitMessagesTable Creates the messages table
func InitMessagesTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS messages (
		id text,
		type integer,
		channelID text,
		content text,
		characters integer,
		timestamp text,
		editedTimestamp text,
		mentionRoles text,
		tts integer,
		mentionEveryone integer,
		authorID text,
		attachmentIDs text,
		embedsIDs text,
		mentionsIDs text,
		reactionsIDs text,
		emojisIDs text
	)`)
	stmt.Exec()
}

// prepareMessage Transforms a discordgo.Message into a DB friendly message
func prepareMessage(message *discordgo.Message) SQLMessage {
	var convertedMsg SQLMessage
	convertedMsg.ID = message.ID
	convertedMsg.Type = int(message.Type)
	convertedMsg.ChannelID = message.ChannelID
	convertedMsg.Content = message.Content
	convertedMsg.Characters = count.CountCharacters(message.Content, message.Mentions)
	convertedMsg.Timestamp = string(message.Timestamp)
	convertedMsg.EditedTimestamp = string(message.EditedTimestamp)
	convertedMsg.MentionRoles = UpdateRoleMentions(message.MentionRoles)
	convertedMsg.Tts = message.Tts
	convertedMsg.MentionEveryone = UpdateEveryoneMentions(message.MentionEveryone)
	convertedMsg.Author = message.Author.ID
	convertedMsg.Attachments = AddAttachments(message.Attachments)
	convertedMsg.Embeds = AddEmbeds(message.Embeds)
	convertedMsg.Mentions = UpdateMemberMentions(message.Mentions)
	convertedMsg.Reactions = emoji.UpdateOrCreateReaction(message)
	convertedMsg.Emojis = emoji.UpdateOrCreateEmoji(message)

	return convertedMsg
}

// AddMessage Inserts a new entry in the messages table
func AddMessage(message *discordgo.Message, db *sql.DB) {
	tx, _ := db.Begin()

	msg := prepareMessage(message)
	count.UpdateCounters(message)

	stmt, _ := tx.Prepare(`INSERT INTO messages
		(id, type, channelID, content, characters, timestamp, editedTimestamp, mentionRoles, tts, mentionEveryone, authorID, attachmentIDs, embedsIDs, mentionsIDs, reactionsIDs, emojisIDs)
	values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	stmt.Exec(
		msg.ID,
		msg.Type,
		msg.ChannelID,
		msg.Content,
		msg.Characters,
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
		msg.Emojis)

	stmt2, _ := tx.Prepare(`UPDATE members SET messages = messages + 1 WHERE id = ?`)
	stmt2.Exec(msg.Author)
	tx.Commit()
}
