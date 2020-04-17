package util

import (
	"github.com/bwmarrin/discordgo"
)

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
