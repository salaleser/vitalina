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
			language := util.GetLanguageByCode(detections[0].Language)
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
	}

	if strings.HasPrefix(m.Content, ",") {
		detections := util.DetectLanguage(m.Content[1:])

		if len(detections) > 0 {
			SendLanguageDetection(s, m, detections)
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

	ft := "Author: " + metadata.ArtistName + ", " + getStars(int(metadata.Rating))

	util.Send(s, m, metadata.Title, metadata.Subtitle, link, metadata.Screenshot1, metadata.Logo, ft, footerIconURL)
}

func SendStory(s *discordgo.Session, m *discordgo.MessageCreate, storyID string) {
	location := "ru"
	language := "ru"

	story := scraper.AsStory(storyID, location, language)

	// TODO улучшить проверку
	if story.ID == "" {
		return
	}

	iu := ""
	for _, v := range story.EditorialArtwork {
		iu = strings.Replace(v.URL, "{w}x{h}{c}.{f}", "512x512bb.png", -1)
	}

	t := story.Label
	d := "**" + story.EditorialNotes.Name + "**\n" + story.EditorialNotes.Short
	l := story.Link.URL

	ft := "\n"
	for _, card := range story.CardIds {
		ft += "\n" + card
	}

	util.Send(s, m, t, d, l, iu, "", ft, asLogoURL)
}

func SendAsAppIDs(s *discordgo.Session, m *discordgo.MessageCreate, keyword string) {
	location := "ru"
	language := "ru"
	count := 5

	var d bytes.Buffer

	metadatas := scraper.AsAppIDs(keyword, location, language)
	for i, m := range metadatas {
		d.WriteString(fmt.Sprintf("**%d**: %s (`%s`) %s\n", i+1, m.Title, m.AppID, getStars(int(m.Rating))))
		if i >= count {
			break
		}
	}

	t := "Приложения App Store по ключевому слову «" + keyword + "»:"
	ft := fmt.Sprintf("Total: %d", len(metadatas))

	util.Send(s, m, t, d.String(), "", "", asLogoURL, ft, asLogoURL)
}

func SendGpAppIDs(s *discordgo.Session, m *discordgo.MessageCreate, keyword string) {
	location := "ru"
	language := "ru"
	count := 5

	var d bytes.Buffer

	metadatas := scraper.GpAppIDs(keyword, location, language)
	for i, m := range metadatas {
		d.WriteString(fmt.Sprintf("**%d**: %s (`%s`) %s\n", i+1, m.Title, m.AppID, getStars(int(m.Rating))))
		if i >= count {
			break
		}
	}

	t := "Приложения Google Play по ключевому слову «" + keyword + "»:"
	l := "https://play.google.com/store/search?q=" + keyword + "&c=apps&gl=" + location + "&hl=" + language
	ft := fmt.Sprintf("Total: %d", len(metadatas))

	util.Send(s, m, t, d.String(), l, "", gpLogoURL, ft, gpLogoURL)
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

func SendLanguageDetection(s *discordgo.Session, m *discordgo.MessageCreate, detections []util.LanguageDetection) {
	var description bytes.Buffer

	for i, d := range detections {
		l := util.GetLanguageByCode(d.Language)

		r := "☐"
		if d.IsReliable {
			r = "☑︎"
		}

		ds := fmt.Sprintf("**%d**: %s [**%f**] %s | %s | %s\n", i+1, r, d.ConfidenceScore/10000, l.English, l.Russian, l.Native)
		description.WriteString(ds)
	}

	t := "Возможный язык:"
	ft := "Используя API https://detectlanguage.com"

	util.Send(s, m, t, description.String(), "", "", "", ft, "")
}
