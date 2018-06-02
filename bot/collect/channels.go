package collect

import (
	"database/sql"
	"fmt"

	"../../config"
	"./channels"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// initTables Creates all tables necessary to process channels
func initTables(db *sql.DB) {
	channels.InitVoiceChannelsTable(db)
	channels.InitTextChannelsTable(db)
	channels.InitCategoriesTable(db)
}

// GetChannelsData Filters the channel by it's type and then adds it to the channels table
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

// GetTextChannels Returns only text channels
func GetTextChannels(channelsArr []*discordgo.Channel) []*discordgo.Channel {
	var channels []*discordgo.Channel
	for _, channel := range channelsArr {
		if channel.Type == 0 {
			channels = append(channels, channel)
		}
	}

	return channels
}
