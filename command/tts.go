package command

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/salaleser/vitalina/db"
	"github.com/salaleser/vitalina/util"
	voicerssgo "github.com/salaleser/voicerss-api-go"
)

const path = "cache/"

func Tts(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")

	dir, err := ioutil.ReadDir("cache/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading cache directory: %v\n", err)
		return
	}

	if len(args) == 1 {
		var dirSize int64
		for _, file := range dir {
			dirSize += file.Size()
		}

		c := db.GetTTSCount()
		text := fmt.Sprintf("База данных содержит %d ссылок.\nВсего в кэше %d файлов (%d Mb)", c, len(dir), dirSize/1024/1024)
		s.ChannelMessageSend(m.ChannelID, text)

		return
	}

	tts := strings.Join(args[1:], " ")

	if !isProper(tts) {
		a := []string{
			"Я отказываюсь это произносить",
			"Я в этой хуйне не участвую",
			"Сам попробуй такое произнести",
			"Я на такое не подписывалась",
		}

		s.ChannelMessageSend(m.ChannelID, a[rand.Intn(len(a))])
		return
	}

	language := getLanguage(tts, s.State.Settings.Locale)

	filename, counter := db.GetTTS(tts, language)

	var cached bool
	var file *os.File
	if counter > 0 {
		counter++
		file, err = os.Open(path + filename + ".wav")
		if err != nil {
			fmt.Fprintf(os.Stderr, "sound file opening: %v\n", err)
		}
		cached = true
	} else {
		var zeroes string
		switch len(strconv.Itoa(counter)) {
		case 1:
			zeroes = "0000"
			break
		case 2:
			zeroes = "000"
			break
		case 3:
			zeroes = "00"
			break
		case 4:
			zeroes = "0"
			break
		default:
			zeroes = ""
		}

		filename = path + zeroes + strconv.Itoa(counter+1)
		file, err = voicerssgo.Get(util.Config["voice-rss-api-key"], language, tts, filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "getting sound file: %v\n", err)
			return
		}
	}

	db.UpdateTTS(tts, filename, counter, language, util.Now())

	util.PlayFile(s, m, file.Name())

	emoji := util.GetFlagByCountry(language)
	s.MessageReactionAdd(m.ChannelID, m.ID, emoji) // TODO

	time.Sleep(100 * time.Millisecond)
	if cached {
		s.MessageReactionAdd(m.ChannelID, m.ID, "💾")
	}
}

func isProper(text string) bool {
	p1, err := regexp.Compile("\\d{8,}")
	if err != nil {
		fmt.Fprintf(os.Stderr, "p1 regexp compilation: %v\n", err)
	}

	p2, err := regexp.Compile("[а-яёА-ЯЁa-zA-Z0-9]{16,}")
	if err != nil {
		fmt.Fprintf(os.Stderr, "p2 regexp compilation: %v\n", err)
	}

	p3, err := regexp.Compile("[^а-яёА-ЯЁa-zA-Z0-9]{3,}")
	if err != nil {
		fmt.Fprintf(os.Stderr, "p3 regexp compilation: %v\n", err)
	}

	return !p1.MatchString(text) && !p2.MatchString(text) && !p3.MatchString(text)
}

func getLanguage(tts string, locale string) string {
	language := voicerssgo.EnglishUnitedStates
	switch locale {
	case "Brazil":
		language = voicerssgo.PortugueseBrazil
		break
	case "Western Europe":
	case "Central Europe":
		language = voicerssgo.EnglishGreatBritain
		break
	case "Hong Kong":
		language = voicerssgo.ChineseHongKong
		break
	case "Japan":
		language = voicerssgo.Japanese
		break
	case "Russia":
		language = voicerssgo.Russian
		break
	case "Sydney":
		language = voicerssgo.EnglishAustralia
		break
	case "Singapore":
	case "US Central":
	case "US East":
	case "US South":
	case "US West":
		language = voicerssgo.EnglishUnitedStates
		break
	}

	chPattern, err := regexp.Compile("^[\u4E00-\u9FA5]+$")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch regexp compilation: %v\n", err)
	}

	ruPattern, err := regexp.Compile("^[А-ЯЁа-яё]+$")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ru regexp compilation: %v\n", err)
	}

	enPattern, err := regexp.Compile("^[A-Za-z]+$")
	if err != nil {
		fmt.Fprintf(os.Stderr, "en regexp compilation: %v\n", err)
	}

	if chPattern.MatchString(tts) {
		language = voicerssgo.ChineseChina
	}

	if ruPattern.MatchString(tts) {
		language = voicerssgo.Russian
	}

	if enPattern.MatchString(tts) {
		language = voicerssgo.EnglishUnitedStates
	}

	return language
}
