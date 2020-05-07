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

const asLogoURL = "https://www.freepnglogos.com/uploads/" +
	"app-store-logo-png/file-app-store-ios-custom-size-18.png"
const gpLogoURL = "https://www.freepnglogos.com/uploads/" +
	"google-play-png-logo/google-severs-music-studio-png-logo-21.png"

func Vitalina(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")

	for _, arg := range args {
		if util.IsAppID(arg) == util.AppStore {
			// Кроме как запросить данные с сервера нет возможности заранее
			// узнать тип (TODO изучить возможность)
			// TODO ограничить доступными диапазонами
			util.Send(s, m, GetMetadataMessage(arg, util.AppStore))
			util.Send(s, m, GetStoryMessage(arg))
			util.Send(s, m, GetRoomMessage(arg))
			// добавить Genre
			// добавить Store Front
			// добавить Language
			// добавить Platform
			// добавить Artist
		} else if util.IsAppID(arg) == util.GooglePlay {
			util.Send(s, m, GetMetadataMessage(arg, util.GooglePlay))
		} else if util.IsTimestamp(arg) {
			// GetTimestamp(arg) // FIXME
		} else if strings.HasPrefix(m.Content, "!") {
			util.Send(s, m, GetAsAppIDsMessage(m.Content[1:]))
			util.Send(s, m, GetGpAppIDsMessage(m.Content[1:]))
		}
	}

	if isPhrase(m.Content) {
		time.Sleep(100)
		s.MessageReactionAdd(m.ChannelID, m.ID, "➕")
		time.Sleep(100)
		s.MessageReactionAdd(m.ChannelID, m.ID, "➖")
	}

	if strings.HasPrefix(m.Content, ".") {
		detections := util.DetectLanguage(m.Content[1:])

		if len(detections) > 0 {
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
	}

	if strings.HasPrefix(m.Content, ",") {
		detections := util.DetectLanguage(m.Content[1:])

		if len(detections) > 0 {
			util.Send(s, m, GetLanguageDetectionMessage(detections))
		}
	}
}

func GetTimestampMessage(timestamp string) util.Message {
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
		value = future + strconv.Itoa(int(quantity)) + " " +
			getEnding("секунда", int(quantity)) + past + date
	} else if diff <= time.Second.Milliseconds() {
		quantity = diff * time.Minute.Milliseconds()
		value = future + strconv.Itoa(int(quantity)) + " " +
			getEnding("минута", int(quantity)) + past + date
	} else if diff <= time.Hour.Milliseconds()*24 {
		quantity = diff * time.Hour.Milliseconds()
		value = future + strconv.Itoa(int(quantity)) + " " +
			getEnding("час", int(quantity)) + past + date
	} else {
		quantity = diff * time.Hour.Milliseconds() * 24
		value = future + strconv.Itoa(int(quantity)) + " " +
			getEnding("день", int(quantity)) + past + date
	}

	return util.Message{
		Title:       "UNIX-time",
		Description: value,
	}
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

func GetMetadataMessage(appID string, store int) util.Message {
	location := "ru"
	language := "ru"

	var metadata scraper.MetadataResponse
	var link string
	var footerIconURL string
	if store == util.AppStore {
		metadata = scraper.AsMetadata(appID, location, language)
		link = metadata.Link
		footerIconURL = asLogoURL
	} else if store == util.GooglePlay {
		metadata = scraper.GpMetadata(appID, location, language)
		link = fmt.Sprintf("https://play.google.com/store/apps/details?id=%s", appID)
		footerIconURL = gpLogoURL
	}

	// TODO улучшить проверку
	if metadata.Title == "" {
		return util.Message{}
	}

	return util.Message{
		Title:        metadata.Title,
		Description:  metadata.Subtitle,
		Link:         link,
		ImageURL:     metadata.Screenshot1,
		ThumbnailURL: metadata.Logo,
		FooterText: fmt.Sprintf("Application\nAuthor: %s, %s", metadata.ArtistName,
			getStars(int(metadata.Rating))),
		FooterIconURL: footerIconURL,
	}
}

func GetStoryMessage(storyID string) util.Message {
	location := "ru"
	language := "ru"

	story := scraper.AsStory(storyID, location, language)

	// TODO улучшить проверку
	if story.ID == "" {
		return util.Message{}
	}

	iu := ""
	for _, v := range story.EditorialArtwork {
		iu = strings.Replace(v.URL, "{w}x{h}{c}.{f}", "512x512bb.png", -1)
	}

	ft := "\n"
	for _, card := range story.CardIds {
		ft += "\n" + card
	}

	return util.Message{
		Title: story.Label,
		Description: fmt.Sprintf("Story\n**%s**\n%s",
			story.EditorialNotes.Name, story.EditorialNotes.Short),
		Link:          story.Link.URL,
		ImageURL:      iu,
		FooterText:    ft,
		FooterIconURL: asLogoURL,
	}
}

func GetRoomMessage(adamID string) util.Message {
	location := "ru"
	language := "ru"

	room := scraper.AsRoom(adamID, location, language)

	// TODO улучшить проверку
	if room.Title == "" {
		return util.Message{}
	}

	return util.Message{
		Title:        room.Title,
		Description:  room.Subtitle,
		Link:         room.Link,
		ImageURL:     room.Screenshot1,
		ThumbnailURL: room.Logo,
		FooterText: fmt.Sprintf("Room\nAuthor: %s, %s",
			room.ArtistName, getStars(int(room.Rating))),
		FooterIconURL: asLogoURL,
	}
}

func GetAsAppIDsMessage(keyword string) util.Message {
	location := "ru"
	language := "ru"

	var d bytes.Buffer

	metadatas := scraper.AsAppIDs(keyword, location, language)
	for i, m := range metadatas {
		d.WriteString(fmt.Sprintf("**%d**: %s (`%s`) %s\n",
			i+1, m.Title, m.AppID, getStars(int(m.Rating))))
	}

	return util.Message{
		Title: fmt.Sprintf("Приложения App Store по ключевому слову «%s»:",
			keyword),
		Description:   d.String(),
		ThumbnailURL:  asLogoURL,
		FooterText:    fmt.Sprintf("Total: %d", len(metadatas)),
		FooterIconURL: asLogoURL,
	}
}

func GetGpAppIDsMessage(keyword string) util.Message {
	location := "ru"
	language := "ru"

	var d bytes.Buffer

	metadatas := scraper.GpAppIDs(keyword, location, language)
	for i, m := range metadatas {
		d.WriteString(fmt.Sprintf("**%d**: %s (`%s`) %s\n",
			i+1, m.Title, m.AppID, getStars(int(m.Rating))))
	}

	return util.Message{
		Title: fmt.Sprintf("Приложения Google Play по ключевому слову «%s»:",
			keyword),
		Description: d.String(),
		Link: fmt.Sprintf("https://play.google.com/store/search?q=%s&c=apps&gl=%s&hl=%s",
			keyword, location, language),
		ThumbnailURL:  gpLogoURL,
		FooterText:    fmt.Sprintf("Total: %d", len(metadatas)),
		FooterIconURL: gpLogoURL,
	}
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

func GetLanguageDetectionMessage(detections []util.LanguageDetection) util.Message {
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
