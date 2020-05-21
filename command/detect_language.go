package command

import (
	"bytes"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/salaleser/vitalina/util"
)

// DetectLanguage detects language and adds reactions.
func DetectLanguage(s *discordgo.Session, m *discordgo.MessageCreate) {
	detections := util.DetectLanguage(m.Content[1:])
	if len(detections) == 0 {
		return
	}

	language := util.Languages[detections[0].Language]
	time.Sleep(100)
	s.MessageReactionAdd(m.ChannelID, m.ID, language.Emoji)

	score := util.GetEmojiByDigit(detections[0].ConfidenceScore)
	time.Sleep(100)
	s.MessageReactionAdd(m.ChannelID, m.ID, score)

	var reliable string
	if detections[0].IsReliable {
		reliable = "✅"
	} else {
		reliable = "❌"
	}
	time.Sleep(100)
	s.MessageReactionAdd(m.ChannelID, m.ID, reliable)
}

// DetectLanguageFull detects language and sends message.
func DetectLanguageFull(s *discordgo.Session, m *discordgo.MessageCreate) {
	detections := util.DetectLanguage(m.Content[1:])
	if len(detections) == 0 {
		return
	}

	msg := getLanguageDetectionMessage(detections)
	util.Send(s, m, msg)
}

func getLanguageDetectionMessage(detections []util.LanguageDetection) util.Message {
	var description bytes.Buffer

	for i, d := range detections {
		l := util.Languages[d.Language]

		r := "☐"
		if d.IsReliable {
			r = "☑︎"
		}

		ds := fmt.Sprintf("**%d**: %s [**%f**] %s | %s | %s\n",
			i+1, r, d.ConfidenceScore/10000, l.English, l.Russian, l.Native)
		description.WriteString(ds)
	}

	return util.Message{
		Title:       "Возможный язык:",
		Description: description.String(),
		FooterText:  "Используя API https://detectlanguage.com",
	}
}
