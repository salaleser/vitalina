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
	"github.com/salaleser/scraper"
	"github.com/salaleser/vitalina/util"
)

const asLogoURL = "https://www.freepnglogos.com/uploads/app-store-logo-png/file-app-store-ios-custom-size-18.png"
const gpLogoURL = "https://www.freepnglogos.com/uploads/google-play-png-logo/google-severs-music-studio-png-logo-21.png"

func Vitalina(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")

	for _, arg := range args {
		if util.IsAppID(arg) == util.AppStore {
			SendMetadata(s, m, arg, util.AppStore)
			SendStory(s, m, arg)
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

	if time.Now().Second()%20 == 0 {
		time.Sleep(100)
		s.MessageReactionAdd(m.ChannelID, m.ID, "🙂")
	}

	if strings.HasPrefix(m.Content, ".") {
		detections := util.DetectLanguage(m.Content[1:])

		if len(detections) > 0 {
			flag := util.GetFlagByLanguage(detections[0].Language)
			time.Sleep(100)
			s.MessageReactionAdd(m.ChannelID, m.ID, flag)

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
		Color:       500000,
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

	var metadata scraper.Metadata
	var link string
	var footerIconURL string
	if store == util.AppStore {
		metadata = scraper.AsMetadataBody(appID, location, language)
		link = metadata.Link
		footerIconURL = asLogoURL
	} else if store == util.GooglePlay {
		metadata = scraper.GpMetadataBody(appID, location, language)
		link = "https://play.google.com/store/apps/details?id=" + appID
		footerIconURL = gpLogoURL
	}

	// TODO улучшить проверку
	if metadata.Title == "" {
		return
	}

	footer := discordgo.MessageEmbedFooter{
		Text:    "Author: " + metadata.ArtistName + ", " + getStars(int(metadata.Rating)),
		IconURL: footerIconURL,
	}

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: metadata.Logo,
	}

	image := discordgo.MessageEmbedImage{
		URL: metadata.Screenshot1,
	}

	embed := discordgo.MessageEmbed{
		Color:       100000,
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

func SendStory(s *discordgo.Session, m *discordgo.MessageCreate, storyID string) {
	location := "ru"
	language := "ru"

	story := scraper.AsStory(storyID, location, language)

	// TODO улучшить проверку
	if story.ID == "" {
		return
	}

	imageURL := ""
	for _, v := range story.EditorialArtwork {
		imageURL = strings.Replace(v.URL, "{w}x{h}{c}.{f}", "512x512bb.png", -1)
	}

	cardIDs := "\n"
	for _, card := range story.CardIds {
		cardIDs += "\n" + card
	}

	footer := discordgo.MessageEmbedFooter{
		Text:    cardIDs,
		IconURL: asLogoURL,
	}

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: "",
	}

	image := discordgo.MessageEmbedImage{
		URL: imageURL,
	}

	embed := discordgo.MessageEmbed{
		Color:       100000,
		Title:       story.Label,
		Description: "**" + story.EditorialNotes.Name + "**\n" + story.EditorialNotes.Short,
		Footer:      &footer,
		URL:         story.Link.URL,
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

	metadatas := scraper.AsAppIDs(keyword, location, language)
	for i, m := range metadatas {
		buffer.WriteString(fmt.Sprintf("**%d**: %s (`%s`) %s\n", i+1, m.Title, m.AppID, getStars(int(m.Rating))))
		if i >= count {
			break
		}
	}

	footer := discordgo.MessageEmbedFooter{
		Text:    fmt.Sprintf("Total: %d", len(metadatas)),
		IconURL: asLogoURL,
	}

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: asLogoURL,
	}

	image := discordgo.MessageEmbedImage{
		// URL: "https://developer.apple.com/app-store/marketing/guidelines/images/badge-download-on-the-app-store.svg",
	}

	embed := discordgo.MessageEmbed{
		Color:       100000,
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

	metadatas := scraper.GpAppIDs(keyword, location, language)
	for i, m := range metadatas {
		buffer.WriteString(fmt.Sprintf("**%d**: %s (`%s`) %s\n", i+1, m.Title, m.AppID, getStars(int(m.Rating))))
		if i > count {
			break
		}
	}

	footer := discordgo.MessageEmbedFooter{
		Text:    fmt.Sprintf("Total: %d", len(metadatas)),
		IconURL: gpLogoURL,
	}

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: gpLogoURL,
	}

	image := discordgo.MessageEmbedImage{
		// URL: "https://play.google.com/intl/en_us/badges/static/images/badges/en_badge_web_generic.png",
	}

	embed := discordgo.MessageEmbed{
		Color:       100000,
		Title:       "Приложения Google Play по ключевому слову «" + keyword + "»:",
		Description: buffer.String(),
		Footer:      &footer,
		URL:         "https://play.google.com/store/search?q=" + keyword + "&c=apps&gl=" + location + "&hl=" + language,
		Thumbnail:   &thumbnail,
		Image:       &image,
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

func getStars(value int) string {
	switch value {
	case 0:
		return "☆☆☆☆☆"
	case 1:
		return "★☆☆☆☆"
	case 2:
		return "★★☆☆☆"
	case 3:
		return "★★★☆☆"
	case 4:
		return "★★★★☆"
	case 5:
		return "★★★★★"
	default:
		return "—"
	}
}
