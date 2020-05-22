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

type room struct {
	adamID             string
	imageURL           string
	contentIDs         []string
	designBadge        string
	designTag          string
	displayStyle       string
	doNotFilter        bool
	fcKind             string
	name               string
	seeAllURL          string
	shouldUseGradients bool
	suppressTagline    string
	tagline            string
	title              string
}

// Vitalina is a AI wannabe.
func Vitalina(s *discordgo.Session, m *discordgo.MessageCreate) {
	force := false // принудительно сообщать о ходе работы
	cc := "us"
	l := ""
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
			country := util.Countries[arg]
			s.MessageReactionAdd(m.ChannelID, m.ID, country.Emoji)
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
		id, _ := strconv.Atoi(arg)
		// FIXME коды в верхнем регистре хранятся в мапе
		if util.ContainsMap(scraper.StoreFronts, id) {
			cc := util.GetCcByStoreFront(id)
			msg := getStoreFrontMessage(id, cc)
			util.Send(s, m, msg)
		}

		if force {
			if asLanguageCode, ok := scraper.Languages[arg]; ok {
				msg := getAsLanguageMessage(asLanguageCode, arg)
				util.Send(s, m, msg)
			}
		}

		if util.MatchesAsAppID(arg) {
			msg := getMetadataMessage(arg, util.AppStore, cc, l)
			util.Send(s, m, msg)
		}

		if util.MatchesAsStoryID(arg) {
			id, _ := strconv.Atoi(arg)
			err := processStory(s, m, id, cc, l)
			if err != nil {
				util.Debug(fmt.Sprintf("story [id=%d,cc=%s,l=%s]: %v", id, cc, l,
					err))
				if force {
					util.SendError(s, m,
						fmt.Sprintf("Story [id=%d,cc=%s,l=%s]", id, cc, l),
						err,
					)
				}
			}
		}

		if util.MatchesAsRoomID(arg) {
			id, _ := strconv.Atoi(arg)
			err := processRoom(s, m, id, cc, l)
			if err != nil {
				util.Debug(fmt.Sprintf("room [id=%d,cc=%s,l=%s]: %v", id, cc, l,
					err))
				if force {
					util.SendError(s, m,
						fmt.Sprintf("Room [id=%d,cc=%s,l=%s]", id, cc, l),
						err,
					)
				}
			}
		}

		if util.MatchesAsGroupingID(arg) {
			id, _ := strconv.Atoi(arg)
			err := processGrouping(s, m, id, cc, l)
			if err != nil {
				util.Debug(fmt.Sprintf("grouping [id=%d,cc=%s,l=%s]: %v", id, cc, l,
					err))
				if force {
					util.SendError(s, m,
						fmt.Sprintf("Grouping [id=%d,cc=%s,l=%s]", id, cc, l),
						err,
					)
				}
			}
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

// func getEnding(nominative string, quantity int) string {
// 	genitive, _ := regexp.Compile("^\\d*[234]$")
// 	plural, _ := regexp.Compile("^\\d*[05-9]$|^\\d*1\\d$")
// 	genetiveMatcher := genitive.MatchString(strconv.Itoa(quantity))
// 	pluralMatcher := plural.MatchString(strconv.Itoa(quantity))
// 	switch nominative {
// 	case "день":
// 		if pluralMatcher {
// 			return "дней"
// 		}
// 		if genetiveMatcher {
// 			return "дня"
// 		}
// 	case "час":
// 		if pluralMatcher {
// 			return "часов"
// 		}
// 		if genetiveMatcher {
// 			return "часа"
// 		}
// 	case "минута":
// 		if pluralMatcher {
// 			return "минут"
// 		}
// 		if genetiveMatcher {
// 			return "минуты"
// 		}
// 	case "секунда":
// 		if pluralMatcher {
// 			return "секунд"
// 		}
// 		if genetiveMatcher {
// 			return "секунды"
// 		}
// 	}
// 	return nominative
// }

func getMetadataMessage(appID string, store int, cc string, l string) util.Message {
	var metadata scraper.MetadataResponse
	var link string
	var footerIconURL string
	if store == util.AppStore {
		metadata = scraper.Metadata(appID, cc, l)
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

func getAsGenreMessage(id int, cc string, l string) util.Message {
	page, err := scraper.Genre(id, cc)
	if err != nil {
		util.Debug(err.Error())
		return util.Message{}
	}

	genre := strings.Split(page.PageData.MetricsBase.PageDetails, "_")[0]

	return util.Message{
		Title:       fmt.Sprintf("App Store Genre detected by code «%d»", id),
		Description: genre,
		Link: fmt.Sprintf("https://itunes.apple.com/%s/genre?id=%d",
			cc, id),
		FooterText:    fmt.Sprintf("%d=%s", id, genre),
		FooterIconURL: util.AsLogoURL,
	}
}

func getStoreFrontMessage(sf int, cc string) util.Message {
	return util.Message{
		// FIXME перепутано там что-то
		Title: fmt.Sprintf("App Store Store Front detected by code «%d»",
			sf),
		Description:   cc + " " + util.GetFlagByCountryCode(cc),
		FooterText:    fmt.Sprintf("%d=%s", sf, cc),
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

func processGrouping(s *discordgo.Session, m *discordgo.MessageCreate,
	id int, cc string, l string) error {
	page, err := scraper.Grouping(id, cc, l)
	if err != nil {
		return err
	}

	children := page.PageData.FcStructure.Model.Children[0].Children

	topSection := make([]room, 0)
	regularSection := make([]room, 0)
	for _, child := range children {
		if child.FcKind == "311" ||
			child.FcKind == "312" ||
			child.FcKind == "424" ||
			child.FcKind == "425" ||
			child.FcKind == "437" ||
			child.FcKind == "476" {
			continue
		}

		if child.FcKind == "415" {
			for _, top := range child.Children {
				topSection = append(topSection, room{
					contentIDs:  []string{top.Link.ContentID},
					designBadge: top.DesignBadge,
					fcKind:      top.FcKind,
					imageURL:    util.ConvertArtworkURL(top.Artwork.URL),
				})
			}
		} else {
			contentIDs := make([]string, 0)
			for _, cID := range child.Content {
				contentIDs = append(contentIDs, cID.ContentID)
			}

			regularSection = append(regularSection, room{
				contentIDs:   contentIDs,
				displayStyle: child.DisplayStyle,
				fcKind:       child.FcKind,
				name:         child.Name,
				tagline:      child.Tagline,
			})
		}
	}

	for _, room := range topSection {
		contentIDs := strings.Builder{}
		for _, contentID := range room.contentIDs {
			contentIDs.WriteString(fmt.Sprintf("%s\n", contentID))
		}

		time.Sleep(500)
		util.Send(s, m, util.Message{
			Title: room.designBadge,
			Description: fmt.Sprintf("%s\n%s\n%s\n%s", room.tagline,
				room.designBadge, room.designTag, room.displayStyle),
			ImageURL:      util.ConvertArtworkURL(room.imageURL),
			FooterText:    contentIDs.String(),
			FooterIconURL: util.AsLogoURL,
		})
	}

	for _, room := range regularSection {
		contentIDs := strings.Builder{}
		for _, contentID := range room.contentIDs {
			contentIDs.WriteString(fmt.Sprintf("%s\n", contentID))
		}

		time.Sleep(500)
		util.Send(s, m, util.Message{
			Title: room.name,
			Description: fmt.Sprintf("%s\n%s\n%s\n%s", room.tagline,
				room.designBadge, room.designTag, room.displayStyle),
			ImageURL:      util.ConvertArtworkURL(room.imageURL),
			FooterText:    contentIDs.String(),
			FooterIconURL: util.AsLogoURL,
		})
	}

	return nil
}

func processStory(s *discordgo.Session, m *discordgo.MessageCreate,
	storyID int, cc string, l string) error {
	page, err := scraper.Story(storyID, cc, l)
	if err != nil {
		return err
	}

	results := page.StorePlatformData["editorial-item-product"].Results
	result := results[strconv.Itoa(storyID)]

	iu := ""
	for _, v := range result.EditorialArtwork {
		iu = util.ConvertArtworkURL(v.URL)
		break
	}

	ids := strings.Builder{}
	for _, id := range result.CardIDs {
		ids.WriteString(fmt.Sprintf("%s\n", id))
	}

	util.Send(s, m, util.Message{
		Title: result.Label,
		Description: fmt.Sprintf("%s\n**%s**\n%s", page.PageData.PageType,
			result.EditorialNotes.Name, result.EditorialNotes.Short),
		Link:          result.Link.URL,
		ImageURL:      iu,
		FooterText:    ids.String(),
		FooterIconURL: util.AsLogoURL,
	})

	return nil
}

func processRoom(s *discordgo.Session, m *discordgo.MessageCreate,
	fcID int, cc string, l string) error {
	page, err := scraper.Room(fcID, cc, l)
	if err != nil {
		return err
	}

	contentIDs := strings.Builder{}
	for _, contentID := range page.PageData.AdamIds {
		contentIDs.WriteString(fmt.Sprintf("%d\n", contentID))
	}

	util.Send(s, m, util.Message{
		Title:         page.PageData.PageTitle,
		Description:   fmt.Sprintf("%s\n", page.PageData.MetricsBase.PageType),
		FooterText:    contentIDs.String(),
		FooterIconURL: util.AsLogoURL,
	})

	return nil
}

func isPhrase(s string) bool {
	regexp, _ := regexp.Compile("^.{1,}\\?$")

	return regexp.MatchString(s)
}
