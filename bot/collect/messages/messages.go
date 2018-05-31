package messages

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

type SQLMessage struct {
	ID              string
	Type            int
	ChannelID       string
	Content         string
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
}

func InitMessagesTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS messages (
		id text,
		type integer,
		channelID text,
		content text,
		timestamp text,
		editedTimestamp text,
		mentionRoles text,
		tts boolean,
		mentionEveryone integer,
		authorID text,
		attachmentIDs text,
		embedsIDs text,
		mentionsIDs text,
		reactionsIDs string
	)`)
	stmt.Exec()
}

func prepareMessage(message *discordgo.Message) SQLMessage {
	var convertedMsg SQLMessage
	convertedMsg.ID = message.ID
	convertedMsg.Type = int(message.Type)
	convertedMsg.ChannelID = message.ChannelID
	convertedMsg.Content = message.Content
	convertedMsg.Timestamp = string(message.Timestamp)
	convertedMsg.EditedTimestamp = string(message.EditedTimestamp)
	convertedMsg.MentionRoles = UpdateRoleMentions(message.MentionRoles)
	convertedMsg.Tts = message.Tts
	convertedMsg.MentionEveryone = UpdateEveryoneMentions(message.MentionEveryone)
	convertedMsg.Author = message.Author.ID
	convertedMsg.Attachments = AddAttachments(message.Attachments)
	convertedMsg.Embeds = AddEmbeds(message.Embeds)
	convertedMsg.Mentions = UpdateMemberMentions(message.Mentions)
	convertedMsg.Reactions = UpdateReactions(message.Reactions)

	return convertedMsg
}

func AddMessage(message *discordgo.Message, db *sql.DB) {
	tx, _ := db.Begin()

	msg := prepareMessage(message)

	stmt, _ := tx.Prepare(`INSERT INTO messages
		(id, type, channelID, content, timestamp, editedTimestamp, mentionRoles, tts, mentionEveryone, authorID, attachmentIDs, embedsIDs, mentionsIDs, reactionsIDs)
	values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	stmt.Exec(
		msg.ID,
		msg.Type,
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
		msg.Reactions)
	tx.Commit()
}
