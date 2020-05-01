package util

import (
	"github.com/bwmarrin/discordgo"
)

func Send(s *discordgo.Session, m *discordgo.MessageCreate, title string, description string, url string,
	imageURL string, thumbnailURL string, footerText string, footerIconURL string) {
	footer := discordgo.MessageEmbedFooter{
		Text:    footerText,
		IconURL: footerIconURL,
	}

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: thumbnailURL,
	}

	image := discordgo.MessageEmbedImage{
		URL: imageURL,
	}

	embed := discordgo.MessageEmbed{
		Color:       100000,
		Title:       title,
		Description: description,
		Footer:      &footer,
		URL:         url,
		Thumbnail:   &thumbnail,
		Image:       &image,
	}

	message := discordgo.MessageSend{
		Embed: &embed,
	}

	s.ChannelMessageSendComplex(m.ChannelID, &message)
}

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
