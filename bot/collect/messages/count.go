package messages

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

// CountCharacters Counts the number of characters in a string
func CountCharacters(content string, mentions []*discordgo.User) int {
	length := len(content)

	// Remove the ids of custom emojis
	length -= 20 * len(CountEmojis(content))

	// Replace @<id> with @<username>
	for _, mention := range mentions {
		length += len(mention.Username)
	}
	length -= 20 * len(mentions)

	return length
}

// CountEmojis Gets all custom emojis in a message
func CountEmojis(content string) [][]string {
	r := regexp.MustCompile(`<:(.*?):[0-9]{18}>`)
	return r.FindAllStringSubmatch(content, -1)
}
