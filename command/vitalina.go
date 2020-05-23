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
	contentIDs         []string
	designBadge        string
	designTag          string
	displayStyle       string
	doNotFilter        bool
	fcKind             string
	imageURL           string
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
		if _, ok := scraper.StoreFronts[arg]; ok {
			util.Debug(fmt.Sprintf("Country code %q detected!", arg))
			cc = arg
			country := util.Countries[strings.ToLower(arg)]
			s.MessageReactionAdd(m.ChannelID, m.ID, country.Emoji)
		}

		if _, ok := scraper.Languages[arg]; ok {
			util.Debug(fmt.Sprintf("Language %q detected!", arg))
			l = arg
			language := util.Languages[strings.Split(arg, "-")[0]]
			s.MessageReactionAdd(m.ChannelID, m.ID, language.Emoji)
		}
	}

	// Scan for IDs
	for _, arg := range args {
		// Store Front
		id, _ := strconv.Atoi(arg)
		if util.ContainsMap(scraper.StoreFronts, id) {
			cc := util.GetCcByStoreFront(id)
			country := util.Countries[strings.ToLower(cc)]
			g, _ := s.Guild(m.GuildID)
			rl := util.ConvertDiscordRegionToLanguage(g.Region)
			translate := country.Translate(rl)
			util.Send(s, m, util.Message{
				Title: fmt.Sprintf("App Store Store Front detected by ID «%d»",
					id),
				Description: fmt.Sprintf("%s\n%s (%s)", country.Emoji,
					country.Title, translate),
				FooterText: fmt.Sprintf("Country Code: %s\nStore Front ID: %d",
					cc, id),
				FooterIconURL: util.AsLogoURL,
			})
		}

		if force {
			// Coutry Code
			if sf, ok := scraper.StoreFronts[strings.ToUpper(arg)]; ok {
				country := util.Countries[strings.ToLower(arg)]
				g, _ := s.Guild(m.GuildID)
				rl := util.ConvertDiscordRegionToLanguage(g.Region)
				translate := country.Translate(rl)
				util.Send(s, m, util.Message{
					Title: fmt.Sprintf("App Store Country Code detected by code «%s»",
						arg),
					Description: fmt.Sprintf("%s %s", country.Emoji,
						translate),
					FooterText: fmt.Sprintf("ISO 3166-1 alpha-2 code: %s\n"+
						"Store Front ID: %d\nTitle: %s", strings.ToUpper(arg),
						sf, country.Title),
					FooterIconURL: util.AsLogoURL,
				})
			}

			// App Store Langauge Code
			if asl, ok := scraper.Languages[arg]; ok {
				language := util.Languages[strings.Split(arg, "-")[0]]
				g, _ := s.Guild(m.GuildID)
				rl := util.ConvertDiscordRegionToLanguage(g.Region)
				translate := language.Translate(rl)
				util.Send(s, m, util.Message{
					Title: fmt.Sprintf("App Store Langauge detected by code «%s»",
						arg),
					Description: fmt.Sprintf("%s %s (%s)",
						language.Emoji, translate, language.Native),
					FooterText: fmt.Sprintf("Language Code: %s\n"+
						"Language ID: %d\nTitle: %s", arg, asl, language.Title),
					FooterIconURL: util.AsLogoURL,
				})
			}

			// ISO 639-1 code
			if isol, ok := util.Languages[arg]; ok {
				g, _ := s.Guild(m.GuildID)
				rl := util.ConvertDiscordRegionToLanguage(g.Region)
				translate := isol.Translate(rl)
				util.Send(s, m, util.Message{
					Title: fmt.Sprintf("ISO 639-1 code detected by code «%s»",
						arg),
					Description: fmt.Sprintf("%s %s (%s)",
						isol.Emoji, translate, isol.Native),
					FooterText: fmt.Sprintf("ISO 639-1 code: %s\n"+
						"Title: %s", arg, isol.Title),
					FooterIconURL: util.AsLogoURL,
				})
			}
		}

		if util.MatchesAsAppID(arg) {
			id, _ := strconv.Atoi(arg)
			err := processApp(s, m, id, cc, l)
			if err != nil {
				util.Debug(fmt.Sprintf("app [id=%d,cc=%s,l=%s]: %v", id, cc, l,
					err))
				if force {
					util.SendError(s, m,
						fmt.Sprintf("App [id=%d,cc=%s,l=%s]", id, cc, l),
						err,
					)
				}
			}
		}

		if util.MatchesAsStoryID(arg) {
			id, _ := strconv.Atoi(arg)
			err := processStory(s, m, id, cc, l)
			if err != nil {
				util.Debug(fmt.Sprintf("story [id=%d,cc=%s,l=%s]: %v",
					id, cc, l, err))
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
				util.Debug(fmt.Sprintf("room [id=%d,cc=%s,l=%s]: %v",
					id, cc, l, err))
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
				util.Debug(fmt.Sprintf("grouping [id=%d,cc=%s,l=%s]: %v",
					id, cc, l, err))
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
			err := processGenre(s, m, id, cc, l)
			if err != nil {
				util.Debug(fmt.Sprintf("genre [id=%d,cc=%s,l=%s]: %v",
					id, cc, l, err))
				if force {
					util.SendError(s, m,
						fmt.Sprintf("Genre [id=%d,cc=%s,l=%s]", id, cc, l),
						err,
					)
				}
			}
		}

		if util.GetStoreFromAppID(arg) == util.GooglePlay {
			err := processGpApp(s, m, arg, cc, l)
			if err != nil {
				util.Debug(fmt.Sprintf("gp app"+
					" [package_name=%s,gl=%s,hl=%s]: %v", arg, cc, l,
					err))
				if force {
					util.SendError(s, m,
						fmt.Sprintf("Google Play application"+
							" [package_name=%s,gl=%s,hl=%s]", arg, cc, l),
						err,
					)
				}
			}
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

func processGpApp(s *discordgo.Session, m *discordgo.MessageCreate,
	packageName string, gl string, hl string) error {
	metadata := scraper.GpMetadata(packageName, gl, hl)

	util.Send(s, m, util.Message{
		Title:       metadata.Title,
		Description: metadata.Subtitle,
		Link: fmt.Sprintf("https://play.google.com/store/apps/details?id=%s",
			packageName),
		ImageURL:     metadata.Screenshot1,
		ThumbnailURL: metadata.Logo,
		FooterText: fmt.Sprintf(""+
			"ID: %s\n"+
			"Author: %s\n"+
			"Release Date: %s\n"+
			"Rating: %s",
			metadata.AppID,
			metadata.ArtistName,
			metadata.ReleaseDate,
			util.GetStarsBar(int(metadata.Rating)),
		),
		FooterIconURL: util.GpLogoURL,
	})

	return nil
}

func processApp(s *discordgo.Session, m *discordgo.MessageCreate,
	appID int, cc string, l string) error {
	page, err := scraper.App(appID, cc, l)
	if err != nil {
		return err
	}

	results := page.StorePlatformData["product-dv"].Results
	result := results[strconv.Itoa(appID)]

	iu := ""
	for _, screenshots := range result.ScreenshotsByType {
		if len(screenshots) == 0 {
			continue
		}

		iu = util.ConvertArtworkURL(screenshots[0].URL)
	}

	util.Send(s, m, util.Message{
		Title:        result.Name,
		Description:  result.Subtitle,
		Link:         result.URL,
		ImageURL:     iu,
		ThumbnailURL: result.Artwork.URL,
		FooterText: fmt.Sprintf(""+
			"ID: %s\n"+
			"Author: %s\n"+
			"Rating: %s\n"+
			"Rating Count: %d\n"+
			"Bundle ID: %s\n"+
			"Artist ID: %s\n"+
			"Related Editorial Items: [%s]",
			result.ID,
			result.ArtistName,
			util.GetStarsBar(int(result.UserRating.Value)),
			result.UserRating.RatingCount,
			result.BundleID,
			result.ArtistID,
			util.MakeListString(result.RelatedEditorialItems),
		),
		FooterIconURL: util.AsLogoURL,
	})

	return nil
}

func processGenre(s *discordgo.Session, m *discordgo.MessageCreate,
	id int, cc string, l string) error {
	page, err := scraper.Genre(id, cc)
	if err != nil {
		return err
	}

	genreTitle := strings.Split(page.PageData.MetricsBase.PageDetails, "_")[0]
	groupingID, _ := strconv.Atoi(page.PageData.MetricsBase.PageID)

	util.Send(s, m, util.Message{
		Title:       fmt.Sprintf("Genre detected by ID «%d»", id),
		Description: fmt.Sprintf("**%s**", genreTitle),
		Link: strings.Replace(page.PageData.AllCategoriesLink.URL,
			"36", strconv.Itoa(id), 1),
		FooterText: fmt.Sprintf(""+
			"ID: %d\n"+
			"Title: %s\n"+
			"Grouping ID: %d\n"+
			"Store Front: %s\n"+
			"Language ID: %s",
			id,
			genreTitle,
			groupingID,
			page.PageData.MetricsBase.StoreFront,
			page.PageData.MetricsBase.Language,
		),
		FooterIconURL: util.AsLogoURL,
	})

	sections := extractRooms(&page)

	sendTopRooms(s, m, sections[0], id, groupingID,
		page.PageData.MetricsBase.StoreFront,
		page.PageData.MetricsBase.Language)
	sendRegularRooms(s, m, sections[1], id, groupingID,
		page.PageData.MetricsBase.StoreFront,
		page.PageData.MetricsBase.Language)

	return nil
}

func processGrouping(s *discordgo.Session, m *discordgo.MessageCreate,
	id int, cc string, l string) error {
	page, err := scraper.Grouping(id, cc, l)
	if err != nil {
		return err
	}

	sections := extractRooms(&page)

	sendTopRooms(s, m, sections[0], -1, id,
		page.PageData.MetricsBase.StoreFront,
		page.PageData.MetricsBase.Language)
	sendRegularRooms(s, m, sections[1], -1, id,
		page.PageData.MetricsBase.StoreFront,
		page.PageData.MetricsBase.Language)

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

	util.Send(s, m, util.Message{
		Title: fmt.Sprintf("Story detected by ID «%d»", storyID),
		Description: fmt.Sprintf(""+
			"**%s**\n"+
			"%s\n"+
			"%s\n\n"+
			"**App IDs:**\n"+
			"%s",
			result.Label,
			result.EditorialNotes.Name,
			result.EditorialNotes.Short,
			util.MakeListString(result.CardIDs),
		),
		Link:     result.Link.URL,
		ImageURL: iu,
		FooterText: fmt.Sprintf(""+
			"ID: %s\n"+
			"Display Style: %s\n"+
			"Card Display Style: %s\n"+
			"Display Sub Style: %s",
			result.ID,
			result.DisplayStyle,
			result.CardDisplayStyle,
			result.DisplaySubStyle,
		),
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

	util.Send(s, m, util.Message{
		Title: fmt.Sprintf("Room detected by ID «%d»", fcID),
		Description: fmt.Sprintf(""+
			"**%s**\n\n"+
			"**App IDs:**\n"+
			"%s",
			page.PageData.PageTitle,
			util.MakeListString(page.PageData.AdamIDs),
		),
		FooterText: fmt.Sprintf(""+
			"ID: %d\n"+
			"Store Front: %s\n"+
			"Language ID: %s",
			fcID,
			page.PageData.MetricsBase.StoreFront,
			page.PageData.MetricsBase.Language,
		),
		FooterIconURL: util.AsLogoURL,
	})

	return nil
}

func isPhrase(s string) bool {
	regexp, _ := regexp.Compile("^.{1,}\\?$")

	return regexp.MatchString(s)
}

func extractRooms(page *scraper.Page) [][]room {
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
					adamID:      child.AdamID,
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
				adamID:       child.AdamID,
				contentIDs:   contentIDs,
				displayStyle: child.DisplayStyle,
				fcKind:       child.FcKind,
				name:         child.Name,
				tagline:      child.Tagline,
			})
		}
	}

	return [][]room{
		topSection,
		regularSection,
	}
}

func sendTopRooms(s *discordgo.Session, m *discordgo.MessageCreate,
	rooms []room, genreID int, groupingID int, sf string, lc string) {
	description := strings.Builder{}
	for i, room := range rooms {
		description.WriteString(fmt.Sprintf("**%d** [%s]: **%q** (%s)\n",
			i+1, room.fcKind, room.designBadge, room.contentIDs[0]))
	}

	util.Send(s, m, util.Message{
		Title:       "Rooms (top section)",
		Description: description.String(),
		FooterText: fmt.Sprintf(""+
			"Genre ID: %d\n"+
			"Grouping ID: %d\n"+
			"Store Front: %s\n"+
			"Language ID: %s",
			genreID,
			groupingID,
			sf,
			lc,
		),
		FooterIconURL: util.AsLogoURL,
	})
}

func sendRegularRooms(s *discordgo.Session, m *discordgo.MessageCreate,
	rooms []room, genreID int, groupingID int, sf string, lc string) {
	d := strings.Builder{}
	for i, room := range rooms {
		d.WriteString(fmt.Sprintf("**%d** [%s]: **%q** (%s)\n",
			i+1, room.fcKind, room.name, room.adamID))
	}

	util.Send(s, m, util.Message{
		Title:       "Rooms (regular section)",
		Description: d.String(),
		FooterText: fmt.Sprintf(""+
			"Genre ID: %d\n"+
			"Grouping ID: %d\n"+
			"Store Front: %s\n"+
			"Language ID: %s",
			genreID,
			groupingID,
			sf,
			lc,
		),
		FooterIconURL: util.AsLogoURL,
	})
}
