package messages

import (
	"database/sql"
	"math/rand"
	"strconv"
	"strings"

	"./embeds"

	"../../../config"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

// treatedEmbed Database friendly embed struct
type treatedEmbed struct {
	ID          string
	URL         string
	Type        string
	Title       string
	Description string
	Timestamp   string
	Color       int
	Footer      string
	Image       string
	Thumbnail   string
	Video       string
	Provider    string
	Author      string
	Fields      string
}

// InitEmbedsTable Creates the embeds table
func InitEmbedsTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS embeds (
		id text,
		url text,
		type text,
		title text,
		description text,
		timestamp text,
		color int,
		footer text,
		image text,
		thumbnail text,
		video text,
		provider text,
		author text,
		fields text
	)`)
	stmt.Exec()
}

// treatEmbed Transforms a discordgo.MessageEmbed into a DB friendly embed
func treatEmbed(embed *discordgo.MessageEmbed, db *sql.DB) treatedEmbed {
	var newEmbed treatedEmbed
	newEmbed.ID = strconv.Itoa(rand.Int())
	newEmbed.URL = embed.URL
	newEmbed.Type = embed.Type
	newEmbed.Title = embed.Title
	newEmbed.Description = embed.Description
	newEmbed.Timestamp = embed.Timestamp
	newEmbed.Color = embed.Color
	if embed.Footer != nil {
		newEmbed.Footer = embeds.AddEmbedFooter(embed.Footer, db)
	}
	if embed.Image != nil {
		newEmbed.Image = embeds.AddEmbedImage(embed.Image, db)
	}
	if embed.Thumbnail != nil {
		newEmbed.Thumbnail = embeds.AddEmbedThumbnail(embed.Thumbnail, db)
	}
	if embed.Video != nil {
		newEmbed.Video = embeds.AddEmbedVideo(embed.Video, db)
	}
	if embed.Provider != nil {
		newEmbed.Provider = embeds.AddEmbedProvider(embed.Provider, db)
	}
	if embed.Author != nil {
		newEmbed.Author = embeds.AddEmbedAuthor(embed.Author, db)
	}
	if embed.Fields != nil {
		newEmbed.Fields = embeds.AddEmbedField(embed.Fields, db)
	}

	return newEmbed
}

// AddEmbeds Inserts a new entry in the embeds table
func AddEmbeds(embeds []*discordgo.MessageEmbed) string {
	db, _ := sql.Open("sqlite3", config.DB)
	var embedSlice []string

	for _, embed := range embeds {
		tx, _ := db.Begin()

		newEmbed := treatEmbed(embed, db)

		stmt, _ := tx.Prepare(`INSERT INTO embeds
			(id, url, type, title, description, timestamp, color, footer, image, thumbnail, video, provider, author, fields)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
		stmt.Exec(
			newEmbed.ID,
			newEmbed.URL,
			newEmbed.Type,
			newEmbed.Title,
			newEmbed.Description,
			newEmbed.Timestamp,
			newEmbed.Color,
			newEmbed.Footer,
			newEmbed.Image,
			newEmbed.Thumbnail,
			newEmbed.Video,
			newEmbed.Provider,
			newEmbed.Author,
			newEmbed.Fields)
		tx.Commit()

		embedSlice = append(embedSlice, newEmbed.ID)
	}

	return strings.Join(embedSlice, ",")
}
