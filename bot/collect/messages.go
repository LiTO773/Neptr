package collect

import (
	"database/sql"
	"fmt"

	"./messages"

	"../../config"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// GetMessagesData Inserts a new entry in the messages table
func GetMessagesData(msgs []*discordgo.Message) {
	fmt.Println("GetMessagesData")
	db, _ := sql.Open("sqlite3", config.DB)
	messages.InitTables(db)

	for _, message := range msgs {
		messages.AddMessage(message, db)
	}
}
