package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/detectlanguage/detectlanguage-go"
	"github.com/salaleser/scraper"
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

	err = scanner.Err()
	if err != nil {
		log.Println(err)
	}
}

// InitLangaugeDetection initializes language detection client.
func InitLangaugeDetection() {
	languageDetectionClient = detectlanguage.New(Config["language-detection-api-key"])
}

// MatchesAsAppID reports whether the string s matches to App Store
// Application ID.
func MatchesAsAppID(s string) bool {
	r, _ := regexp.Compile("^\\d{9,10}$")
	if !r.MatchString(s) {
		return false
	}

	id, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	// TODO уточнить
	if id < 100000000 {
		return false
	}

	// TODO уточнить
	if id > 9999999999 {
		return false
	}

	Debug("App ID detected!")
	return true
}

// MatchesAsGroupingID reports whether the string s matches to App Store
// Grouping ID.
func MatchesAsGroupingID(s string) bool {
	r, _ := regexp.Compile("^\\d{5,6}$")
	if !r.MatchString(s) {
		return false
	}

	id, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	// TODO уточнить
	if id < 25000 {
		return false
	}

	// TODO уточнить
	if id > 172000 {
		return false
	}

	Debug("Grouping ID detected!")
	return true
}

// MatchesAsRoomID reports whether the string s matches to App Store Room ID.
func MatchesAsRoomID(s string) bool {
	r, _ := regexp.Compile("^\\d{2,20}$")
	return r.MatchString(s)
}

// MatchesAsGenreID reports whether the string s matches to App Store Genre ID.
func MatchesAsGenreID(s string) bool {
	r, _ := regexp.Compile("^\\d{2,5}$")
	if !r.MatchString(s) {
		return false
	}

	id, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	// Исключение для 36
	if id == 36 {
		return true
	}

	// TODO уточнить
	if id < 6000 {
		return false
	}

	// TODO уточнить
	if id > 15000 {
		return false
	}

	if !contains(scraper.GenreIDs, id) {
		return false
	}

	Debug("Genre ID detected!")
	return true
}

// MatchesAsStoryID reports whether the string s matches to App Store Story ID.
func MatchesAsStoryID(s string) bool {
	r, _ := regexp.Compile("^\\d{9,10}$")
	return r.MatchString(s)
}

// GetStoreFromAppID returns 0 if appID is not an application ID.
func GetStoreFromAppID(s string) int {
	gpAppIDRegexp, _ := regexp.Compile("^[a-zA-Z][a-zA-Z0-9_]*(\\.[a-zA-Z0-9_]+)+[0-9a-zA-Z_]$")

	if gpAppIDRegexp.MatchString(s) {
		return GooglePlay
	}

	if MatchesAsAppID(s) {
		return AppStore
	}

	return NA
}

// GetFlagByCountryCode returns flag emoji.
// TODO add countries
func GetFlagByCountryCode(code string) string {
	switch code {
	case "ru":
		return "🇷🇺"
	case voicerssgo.Russian:
		return "🇷🇺"
	case "us":
		return "🇺🇸"
	case voicerssgo.EnglishUnitedStates:
		return "🇺🇸"
	case "au":
		return "🇦🇺"
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
		return "🇬🇧"
	case voicerssgo.EnglishGreatBritain:
		return "🇬🇧"
	case "es":
		return "🇪🇸"
	case "pt":
		return "🇵🇹"
	case "ca":
		return "🇨🇦"
	case "br":
		return "🇧🇷"
	case voicerssgo.PortugueseBrazil:
		return "🇧🇷"
	case "hk":
		return "🇭🇰"
	case voicerssgo.ChineseHongKong:
		return "🇭🇰"
	case "cn":
		return "🇨🇳"
	case voicerssgo.ChineseChina:
		return "🇨🇳"
	case "jp":
		return "🇯🇵"
	case voicerssgo.Japanese:
		return "🇯🇵"
	default:
		return ""
	}
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

// GetEmojiByDigit returns emoji with the digit by given number.
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

func contains(a []int, x int) bool {
	for _, e := range a {
		if e == x {
			return true
		}
	}

	return false
}

// ContainsMap reports whether the map m contains value x.
func ContainsMap(m map[string]int, x int) bool {
	for _, v := range m {
		if v == x {
			return true
		}
	}

	return false
}

// GetCcByStoreFront returns App Store country code by store front.
func GetCcByStoreFront(storeFront int) string {
	for cc, sf := range scraper.StoreFronts {
		if sf == storeFront {
			// FIXME scraper.StoreFronts содержит ключи в верхнем регистре
			return strings.ToLower(cc)
		}
	}

	return ""
}

// GetStarsBar returns bar with stars.
func GetStarsBar(x int) string {
	switch x {
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

// ConvertArtworkURL returns valid image URL by App Store artwork special URL.
func ConvertArtworkURL(url string) string {
	return strings.Replace(url, "{w}x{h}{c}.{f}", "512x512bb.png", -1)
}
