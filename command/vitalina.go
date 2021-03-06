package command

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	as "github.com/salaleser/appstoreapi"
	pb "github.com/salaleser/vitalina/proto"
	"github.com/salaleser/vitalina/util"
	"google.golang.org/grpc"
)

type room struct {
	adamID string
	// contentIDs         []string
	designBadge        string
	designTag          string
	displayStyle       string
	doNotFilter        bool
	fcKind             string
	imageURL           string
	links              []string
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
	gl := "us"
	hl := ""

	content := m.Content

	if strings.HasPrefix(m.Content, "?") {
		force = true
		content = m.Content[1:]
	}

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

	// Scan for IDs
	for _, arg := range args {
		// Store Front
		id, _ := strconv.Atoi(arg)
		if util.ContainsMap(as.StoreFronts, id) {
			cc := util.GetCcByStoreFront(id)
			country := util.Countries[strings.ToLower(cc)]
			g, _ := s.Guild(m.GuildID)
			rl := util.ConvertDiscordRegionToLanguage(g.Region)
			translate := country.Translate(rl)
			util.Send(s, m, util.Message{
				Title: fmt.Sprintf(""+
					"__App Store Store Front ID__ detected by ID `%d`",
					id,
				),
				Description: fmt.Sprintf(""+
					"%s %s (%s)",
					country.Emoji,
					country.Title,
					translate,
				),
				FooterText: fmt.Sprintf(""+
					"Country Code: %s\n"+
					"Store Front ID: %d",
					cc,
					id,
				),
				FooterIconURL: util.AsLogoURL,
			})
		}

		if force {
			// Coutry Code
			if sf, ok := as.StoreFronts[strings.ToUpper(arg)]; ok {
				country := util.Countries[strings.ToLower(arg)]
				g, _ := s.Guild(m.GuildID)
				rl := util.ConvertDiscordRegionToLanguage(g.Region)
				translate := country.Translate(rl)
				util.Send(s, m, util.Message{
					Title: fmt.Sprintf(""+
						"__App Store Country Code__ detected by code %q",
						arg,
					),
					Description: fmt.Sprintf(""+
						"%s %s",
						country.Emoji,
						translate,
					),
					FooterText: fmt.Sprintf(""+
						"ISO 3166-1 alpha-2 code: %s\n"+
						"Store Front ID: %d\n"+
						"Title: %s",
						strings.ToUpper(arg),
						sf,
						country.Title,
					),
					FooterIconURL: util.AsLogoURL,
				})
			}

			// App Store Langauge Code
			if asl, ok := as.Languages[arg]; ok {
				language := util.Languages[strings.Split(arg, "-")[0]]
				g, _ := s.Guild(m.GuildID)
				rl := util.ConvertDiscordRegionToLanguage(g.Region)
				translate := language.Translate(rl)
				util.Send(s, m, util.Message{
					Title: fmt.Sprintf(""+
						"__App Store Langauge__ detected by code %q",
						arg,
					),
					Description: fmt.Sprintf("%s %s (%s)",
						language.Emoji, translate, language.Native),
					FooterText: fmt.Sprintf(""+
						"Language Code: %s\n"+
						"Language ID: %d\n"+
						"Title: %s",
						arg,
						asl,
						language.Title,
					),
					FooterIconURL: util.AsLogoURL,
				})
			}

			// ISO 639-1 code
			if isol, ok := util.Languages[arg]; ok {
				g, _ := s.Guild(m.GuildID)
				rl := util.ConvertDiscordRegionToLanguage(g.Region)
				translate := isol.Translate(rl)
				util.Send(s, m, util.Message{
					Title: fmt.Sprintf(""+
						"ISO 639-1 code detected by code %q",
						arg,
					),
					Description: fmt.Sprintf(""+
						"%s %s (%s)",
						isol.Emoji,
						translate,
						isol.Native,
					),
					FooterText: fmt.Sprintf(""+
						"ISO 639-1 code: %s\n"+
						"Title: %s",
						arg,
						isol.Title,
					),
					FooterIconURL: util.AsLogoURL,
				})
			}
		}

		if util.MatchesAsAppID(arg) {
			id, _ := strconv.Atoi(arg)
			err := processApp(s, m, id, cc, l)
			if err != nil {
				util.Debug(fmt.Sprintf(
					"app [id=%d,cc=%s,l=%s]: %v",
					id, cc, l, err))
				if force {
					util.SendError(s, m,
						fmt.Sprintf(
							"App [id=%d,cc=%s,l=%s]",
							id, cc, l),
						err,
					)
				}
			}
		}

		if util.MatchesBundleID(arg) {
			id, _ := strconv.Atoi(arg)
			err := processBundle(s, m, id, cc, l)
			if err != nil {
				util.Debug(fmt.Sprintf(
					"bundle [id=%d,cc=%s,l=%s]: %v",
					id, cc, l, err))
				if force {
					util.SendError(s, m,
						fmt.Sprintf(
							"Bundle [id=%d,cc=%s,l=%s]",
							id, cc, l),
						err,
					)
				}
			}
		}

		if util.MatchesAsStoryID(arg) {
			id, _ := strconv.Atoi(arg)
			err := processStory(s, m, id, cc, l)
			if err != nil {
				util.Debug(fmt.Sprintf(
					"story [id=%d,cc=%s,l=%s]: %v",
					id, cc, l, err))
				if force {
					util.SendError(s, m,
						fmt.Sprintf(
							"Story [id=%d,cc=%s,l=%s]",
							id, cc, l),
						err,
					)
				}
			}
		}

		if util.MatchesAsRoomID(arg) {
			id, _ := strconv.Atoi(arg)
			err := processRoom(s, m, uint32(id), cc, l)
			if err != nil {
				util.Debug(fmt.Sprintf(
					"room [id=%d,cc=%s,l=%s]: %v",
					id, cc, l, err))
				if force {
					util.SendError(s, m,
						fmt.Sprintf(
							"Room [id=%d,cc=%s,l=%s]",
							id, cc, l),
						err,
					)
				}
			}
		}

		if util.MatchesAsGroupingID(arg) {
			id, _ := strconv.Atoi(arg)
			err := processGrouping(s, m, id, cc, l)
			if err != nil {
				util.Debug(fmt.Sprintf(
					"grouping [id=%d,cc=%s,l=%s]: %v",
					id, cc, l, err))
				if force {
					util.SendError(s, m,
						fmt.Sprintf(
							"Grouping [id=%d,cc=%s,l=%s]",
							id, cc, l),
						err,
					)
				}
			}
		}

		if util.MatchesAsGenreID(arg) {
			id, _ := strconv.Atoi(arg)
			err := processGenre(s, m, id, cc, l)
			if err != nil {
				util.Debug(fmt.Sprintf(
					"genre [id=%d,cc=%s,l=%s]: %v",
					id, cc, l, err))
				if force {
					util.SendError(s, m,
						fmt.Sprintf(
							"Genre [id=%d,cc=%s,l=%s]",
							id, cc, l),
						err,
					)
				}
			}
		}

		if util.GetStoreFromAppID(arg) == util.GooglePlay {
			err := processAppGp(s, m, arg, gl, hl)
			if err != nil {
				util.Debug(fmt.Sprintf(
					"gp app [package_name=%s,gl=%s,hl=%s]: %v",
					arg, gl, hl, err))
				if force {
					util.SendError(s, m,
						fmt.Sprintf(
							"Application [package_name=%s,gl=%s,hl=%s]",
							arg, gl, hl),
						err,
					)
				}
			}
		}
	}

	// Scan for phrases
	if isQuestion(m.Content) {
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

func processAppGp(s *discordgo.Session, m *discordgo.MessageCreate,
	packageName string, gl string, hl string) error {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", util.Config["grpc-host"],
		util.Config["grpc-port"]), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewScraperClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.GooglePlayAppRequest{
		PackageName:   packageName,
		GeoLocation:   gl,
		HumanLanguage: hl,
	}
	r, err := client.GooglePlayApp(ctx, req)
	if err != nil {
		log.Fatalf("could not get gp app: %v", err)
	}

	if r.Title == "" {
		return nil
	}

	util.Send(s, m, util.Message{
		Title: fmt.Sprintf(""+
			"__Application__ detected by package name `%s`",
			packageName,
		),
		Description: fmt.Sprintf(""+
			"__**%s**__\n"+
			"%s",
			r.Title,
			r.Subtitle,
		),
		Link: fmt.Sprintf("https://play.google.com/store/apps/details?id=%s",
			packageName),
		ImageURL:     r.Screenshot1,
		ThumbnailURL: r.Logo,
		FooterText: fmt.Sprintf(""+
			"Package Name: %s\n"+
			"Author: %s\n"+
			"Release Date: %s\n"+
			"Rating: %s",
			r.PackageName,
			r.ArtistName,
			r.ReleaseDate,
			util.GetStarsBar(int(r.Rating)),
		),
		FooterIconURL: util.GpLogoURL,
	})

	return nil
}

func processApp(s *discordgo.Session, m *discordgo.MessageCreate,
	id int, cc string, l string) error {
	page, err := as.App(uint32(id), cc, l)
	if err != nil {
		return err
	}

	const spd string = "product-dv"
	result := page.StorePlatformData[spd].Results[strconv.Itoa(id)]

	iu := ""
	for _, screenshots := range result.ScreenshotsByType {
		if len(screenshots) > 0 {
			iu = util.ConvertArtworkURL(screenshots[0].URL, 512, 512)
			break
		}
	}

	util.Send(s, m, util.Message{
		Title: fmt.Sprintf(""+
			"__Application__ detected by ID `%d`",
			id,
		),
		Description: fmt.Sprintf(""+
			"__**%s**__\n"+
			"%s\n\n"+
			"**Related Editorial Items:**\n%s",
			result.Name,
			result.Subtitle,
			util.MakeList(result.RelatedEditorialItems),
		),
		Link:         result.URL,
		ImageURL:     iu,
		ThumbnailURL: util.ConvertArtworkURL(result.Artwork.URL, 512, 512),
		FooterText: fmt.Sprintf(""+
			"ID: %s\n"+
			"Author: %s\n"+
			"Rating: %s\n"+
			"Rating Count: %d\n"+
			"Bundle ID: %s\n"+
			"Artist ID: %s\n"+
			"Kind: %s\n"+
			"Store Front ID: %s\n"+
			"Language ID: %s",
			result.ID,
			result.ArtistName,
			util.GetStarsBar(int(result.UserRating.Value)),
			result.UserRating.RatingCount,
			result.BundleID,
			result.ArtistID,
			result.Kind,
			page.PageData.MetricsBase.StoreFront,
			page.PageData.MetricsBase.Language,
		),
		FooterIconURL: util.AsLogoURL,
	})

	return nil
}

func processBundle(s *discordgo.Session, m *discordgo.MessageCreate,
	id int, cc string, l string) error {
	page, err := as.Bundle(uint32(id), cc, l)
	if err != nil {
		return err
	}

	const spd string = "product-dv"
	result := page.StorePlatformData[spd].Results[strconv.Itoa(id)]

	iu := ""
	for _, screenshots := range result.ScreenshotsByType {
		if len(screenshots) > 0 {
			iu = util.ConvertArtworkURL(screenshots[0].URL, 512, 512)
			break
		}
	}

	util.Send(s, m, util.Message{
		Title: fmt.Sprintf(""+
			"__Bundle__ detected by ID `%d`",
			id,
		),
		Description: fmt.Sprintf(""+
			"__**%s**__\n"+
			"%s\n\n"+
			"**Children IDs:**\n%s",
			result.Name,
			result.Description.Standard,
			util.MakeList(result.ChildrenIDs),
		),
		Link:         result.URL,
		ImageURL:     iu,
		ThumbnailURL: util.ConvertArtworkURL(result.Artwork.URL, 512, 512),
		FooterText: fmt.Sprintf(""+
			"ID: %s\n"+
			"Author: %s\n"+
			"Rating: %s\n"+
			"Rating Count: %d\n"+
			"Artist ID: %s\n"+
			"Kind: %s\n"+
			"Store Front ID: %s\n"+
			"Language ID: %s",
			result.ID,
			result.ArtistName,
			util.GetStarsBar(int(result.UserRating.Value)),
			result.UserRating.RatingCount,
			result.ArtistID,
			result.Kind,
			page.PageData.MetricsBase.StoreFront,
			page.PageData.MetricsBase.Language,
		),
		FooterIconURL: util.AsLogoURL,
	})

	return nil
}

func processGenre(s *discordgo.Session, m *discordgo.MessageCreate,
	id int, cc string, l string) error {
	page, err := as.Genre(uint32(id), cc)
	if err != nil {
		return err
	}

	genreTitle := strings.Split(page.PageData.MetricsBase.PageDetails, "_")[0]
	groupingID, _ := strconv.Atoi(page.PageData.MetricsBase.PageID)
	link := strings.Replace(page.PageData.AllCategoriesLink.URL, "36",
		strconv.Itoa(id), 1)

	sections := buildRooms(&page)
	d := strings.Builder{}
	for i, s := range sections {
		d.WriteString(fmt.Sprintf(""+
			"**Section %d:**\n"+
			"%s\n",
			i+1,
			buildDescription(s),
		))
	}

	util.Send(s, m, util.Message{
		Title: fmt.Sprintf(""+
			"__Genre__ detected by ID `%d`",
			id,
		),
		Description: fmt.Sprintf(""+
			"__**%s**__\n\n"+
			"%s",
			genreTitle,
			d.String(),
		),
		Link: link,
		FooterText: fmt.Sprintf(""+
			"ID: %d\n"+
			"Title: %s\n"+
			"Grouping ID: %d\n"+
			"Store Front ID: %s\n"+
			"Language ID: %s",
			page.PageData.GenreID,
			genreTitle,
			groupingID,
			page.PageData.MetricsBase.StoreFront,
			page.PageData.MetricsBase.Language,
		),
		FooterIconURL: util.AsLogoURL,
	})

	return nil
}

func processGrouping(s *discordgo.Session, m *discordgo.MessageCreate,
	id int, cc string, l string) error {
	page, err := as.Grouping(uint32(id), cc, l)
	if err != nil {
		return err
	}

	sections := buildRooms(&page)
	d := strings.Builder{}
	for i, s := range sections {
		d.WriteString(fmt.Sprintf(""+
			"**Section %d:**\n"+
			"%s\n",
			i+1,
			buildDescription(s),
		))
	}

	util.Send(s, m, util.Message{
		Title: fmt.Sprintf(""+
			"__Grouping__ detected by ID `%d`",
			id,
		),
		Description: d.String(),
		FooterText: fmt.Sprintf(""+
			"ID: %s\n"+
			"Content ID: %s\n"+
			"Store Front ID: %s\n"+
			"Language ID: %s",
			page.PageData.MetricsBase.PageID,
			page.PageData.ContentID,
			page.PageData.MetricsBase.StoreFront,
			page.PageData.MetricsBase.Language,
		),
		FooterIconURL: util.AsLogoURL,
	})

	return nil
}

func processStory(s *discordgo.Session, m *discordgo.MessageCreate,
	id int, cc string, l string) error {
	page, err := as.Story(uint32(id), cc, l)
	if err != nil {
		return err
	}

	const spd string = "editorial-item-product"
	result := page.StorePlatformData[spd].Results[strconv.Itoa(id)]

	iu := ""
	for _, v := range result.EditorialArtwork {
		iu = util.ConvertArtworkURL(v.URL, 512, 512)
		break
	}

	util.Send(s, m, util.Message{
		Title: fmt.Sprintf(""+
			"__Story__ detected by ID `%d`",
			id,
		),
		Description: fmt.Sprintf(""+
			"__**%s**__\n"+
			"%s\n"+
			"%s\n\n"+
			"**App IDs:**\n%s",
			result.Label,
			result.EditorialNotes.Name,
			result.EditorialNotes.Short,
			util.MakeList(result.CardIDs),
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
	id uint32, cc string, l string) error {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", util.Config["grpc-host"],
		util.Config["grpc-port"]), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewScraperClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.AppStoreRoomRequest{Id: id, CountryCode: cc, Language: l}
	resp, err := client.AppStoreRoom(ctx, req)
	if err != nil {
		return fmt.Errorf("could not get room: %v", err)
	}

	util.Send(s, m, util.Message{
		Title: fmt.Sprintf(""+
			"__Room__ detected by ID `%d`",
			id,
		),
		Description: fmt.Sprintf(""+
			"__**%s**__\n\n"+
			"**App IDs:**\n%s",
			resp.Title,
			util.MakeList(resp.ContentIds),
		),
		FooterText: fmt.Sprintf(""+
			"ID: %d\n"+
			"FC Kind: %d\n"+
			"Store Front ID: %d\n"+
			"Language ID: %d",
			resp.Id,
			resp.FcKind,
			resp.StoreFront,
			resp.LanguageId,
		),
		FooterIconURL: util.AsLogoURL,
	})

	return nil
}

func isQuestion(s string) bool {
	regexp, _ := regexp.Compile("^.{1,}\\?$")

	return regexp.MatchString(s)
}

func buildRooms(page *as.Page) [][]room {
	children := page.PageData.FcStructure.Model.Children[0].Children

	sections := make([][]room, 0)
	regularSection := make([]room, 0)
	for _, c := range children {
		switch c.FcKind {
		case "415":
			s := make([]room, 0)
			for _, r := range c.Children {
				s = append(s, room{
					adamID:      r.AdamID,
					designBadge: r.DesignBadge,
					designTag:   r.DesignTag,
					fcKind:      r.FcKind,
					imageURL:    util.ConvertArtworkURL(r.Artwork.URL, 512, 512),
					title:       r.Title,
				})
			}
			sections = append(sections, s)
		case "418", "420", "429":
			regularSection = append(regularSection, room{
				adamID:       c.AdamID,
				displayStyle: c.DisplayStyle,
				fcKind:       c.FcKind,
				name:         c.Name,
				tagline:      c.Tagline,
			})
		case "424", "425":
			s := make([]room, 0)
			for _, r := range c.Children {
				s = append(s, room{
					adamID: r.AdamID,
					name:   r.Name,
					fcKind: r.FcKind,
				})
			}
			sections = append(sections, s)
			break
		case "437":
			links := make([]string, 0)
			for _, l := range c.Links {
				links = append(links, fmt.Sprintf("%s — %s", l.Label, l.URL))
			}

			sections = append(sections, []room{
				{
					adamID: c.AdamID,
					links:  links,
					name:   c.Name,
					fcKind: c.FcKind,
				},
			})
		case "311", "312", "476":
			sections = append(sections, []room{
				{
					adamID: c.AdamID,
					name:   c.Name,
					fcKind: c.FcKind,
				},
			})
		}
	}

	sections = append(sections, regularSection)

	return sections
}

func buildDescription(rooms []room) string {
	t := ""
	s := ""
	d := strings.Builder{}
	for i, r := range rooms {
		if r.fcKind == "416" || r.fcKind == "417" {
			t = r.designTag
			s = r.designBadge
		} else if r.fcKind == "426" {
			t = r.name
			s = util.MakeList(r.links)
		} else if r.fcKind == "437" {
			t = r.name
			s = util.MakeList(r.links)
		} else {
			t = r.name
			s = r.tagline
		}

		d.WriteString(fmt.Sprintf("%s **%d**: **%q** %s `%s`\n",
			util.ConvertFcKind(r.fcKind), i+1, t, s, r.adamID))
	}

	return d.String()
}
