package command

import (
	"bytes"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/salaleser/vitalina/http"
	"github.com/salaleser/vitalina/parser"
	"github.com/salaleser/vitalina/util"
)

func Vitalina(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")

	for _, arg := range args {
		if util.IsAppID(arg) == util.AppStore {
			SendMetadata(s, m, arg, util.AppStore)
		} else if util.IsAppID(arg) == util.GooglePlay {
			SendMetadata(s, m, arg, util.GooglePlay)
		} else if util.IsTimestamp(arg) {
			// SendTimestamp(s, m, arg)
		} else if strings.HasPrefix(m.Content, "!") {
			SendAsAppIDs(s, m, m.Content[1:])
			SendGpAppIDs(s, m, m.Content[1:])
		}
	}

	if isPhrase(m.Content) {
		time.Sleep(100)
		s.MessageReactionAdd(m.ChannelID, m.ID, "➕")
		time.Sleep(100)
		s.MessageReactionAdd(m.ChannelID, m.ID, "➖")
	}
}

func SendTimestamp(s *discordgo.Session, m *discordgo.MessageCreate, timestamp string) {
	ts, _ := strconv.Atoi(timestamp)
	date := time.Now().Format("EEEE, dd MMMM YYYY года в HH:mm:ss")
	diff := time.Now().Unix() - int64(ts*1000)
	future := ""
	past := " назад, "
	if diff < 0 {
		future = "через "
		past = ", "
		diff = int64(math.Abs(float64(diff)))
	}

	var quantity int64
	var value string
	if diff <= time.Second.Milliseconds() {
		quantity = diff
		value = future + strconv.Itoa(int(quantity)) + " " + getEnding("секунда", int(quantity)) + past + date
	} else if diff <= time.Second.Milliseconds() {
		quantity = diff * time.Minute.Milliseconds()
		value = future + strconv.Itoa(int(quantity)) + " " + getEnding("минута", int(quantity)) + past + date
	} else if diff <= time.Hour.Milliseconds()*24 {
		quantity = diff * time.Hour.Milliseconds()
		value = future + strconv.Itoa(int(quantity)) + " " + getEnding("час", int(quantity)) + past + date
	} else {
		quantity = diff * time.Hour.Milliseconds() * 24
		value = future + strconv.Itoa(int(quantity)) + " " + getEnding("день", int(quantity)) + past + date
	}

	footer := discordgo.MessageEmbedFooter{
		Text: "Не тормози мужик, моя информация может тебе пригодиться.",
		// IconURL: aso.Logo,
	}

	thumbnail := discordgo.MessageEmbedThumbnail{
		// URL: aso.Logo,
	}

	image := discordgo.MessageEmbedImage{
		// URL: aso.Screenshot1,
	}

	embed := discordgo.MessageEmbed{
		Color:       1200,
		Title:       "UNIX-time",
		Description: value,
		Footer:      &footer,
		// URL:         link,
		Thumbnail: &thumbnail,
		Image:     &image,
	}

	message := discordgo.MessageSend{
		Embed: &embed,
	}

	s.ChannelMessageSendComplex(m.ChannelID, &message)
}

func getEnding(nominative string, quantity int) string {
	genitive, _ := regexp.Compile("^\\d*[234]$")
	plural, _ := regexp.Compile("^\\d*[05-9]$|^\\d*1\\d$")
	genetiveMatcher := genitive.MatchString(strconv.Itoa(quantity))
	pluralMatcher := plural.MatchString(strconv.Itoa(quantity))
	switch nominative {
	case "день":
		if pluralMatcher {
			return "дней"
		}
		if genetiveMatcher {
			return "дня"
		}
	case "час":
		if pluralMatcher {
			return "часов"
		}
		if genetiveMatcher {
			return "часа"
		}
	case "минута":
		if pluralMatcher {
			return "минут"
		}
		if genetiveMatcher {
			return "минуты"
		}
	case "секунда":
		if pluralMatcher {
			return "секунд"
		}
		if genetiveMatcher {
			return "секунды"
		}
	}
	return nominative
}

