package command

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	as "github.com/salaleser/appstoreapi"
	gp "github.com/salaleser/googleplayapi"

	"github.com/bwmarrin/discordgo"
	"github.com/salaleser/vitalina/util"
)

// SearchApps sends message with store applications.
func SearchApps(s *discordgo.Session, m *discordgo.MessageCreate) {
	cc := "us"
	l := ""
	gl := "us"
	hl := ""

	content := m.Content[2:]

	args := strings.Split(content, " ")

	// Scan for country code and language
	for _, arg := range args {
		if _, ok := as.StoreFronts[arg]; ok {
			util.Debug(fmt.Sprintf(""+
				"Country code %q detected!",
				arg,
			))
			cc = arg
			gl = util.ToGooglePlayGeoLocation(arg)
			country := util.Countries[strings.ToLower(arg)]
			s.MessageReactionAdd(m.ChannelID, m.ID, country.Emoji)
		}

		if _, ok := as.Languages[arg]; ok {
			util.Debug(fmt.Sprintf(""+
				"Language %q detected!",
				arg,
			))
			l = arg
			hl = util.ToGooglePlayHumanLanguage(arg)
			language := util.Languages[strings.Split(arg, "-")[0]]
			s.MessageReactionAdd(m.ChannelID, m.ID, language.Emoji)
		}
	}

	msg1 := getAsAppIDsMessage(args[0], cc, l)
	util.Send(s, m, msg1)

	msg2 := getGpAppIDsMessage(args[0], gl, hl)
	util.Send(s, m, msg2)
}

func getAsAppIDsMessage(keyword string, cc string, l string) util.Message {
	var d bytes.Buffer

	data, err := as.AppIDs(keyword, cc, l)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error as app-ids: %v", err)
	}

	for i, m := range data {
		d.WriteString(fmt.Sprintf("**%d**: %s (`%s`) %s\n",
			i+1, m.Title, m.AppID, util.GetStarsBar(int(m.Rating))))
	}

	return util.Message{
		Title: fmt.Sprintf("Приложения App Store по ключевому слову «%s»:",
			keyword),
		Description:   d.String(),
		ThumbnailURL:  util.AsLogoURL,
		FooterText:    fmt.Sprintf("Total: %d", len(data)),
		FooterIconURL: util.AsLogoURL,
	}
}

func getGpAppIDsMessage(keyword string, gl string, hl string) util.Message {
	var d bytes.Buffer

	data, err := gp.AppIDs(keyword, gl, hl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error gp app-ids: %v", err)
	}

	for i, m := range data {
		d.WriteString(fmt.Sprintf("**%d**: %s (`%s`) %s\n",
			i+1, m.Title, m.AppID, util.GetStarsBar(int(m.Rating))))
	}

	return util.Message{
		Title: fmt.Sprintf("Приложения Google Play по ключевому слову «%s»:",
			keyword),
		Description: d.String(),
		Link: fmt.Sprintf("https://play.google.com/store/search?q=%s&c=apps&gl=%s&hl=%s",
			keyword, gl, hl),
		ThumbnailURL:  util.GpLogoURL,
		FooterText:    fmt.Sprintf("Total: %d", len(data)),
		FooterIconURL: util.GpLogoURL,
	}
}
