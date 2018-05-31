package collect

import (
	"database/sql"
	"fmt"

	"./messages"

	"../../config"
	"github.com/bwmarrin/discordgo"
)

func GetMessagesData(msgs []*discordgo.Message) {
	fmt.Println("GetMessagesData")
	db, _ := sql.Open("sqlite3", config.DB)
	messages.InitTables(db)

	for _, message := range msgs {
		messages.AddMessage(message, db)
	}
}

// TODO: Buscar pelos menos as últimas 100 msgs e po-las na DB
