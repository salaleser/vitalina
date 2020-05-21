package command

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/salaleser/scraper"
	"github.com/salaleser/vitalina/util"
)

var (
	force = false // принудительно сообщать о ходе работы
	cc    = "us"
	l     = ""
)

// Vitalina is a AI wannabe.
func Vitalina(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := m.Content

	if strings.HasPrefix(m.Content, "?") {
		force = true
		content = m.Content[1:]
	}

	args := strings.Split(content, " ")

	// Scan for country code and language
	for _, arg := range args {
		if _, ok := scraper.StoreFronts[strings.ToUpper(arg)]; ok {
			util.Debug(fmt.Sprintf("Country code %q detected!", arg))
			cc = arg
			flag := util.GetFlagByCountryCode(arg)
			s.MessageReactionAdd(m.ChannelID, m.ID, flag)
		}

		if _, ok := scraper.Languages[arg]; ok {
			util.Debug(fmt.Sprintf("Language %q detected!", arg))
			l = arg
			isoLanguageCode := strings.Split(arg, "-")[0]
			language := util.Languages[isoLanguageCode]
			s.MessageReactionAdd(m.ChannelID, m.ID, language.Emoji)
		}
	}

	// Scan for IDs
	for _, arg := range args {
		// FIXME коды в верхнем регистре хранятся в мапе
		if util.ContainsMap(scraper.StoreFronts, arg) {
			cc := util.GetCcByStoreFront(arg)
			msg := getStoreFrontMessage(arg, cc)
			util.Send(s, m, msg)
		}

		if asLanguageCode, ok := scraper.Languages[arg]; ok {
			msg := getAsLanguageMessage(asLanguageCode, arg)
			util.Send(s, m, msg)
		}

		if util.MatchesAsAppID(arg) {
			msg := getMetadataMessage(arg, util.AppStore, cc, l)
			util.Send(s, m, msg)
		}

		if util.MatchesAsStoryID(arg) {
			msg := getStoryMessage(arg, cc, l)
			util.Send(s, m, msg)
		}

		if util.MatchesAsRoomID(arg) {
			msg := getRoomMessage(arg, cc, l)
			util.Send(s, m, msg)
		}

		if util.MatchesAsGroupingID(arg) {
			id, _ := strconv.Atoi(arg)
			msg := getGroupingMessage(id, cc, l)
			util.Send(s, m, msg)
		}

		if util.MatchesAsGenreID(arg) {
			id, _ := strconv.Atoi(arg)
			msg := getAsGenreMessage(id, cc, l)
			util.Send(s, m, msg)
		}

		if util.GetStoreFromAppID(arg) == util.GooglePlay {
			msg := getMetadataMessage(arg, util.GooglePlay, cc, l)
			util.Send(s, m, msg)
		}
	}

	// Scan for phrases
	if isPhrase(m.Content) {
		time.Sleep(100)
		s.MessageReactionAdd(m.ChannelID, m.ID, "➕")
		time.Sleep(100)
		s.MessageReactionAdd(m.ChannelID, m.ID, "➖")
	}
}

// func getTimestampMessage(timestamp string) util.Message {
// 	ts, _ := strconv.Atoi(timestamp)
// 	date := time.Now().Format("EEEE, dd MMMM YYYY года в HH:mm:ss")
// 	diff := time.Now().Unix() - int64(ts*1000)
// 	future := ""
// 	past := " назад, "
// 	if diff < 0 {
// 		future = "через "
// 		past = ", "
// 		diff = int64(math.Abs(float64(diff)))
// 	}

// 	var quantity int64
// 	var value string
// 	if diff <= time.Second.Milliseconds() {
// 		quantity = diff
// 		value = future + strconv.Itoa(int(quantity)) + " " +
// 			getEnding("секунда", int(quantity)) + past + date
// 	} else if diff <= time.Second.Milliseconds() {
// 		quantity = diff * time.Minute.Milliseconds()
// 		value = future + strconv.Itoa(int(quantity)) + " " +
// 			getEnding("минута", int(quantity)) + past + date
// 	} else if diff <= time.Hour.Milliseconds()*24 {
// 		quantity = diff * time.Hour.Milliseconds()
// 		value = future + strconv.Itoa(int(quantity)) + " " +
// 			getEnding("час", int(quantity)) + past + date
// 	} else {
// 		quantity = diff * time.Hour.Milliseconds() * 24
// 		value = future + strconv.Itoa(int(quantity)) + " " +
// 			getEnding("день", int(quantity)) + past + date
// 	}