func SendMetadata(s *discordgo.Session, m *discordgo.MessageCreate, appID string, store int) {
	location := "ru"
	language := "ru"

	var body []byte
	var metadata parser.Metadata
	var link string
	var footerText string
	var footerIconURL string
	if store == util.AppStore {
		body = http.GetAsMetadataBody(appID, location, language)
		metadata = parser.ParseAsAsoBody(body)
		link = "https://apps.apple.com/ru/app/id" + appID + "?l=ru"
		footerText = "App Store"
		footerIconURL = "https://www.freepnglogos.com/uploads/app-store-logo-png/file-app-store-ios-custom-size-18.png"
	} else if store == util.GooglePlay {
		body = http.GetGpMetadataBody(appID, location, language)
		metadata = parser.ParseGpAsoBody(body)
		link = "https://play.google.com/store/apps/details?id=" + appID
		footerText = "Google Play"
		footerIconURL = "https://www.freepnglogos.com/uploads/google-play-png-logo/google-severs-music-studio-png-logo-21.png"
	}

	if metadata.Title == "" {
		s.ChannelMessageSend(m.ChannelID, "_Не смогла..._")
		return
	}

	footer := discordgo.MessageEmbedFooter{
		Text:    footerText,
		IconURL: footerIconURL,
	}

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: metadata.Logo,
	}

	image := discordgo.MessageEmbedImage{
		URL: metadata.Screenshot1,
	}

	embed := discordgo.MessageEmbed{
		Color:       800,
		Title:       metadata.Title,
		Description: metadata.Subtitle,
		Footer:      &footer,
		URL:         link,
		Thumbnail:   &thumbnail,
		Image:       &image,
	}

	message := discordgo.MessageSend{
		Embed: &embed,
	}

	s.ChannelMessageSendComplex(m.ChannelID, &message)
}

func SendAsAppIDs(s *discordgo.Session, m *discordgo.MessageCreate, keyword string) {
	location := "ru"
	language := "ru"
	count := 5

	var buffer bytes.Buffer

	// TODO encapsulate
	asBody := http.AsAppIDs(keyword, location, language, count)
	asAppIDs := parser.ParseAsIDsBody(asBody, count)
	if len(asAppIDs) == 0 || asAppIDs[0].Name == "" {
		s.ChannelMessageSend(m.ChannelID, "_Не смогла в App Store..._")
		return
	}

	for i := 0; i < count; i++ {
		buffer.WriteString(fmt.Sprintf("*%d*: %s (`%s`)\n", i+1, asAppIDs[i].Name, asAppIDs[i].ID))
	}

	footer := discordgo.MessageEmbedFooter{
		Text:    "App Store",
		IconURL: "https://www.freepnglogos.com/uploads/app-store-logo-png/file-app-store-ios-custom-size-18.png",
	}

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: "https://www.freepnglogos.com/uploads/app-store-logo-png/file-app-store-ios-custom-size-18.png",
	}

	image := discordgo.MessageEmbedImage{
		// URL: "https://developer.apple.com/app-store/marketing/guidelines/images/badge-download-on-the-app-store.svg",
	}

	embed := discordgo.MessageEmbed{
		Color:       500,
		Title:       "Приложения App Store по ключевому слову «" + keyword + "»:",
		Description: buffer.String(),
		Footer:      &footer,
		// URL:         link,
		Thumbnail: &thumbnail,
		Image:     &image,
	}

	message := discordgo.MessageSend{
		Embed: &embed,
	}

	s.ChannelMessageSendComplex(m.ChannelID, &message)
}

func SendGpAppIDs(s *discordgo.Session, m *discordgo.MessageCreate, keyword string) {
	location := "ru"
	language := "ru"
	count := 5

	var buffer bytes.Buffer

	// TODO encapsulate
	gpBody := http.GpAppIDs(keyword, location, language, count)
	gpAppIDs := parser.ParseGpIDsBody(gpBody, count)
	if len(gpAppIDs) == 0 || gpAppIDs[0].Title == "" {
		s.ChannelMessageSend(m.ChannelID, "_Не смогла в Google Play..._")
		return
	}

	for i := 0; i < count; i++ {
		buffer.WriteString(fmt.Sprintf("*%d*: %s (`%s`)\n", i+1, gpAppIDs[i].Title, gpAppIDs[i].AppID))
	}

	footer := discordgo.MessageEmbedFooter{
		Text:    "Google Play",
		IconURL: "https://www.freepnglogos.com/uploads/google-play-png-logo/google-severs-music-studio-png-logo-21.png",
	}

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: "https://www.freepnglogos.com/uploads/google-play-png-logo/google-severs-music-studio-png-logo-21.png",
	}

	image := discordgo.MessageEmbedImage{
		// URL: "https://play.google.com/intl/en_us/badges/static/images/badges/en_badge_web_generic.png",
	}

	embed := discordgo.MessageEmbed{
		Color:       500,
		Title:       "Приложения Google Play по ключевому слову «" + keyword + "»:",
		Description: buffer.String(),
		Footer:      &footer,
		// URL:         link,
		Thumbnail: &thumbnail,
		Image:     &image,
	}

	message := discordgo.MessageSend{
		Embed: &embed,
	}

	s.ChannelMessageSendComplex(m.ChannelID, &message)
}

func isPhrase(s string) bool {
	regexp, _ := regexp.Compile("^.{1,}\\?$")

	return regexp.MatchString(s)
}
