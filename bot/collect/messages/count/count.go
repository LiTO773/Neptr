package count

import (
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// CountCharacters Counts the number of characters in a string
func CountCharacters(content string, mentions []*discordgo.User) int {
	r := regexp.MustCompile(`<:(.*?):[0-9]{18}>`)
	content = r.ReplaceAllString(content, `:$1:`) // Removes emoji's id and <>

	length := len(content)

	// Replace @<id> with @<username>
	for _, mention := range mentions {
		length += len(mention.Username)
	}
	length -= 20 * len(mentions)

	return length
}

// filterMessage Removes everything that the user can't see, like ids and <>
func filterMessage(content string, mentions []*discordgo.User) string {
	var sb strings.Builder
	r := regexp.MustCompile(`<:(.*?):[0-9]{18}>`)
	content = r.ReplaceAllString(content, `:$1:`) // Removes emoji's id and <>
	r = regexp.MustCompile(`(<@.*?>)`)
	content = r.ReplaceAllString(content, "@") // Removes mentions

	sb.WriteString(content)
	for _, mention := range mentions { // Adds the mentioned username
		sb.WriteString(mention.Username)
	}

	return sb.String()
}

// ProccessCharacters Gets each individual character/emoji and updates unicode emoji uses
func ProccessCharacters(message *discordgo.Message) map[string]int {
	content := []byte(filterMessage(message.Content, message.Mentions))
	chars := make(map[string]int)
	var specialCharacters []byte

	for i, char := range content {
		if char <= 127 {
			if len(specialCharacters) != 0 {
				chars = addToMap(chars, string(specialCharacters))
				specialCharacters = specialCharacters[:0]
			}
			chars = addToMap(chars, string(char))
		} else {
			specialCharacters = append(specialCharacters, char)
			if i == len(content)-1 {
				chars = addToMap(chars, string(specialCharacters))
			}
		}
	}

	return chars
}

func addToMap(chars map[string]int, name string) map[string]int {
	_, exists := chars[name]
	if exists {
		chars[name]++
	} else {
		chars[name] = 1
	}

	return chars
}