// 	return util.Message{
// 		Title:       "UNIX-time",
// 		Description: value,
// 	}
// }

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

func getMetadataMessage(appID string, store int, cc string, l string) util.Message {
	var metadata scraper.MetadataResponse
	var link string
	var footerIconURL string
	if store == util.AppStore {
		metadata = scraper.AsMetadata(appID, cc, l)
		link = metadata.Link
		footerIconURL = util.AsLogoURL
	} else if store == util.GooglePlay {
		metadata = scraper.GpMetadata(appID, cc, l)
		link = fmt.Sprintf("https://play.google.com/store/apps/details?id=%s", appID)
		footerIconURL = util.GpLogoURL
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
			util.GetStarsBar(int(metadata.Rating))),
		FooterIconURL: footerIconURL,
	}
}

func getStoryMessage(storyID string, cc string, l string) util.Message {
	story := scraper.AsStory(storyID, cc, l)

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
		FooterIconURL: util.AsLogoURL,
	}
}

func getAsGenreMessage(id int, cc string, l string) util.Message {
	page, err := scraper.AsGenre(id, cc)
	if err != nil {
		util.Debug(err.Error())
		return util.Message{}
	}

	genre := strings.Split(page.PageData.MetricsBase.PageDetails, "_")[0]

	return util.Message{
		Title:         fmt.Sprintf("App Store Genre detected by code «%d»", id),
		Description:   genre,
		Link:          fmt.Sprintf("https://itunes.apple.com/%s/genre?id=%d", cc, id),
		FooterText:    fmt.Sprintf("%d=%s", id, genre),
		FooterIconURL: util.AsLogoURL,
	}
}

func getStoreFrontMessage(sf string, cc string) util.Message {
	return util.Message{
		Title: fmt.Sprintf("App Store Store Front detected by code «%s»",
			sf),
		Description:   cc + " | " + util.GetFlagByCountryCode(cc),
		FooterText:    fmt.Sprintf("%s=%s", sf, cc),
		FooterIconURL: util.AsLogoURL,
	}
}

func getAsLanguageMessage(asLanguageCode string, l string) util.Message {
	return util.Message{
		Title: fmt.Sprintf("App Store Langauge detected by code «%s»",
			asLanguageCode),
		Description:   l,
		FooterText:    fmt.Sprintf("%s=%s", l, asLanguageCode),
		FooterIconURL: util.AsLogoURL,
	}
}

func getGroupingMessage(groupingID int, cc string, l string) util.Message {
	page, err := scraper.AsGrouping(groupingID, cc, l)
	if err != nil {
		util.Debug(fmt.Sprintf("command.GetGroupingMessage(%d,%s,%s): %v",
			groupingID, cc, l, err))

		// TODO отправлять сообщение об ошибке
		if force {
			return util.Message{
				Title: "[ERR]",
				Description: fmt.Sprintf("command.GetGroupingMessage(%d,%s,%s): %v",
					groupingID, cc, l, err),
				FooterText: err.Error(),
			}
		}

		return util.Message{}
	}

	grouping := page.PageData.FcStructure.Model

	return util.Message{
		Title:       grouping.AdamID,
		Description: grouping.FcKind,
		// Link:          story.Link.URL,
		// ImageURL:      iu,
		// FooterText:    ft,
		FooterIconURL: util.AsLogoURL,
	}
}

func getRoomMessage(adamID string, cc string, l string) util.Message {
	room := scraper.AsRoom(adamID, cc, l)

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
			room.ArtistName, util.GetStarsBar(int(room.Rating))),
		FooterIconURL: util.AsLogoURL,
	}
}

func isPhrase(s string) bool {
	regexp, _ := regexp.Compile("^.{1,}\\?$")

	return regexp.MatchString(s)
}
