package bot

import (
	"fmt"
	"strings"

	"./collect"

	"../config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Get the guild
	channel, _ := s.State.Channel(m.ChannelID)

	// !strings.HasPrefix(m.Content, config.BotPrefix) || m.Author.ID == BotID ||
	if channel.GuildID == "" {
		return
	}

	// START - First time collection
	guild, _ := s.State.Guild(channel.GuildID)

	// Collects everything
	if strings.Contains(m.Content, config.BotPrefix+"start") {
		collect.GetMembersData(guild.Members)
		collect.GetUsersData(guild.Members)
		collect.GetChannelsData(guild.Channels)
		collect.GetRolesData(guild.Roles)
		collect.GetEmojisData(guild.Emojis)

		for _, channel := range collect.GetTextChannels(guild.Channels) {
			fmt.Printf("Collecting channel %v - ", channel.ID)
			messages, _ := s.ChannelMessages(channel.ID, 100, "", "", "")
			collect.GetMessagesData(messages)
		}

		fmt.Println("Done!")
		return
	}

	// Collects every member's info
	if strings.Contains(m.Content, config.BotPrefix+"cm") {
		collect.GetMembersData(guild.Members)
		return
	}

	// Collects every user's info
	if strings.Contains(m.Content, config.BotPrefix+"cu") {
		collect.GetUsersData(guild.Members)
		return
	}

	// Collects all channels in a guild
	if strings.Contains(m.Content, config.BotPrefix+"cc") {
		collect.GetChannelsData(guild.Channels)
		return
	}

	// Collects all roles in a guild
	if strings.Contains(m.Content, config.BotPrefix+"cr") {
		collect.GetRolesData(guild.Roles)
		return
	}

	// Collects all emojis in a guild
	if strings.Contains(m.Content, config.BotPrefix+"ce") {
		collect.GetEmojisData(guild.Emojis)
		return
	}

	// Collects latest 100 messages in a channel
	if strings.Contains(m.Content, config.BotPrefix+"ct") {
		command := strings.Split(m.Content, " ")
		if len(command) == 2 {
			messages, _ := s.ChannelMessages(command[1], 100, "", "", "")
			collect.GetMessagesData(messages)
		}
		return
	}

	// END - First time collection

	// START - Event triggered collection
	// collect.NewMessage()
	// END - Event triggered collection
}
