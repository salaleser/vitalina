package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/detectlanguage/detectlanguage-go"
	voicerssgo "github.com/salaleser/voicerss-api-go"
)

// Config is a configuration.
var Config = make(map[string]string)

var languageDetectionClient *detectlanguage.Client

// LanguageDetection is a vitalina's Detection structure.
type LanguageDetection struct {
	Language        string
	ConfidenceScore float32
	IsReliable      bool
}

// Message is a vitalina's Message structure.
type Message struct {
	Title         string
	Description   string
	Link          string
	ImageURL      string
	ThumbnailURL  string
	FooterText    string
	FooterIconURL string
}

// Regexp patterns
const (
	ASAppIDPattern   = "^\\d{9,10}$"
	GPAppIDPattern   = "^[a-zA-Z][a-zA-Z0-9_]*(\\.[a-zA-Z0-9_]+)+[0-9a-zA-Z_]$"
	TimestampPattern = "[MWN]" // FIXME
)

// Store
const (
	NA = iota
	GooglePlay
	AppStore
)

const separator = "=" // разделитель для парсинга файлов

// ReadConfig reads lines from config file into the Config map.
func ReadConfig() {
	file, err := os.Open("config")
	if err != nil {
		log.Printf("Не удалось найти конфигурационный файл (%s)", err)
		return
	}
	defer file.Close()

	var line []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = strings.Split(scanner.Text(), separator)
		if len(line) == 2 {
			key := line[0]
			value := line[1]
			Config[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func InitLangaugeDetection() {
	languageDetectionClient = detectlanguage.New(Config["language-detection-api-key"])
}

// IsAppID returns 0 if appID is not an application ID.
func IsAppID(s string) int {
	gpAppIDRegexp, _ := regexp.Compile(GPAppIDPattern)
	asAppIDRegexp, _ := regexp.Compile(ASAppIDPattern)

	if gpAppIDRegexp.MatchString(s) {
		return GooglePlay
	}

	if asAppIDRegexp.MatchString(s) {
		return AppStore
	}

	return NA
}

func IsTimestamp(s string) bool {
	timestampRegexp, _ := regexp.Compile(TimestampPattern)

	return timestampRegexp.MatchString(s)
}

func GetFlagByCountry(code string) string {
	switch code {
	case "ru":
	case voicerssgo.Russian:
		return "🇷🇺"
	case "us":
	case voicerssgo.EnglishUnitedStates:
		return "🇺🇸"
	case "au":
	case voicerssgo.EnglishAustralia:
		return "🇦🇺"
	case "fr":
		return "🇫🇷"
	case "de":
		return "🇩🇪"
	case "no":
		return "🇳🇴"
	case "pl":
		return "🇵🇱"
	case "gb":
	case voicerssgo.EnglishGreatBritain:
		return "🇬🇧"
	case "es":
		return "🇪🇸"
	case "pt":
		return "🇵🇹"
	case "ca":
		return "🇨🇦"
	case "br":
	case voicerssgo.PortugueseBrazil:
		return "🇧🇷"
	case voicerssgo.ChineseHongKong:
		return "🇭🇰"
	case voicerssgo.ChineseChina:
		return "🇨🇳"
	case voicerssgo.Japanese:
		return "🇯🇵"
	default:
		return "🏳"
	}

	return "🏳"
}

// DetectLanguage trying to detect language by given text and returns detections array.
func DetectLanguage(value string) []LanguageDetection {
	detections, err := languageDetectionClient.Detect(value)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error detecting language:", err)
		return []LanguageDetection{}
	}

	result := make([]LanguageDetection, 0)
	for _, d := range detections {
		detection := LanguageDetection{d.Language, d.Confidence, d.Reliable}
		result = append(result, detection)
	}

	return result
}

func GetEmojiByDigit(digit float32) string {
	value := int(digit)

	if value < 0 {
		value = 0
	} else if value > 10 {
		value = 10
	}

	switch value {
	case 0:
		return "0️⃣"
	case 1:
		return "1️⃣"
	case 2:
		return "2️⃣"
	case 3:
		return "3️⃣"
	case 4:
		return "4️⃣"
	case 5:
		return "5️⃣"
	case 6:
		return "6️⃣"
	case 7:
		return "7️⃣"
	case 8:
		return "8️⃣"
	case 9:
		return "9️⃣"
	case 10:
		return "🔟"
	default:
		return ""
	}
}

// Now returns formatted current time.
func Now() string {
	return time.Now().Format(time.RFC3339)
}
