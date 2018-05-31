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

	if !strings.HasPrefix(m.Content, config.BotPrefix) || m.Author.ID == BotID || channel.GuildID == "" {
		return
	}

	// Collects every member's info
	if strings.Contains(m.Content, config.BotPrefix+"cm") {
		guild, _ := s.State.Guild(channel.GuildID)

		collect.GetMembersData(guild.Members)
	}

	// Collects every user's info
	if strings.Contains(m.Content, config.BotPrefix+"cu") {
		guild, _ := s.State.Guild(channel.GuildID)

		collect.GetUsersData(guild.Members)
	}

	// Collects all channels in a guild
	if strings.Contains(m.Content, config.BotPrefix+"cc") {
		guild, _ := s.State.Guild(channel.GuildID)

		collect.GetChannelsData(guild.Channels)
	}

	// Collects all roles in a guild
	if strings.Contains(m.Content, config.BotPrefix+"cr") {
		guild, _ := s.State.Guild(channel.GuildID)

		collect.GetRolesData(guild.Roles)
	}

	// Collects all emojis in a guild
	if strings.Contains(m.Content, config.BotPrefix+"ce") {
		guild, _ := s.State.Guild(channel.GuildID)

		collect.GetEmojisData(guild.Emojis)
	}

	// Collects latest 100 messages in a channel
	if strings.Contains(m.Content, config.BotPrefix+"ct") {
		command := strings.Split(m.Content, " ")
		if len(command) == 2 {
			fmt.Println("yeyah!")
			messages, _ := s.ChannelMessages(command[1], 100, "", "", "")
			collect.GetMessagesData(messages)
		}

		// collect.GetEmojisData(guild.Emojis)
	}
}
