package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/salaleser/vitalina/util"
)

var commands map[string]string

// Help sends user manual
func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content[1:], " ")

	if len(args) == 1 {
		var builder strings.Builder

		builder.WriteString("Список поддерживаемых комад:\n")
		for title, description := range commands {
			builder.WriteString(title)
			builder.WriteString(" — ")
			builder.WriteString(description)
			builder.WriteByte('\n')
		}

		util.SendInfo(s, m, builder.String())
	}
}

// FIXME
func InitHelp(c map[string]string) {
	commands = c
}
