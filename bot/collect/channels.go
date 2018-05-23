package collect

import (
	"database/sql"
	"fmt"

	"../../config"
	"./channels"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

func initTables(db *sql.DB) {
	channels.InitVoiceChannelsTable(db)
	channels.InitTextChannelsTable(db)
	channels.InitCategoriesTable(db)
}

func GetChannelsData(channelsArr []*discordgo.Channel) {
	fmt.Println("GetChannelsData")
	db, _ := sql.Open("sqlite3", config.DB)
	initTables(db)

	for _, channel := range channelsArr {
		if channel.Type == 0 {
			channels.AddTextChannel(channel, db)
			continue
		}
		if channel.Type == 2 {
			channels.AddVoiceChannel(channel, db)
			continue
		}
		if channel.Type == 4 {
			channels.AddCategory(channel, db)
			continue
		}
	}
}
