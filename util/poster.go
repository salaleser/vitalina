package util

import (
	"github.com/bwmarrin/discordgo"
)

// Send publishes message.
func Send(s *discordgo.Session, m *discordgo.MessageCreate, msg Message) {
	if msg.Title == "" {
		return
	}

	footer := discordgo.MessageEmbedFooter{
		Text:    msg.FooterText,
		IconURL: msg.FooterIconURL,
	}

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: msg.ThumbnailURL,
	}

	image := discordgo.MessageEmbedImage{
		URL: msg.ImageURL,
	}

	embed := discordgo.MessageEmbed{
		Color:       100000,
		Title:       msg.Title,
		Description: msg.Description,
		Footer:      &footer,
		URL:         msg.Link,
		Thumbnail:   &thumbnail,
		Image:       &image,
	}

	message := discordgo.MessageSend{
		Embed: &embed,
	}

	s.ChannelMessageSendComplex(m.ChannelID, &message)
}

// SendInfo publishes information message.
func SendInfo(s *discordgo.Session, m *discordgo.MessageCreate, description string) {
	footer := discordgo.MessageEmbedFooter{
		Text: GetRandomLines(),
	}

	embed := discordgo.MessageEmbed{
		Color:       200,
		Title:       "[INF]",
		Description: description,
		Footer:      &footer,
	}

	message := discordgo.MessageSend{
		Embed: &embed,
	}

	s.ChannelMessageSendComplex(m.ChannelID, &message)
}

// SendError publishes error message.
func SendError(s *discordgo.Session, m *discordgo.MessageCreate, description string, err error) {
	footer := discordgo.MessageEmbedFooter{
		Text: err.Error(),
	}

	embed := discordgo.MessageEmbed{
		Color:       500,
		Title:       "[ERR]",
		Description: description,
		Footer:      &footer,
	}

	message := discordgo.MessageSend{
		Embed: &embed,
	}

	s.ChannelMessageSendComplex(m.ChannelID, &message)
}
