package emoji

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// ProccessReactions Gets all reactions and returns them in a map[string]int
func ProccessReactions(reactions []*discordgo.MessageReactions) map[string]int {
	response := make(map[string]int)
	for _, reaction := range reactions {
		response[reaction.Emoji.ID] = reaction.Count
	}

	return response
}

// ProccessEmojis Gets all emojis and returns them in a map[string]int
func ProccessEmojis(content string) map[string]int {
	// Guild emojis
	r := regexp.MustCompile(`<:.*?:([0-9]{18})>`)
	match := r.FindAllStringSubmatch(content, -1)
	response := make(map[string]int)

	// Checks if there are no repeated emojis
	for _, id := range match {
		_, exists := response[id[1]]
		if exists {
			response[id[1]]++
		} else {
			response[id[1]] = 1
		}
	}

	// Unicode emojis
	unicode := getUnicodeEmojis(content)
	for emoji, count := range unicode {
		response[emoji] = count
	}

	return response
}

func getUnicodeEmojis(content string) map[string]int {
	contentSlice := []byte(content)
	response := make(map[string]int)
	var specialCharacters []byte

	for i, char := range contentSlice {
		if char <= 127 && len(specialCharacters) != 0 {
			specialCharacters = specialCharacters[:0]
			response = addToMap(response, string(specialCharacters))
		} else if char == 240 {
			specialCharacters = append(specialCharacters, char)
		} else if char >= 128 {
			specialCharacters = append(specialCharacters, char)
			if i == len(content)-1 {
				response = addToMap(response, string(specialCharacters))
			}
		}
	}

	return response
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

// TransformToString Transforms a reactions/emojis map into a DB friendly string
func TransformToString(reMap map[string]int) string {
	response := []string{}
	for reaction, count := range reMap {
		response = append(response, reaction+" "+strconv.Itoa(count))
	}

	return strings.Join(response, ",")
}
