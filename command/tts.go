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

// Tts pronounces the text.
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

	language := "en-us"
	detections := util.DetectLanguage(tts)
	if len(detections) > 0 {
		language = detections[0].Language
	}

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

	country := util.Countries[language]
	s.MessageReactionAdd(m.ChannelID, m.ID, country.Emoji) // TODO

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
