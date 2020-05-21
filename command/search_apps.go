package command

import (
	"bytes"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/salaleser/scraper"
	"github.com/salaleser/vitalina/util"
)

// SearchApps sends message with store applications.
func SearchApps(s *discordgo.Session, m *discordgo.MessageCreate) {
	cc := "us"
	l := "en-us"

	msg1 := getAsAppIDsMessage(m.Content[1:], cc, l)
	util.Send(s, m, msg1)

	msg2 := getGpAppIDsMessage(m.Content[1:], cc, l)
	util.Send(s, m, msg2)
}

func getAsAppIDsMessage(keyword string, cc string, l string) util.Message {
	var d bytes.Buffer

	metadatas := scraper.AsAppIDs(keyword, cc, l)
	for i, m := range metadatas {
		d.WriteString(fmt.Sprintf("**%d**: %s (`%s`) %s\n",
			i+1, m.Title, m.AppID, util.GetStarsBar(int(m.Rating))))
	}

	return util.Message{
		Title: fmt.Sprintf("Приложения App Store по ключевому слову «%s»:",
			keyword),
		Description:   d.String(),
		ThumbnailURL:  util.AsLogoURL,
		FooterText:    fmt.Sprintf("Total: %d", len(metadatas)),
		FooterIconURL: util.AsLogoURL,
	}
}

func getGpAppIDsMessage(keyword string, cc string, l string) util.Message {
	var d bytes.Buffer

	metadatas := scraper.GpAppIDs(keyword, cc, l)
	for i, m := range metadatas {
		d.WriteString(fmt.Sprintf("**%d**: %s (`%s`) %s\n",
			i+1, m.Title, m.AppID, util.GetStarsBar(int(m.Rating))))
	}

	return util.Message{
		Title: fmt.Sprintf("Приложения Google Play по ключевому слову «%s»:",
			keyword),
		Description: d.String(),
		Link: fmt.Sprintf("https://play.google.com/store/search?q=%s&c=apps&gl=%s&hl=%s",
			keyword, cc, l),
		ThumbnailURL:  util.GpLogoURL,
		FooterText:    fmt.Sprintf("Total: %d", len(metadatas)),
		FooterIconURL: util.GpLogoURL,
	}
}
