package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/detectlanguage/detectlanguage-go"
)

// Config is a configuration
var Config = make(map[string]string)

var languageDetectionClient *detectlanguage.Client

type LanguageDetection struct {
	Language        string
	ConfidenceScore int
	IsReliable      bool
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

// ReadConfig reads lines from config file into the Config map
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

// IsAppID returns 0 if appID is not an application ID
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

func IsLocation(s string) bool {
	switch s {
	case "ru":
	case "us":
	case "au":
	case "fr":
	case "de":
	case "no":
	case "pl":
	case "gb":
	case "es":
	case "pt":
	case "ca":
	case "br":
		return true
	}

	return false
}

func IsTimestamp(s string) bool {
	timestampRegexp, _ := regexp.Compile(TimestampPattern)

	return timestampRegexp.MatchString(s)
}

func GetFlagByCountry(countryCode string) string {
	switch countryCode {
	case "ru":
		return "🇷🇺"
	case "us":
		return "🇺🇸"
	case "au":
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
	case "es":
		return "🇪🇸"
	case "pt":
		return "🇵🇹"
	case "ca":
		return "🇨🇦"
	case "br":
		return "🇧🇷"
	default:
		return "🏳"
	}
}

func GetFlagByLanguage(languageCode string) string {
	switch languageCode {
	case "aa":
		return "🇪🇹" // Afar
	case "ab":
		return "🌏"
	case "af":
		return "🇿🇦" // Африкаанс
	case "ak":
		return "🇬🇭" // Akan
	case "am":
		return "🌏" // Амхарский язык
	case "ar":
		return "🌍" // Арабский
	case "as":
		return "ASSAMESE"
	case "ay":
		return "AYMARA"
	case "az":
		return "🇦🇿"
	case "ba":
		return "BASHKIR"
	case "be":
		return "🇧🇾"
	case "bg":
		return "🇧🇬"
	case "bh":
		return "BIHARI"
	case "bi":
		return "BISLAMA"
	case "bn":
		return "🇧🇩"
	case "bo":
		return "🇨🇳" // Tibetan
	case "br":
		return "BRETON"
	case "bs":
		return "🇧🇦"
	case "bug":
		return "BUGINESE"
	case "ca":
		return "🇪🇸" // Катала́нский язы́к
	case "ceb":
		return "CEBUANO"
	case "chr":
		return "🌎" // CHEROKEE
	case "co":
		return "CORSICAN" // CORSICAN
	case "crs":
		return "SESELWA"
	case "cs":
		return "🇨🇿"
	case "cy":
		return "🏴󠁧󠁢󠁷󠁬󠁳󠁿"
	case "da":
		return "🇩🇰"
	case "de":
		return "🇩🇪"
	case "dv":
		return "🇲🇻" // Мальди́вский язы́к
	case "dz":
		return "DZONGKHA"
	case "egy":
		return "🇪🇬"
	case "el":
		return "🇬🇷"
	case "en":
		return "🇺🇸"
	case "eo":
		return "🇺🇳" // ESPERANTO
	case "es":
		return "🇪🇸"
	case "et":
		return "🇪🇪"
	case "eu":
		return "BASQUE"
	case "fa":
		return "PERSIAN"
	case "fi":
		return "🇫🇮"
	case "fj":
		return "🇫🇯"
	case "fo":
		return "🇩🇰" // FAROESE
	case "fr":
		return "🇫🇷"
	case "fy":
		return "🇮🇷" // Персидский язык
	case "ga":
		return "🇮🇪"
	case "gd":
		return "🏴󠁧󠁢󠁳󠁣󠁴󠁿" // SCOTS_GAELIC
	case "gl":
		return "🇪🇸" // GALICIAN Галисийский язык
	case "gn":
		return "GUARANI"
	case "got":
		return "GOTHIC"
	case "gu":
		return "GUJARATI"
	case "gv":
		return "🇮🇲"
	case "ha":
		return "HAUSA"
	case "haw":
		return "🌏" // Гавайский язык
	case "hi":
		return "🇮🇳" // Хинди
	case "hmn":
		return "HMONG"
	case "hr":
		return "🇭🇷"
	case "ht":
		return "🇭🇹"
	case "hu":
		return "🇭🇺"
	case "hy":
		return "🇦🇲"
	case "ia":
		return "INTERLINGUA"
	case "id":
		return "🇮🇩"
	case "ie":
		return "INTERLINGUE"
	case "ig":
		return "IGBO"
	case "ik":
		return "INUPIAK"
	case "is":
		return "🇮🇸"
	case "it":
		return "🇮🇹"
	case "iu":
		return "INUKTITUT"
	case "iw":
		return "🇮🇱" // HEBREW
	case "ja":
		return "🇯🇵"
	case "jw":
		return "🌏" // JAVANESE
	case "ka":
		return "🇬🇪"
	case "kha":
		return "KHASI"
	case "kk":
		return "🇰🇿"
	case "kl":
		return "🇬🇱"
	case "km":
		return "KHMER"
	case "kn":
		return "KANNADA"
	case "ko":
		return "🇰🇷" // KOREAN
	case "ks":
		return "🇮🇳" // KASHMIRI
	case "ku":
		return "🇹🇷" // KURDISH для упрощения Турции флаг
	case "ky":
		return "🇰🇬"
	case "la":
		return "LATIN"
	case "lb":
		return "🇱🇺"
	case "lg":
		return "GANDA"
	case "lif":
		return "LIMBU"
	case "ln":
		return "LINGALA"
	case "lo":
		return "LAOTHIAN"
	case "lt":
		return "🇱🇹"
	case "lv":
		return "🇱🇻"
	case "mfe":
		return "MAURITIAN_CREOLE"
	case "mg":
		return "🇲🇬"
	case "mi":
		return "🌏" // MAORI
	case "mk":
		return "🇲🇰"
	case "ml":
		return "MALAYALAM"
	case "mn":
		return "🇲🇳"
	case "mr":
		return "MARATHI"
	case "ms":
		return "🇲🇾"
	case "mt":
		return "🇲🇹"
	case "my":
		return "🇲🇲"
	case "na":
		return "🇳🇷"
	case "ne":
		return "🇳🇵"
	case "nl":
		return "🇳🇱"
	case "no":
		return "🇳🇴"
	case "nr":
		return "🇿🇦" // NDEBELE Ndebele
	case "nso":
		return "🇿🇦" // PEDI
	case "ny":
		return "🇲🇼" // NYANJA
	case "oc":
		return "🇫🇷" // Окситанский язык
	case "om":
		return "🇸🇴" // OROMO
	case "or":
		return "🇮🇳" // ORIYA
	case "pa":
		return "🇮🇳" // PUNJABI Панджа́би
	case "pl":
		return "🇵🇱"
	case "ps":
		return "🇦🇫" // PASHTO Пушту
	case "pt":
		return "🇵🇹"
	case "qu":
		return "🌎" // QUECHUA Quechuan languages Кечуа
	case "rm":
		return "RHAETO_ROMANCE"
	case "rn":
		return "🇧🇮" // Kirundi
	case "ro":
		return "🇷🇴"
	case "ru":
		return "🇷🇺"
	case "rw":
		return "🇷🇼" // Kinyarwanda
	case "sa":
		return "🇮🇳" // SANSKRIT
	case "sco":
		return "🏴󠁧󠁢󠁳󠁣󠁴󠁿"
	case "sd":
		return "SANGO"
	case "sg":
		return "SANGO"
	case "si":
		return "SINHALESE"
	case "sk":
		return "🇸🇰"
	case "sl":
		return "🇸🇮"
	case "sm":
		return "🇼🇸"
	case "sn":
		return "SHONA"
	case "so":
		return "🇸🇴"
	case "sq":
		return "🇦🇱"
	case "sr":
		return "🇷🇸"
	case "ss":
		return "SISWANT"
	case "st":
		return "SESOTHO"
	case "su":
		return "SUNDANESE"
	case "sv":
		return "🇸🇪"
	case "sw":
		return "SWAHILI"
	case "syr":
		return "SYRIAC"
	case "ta":
		return "TAMIL"
	case "te":
		return "TELUGU"
	case "tg":
		return "🇹🇯"
	case "th":
		return "🇹🇭"
	case "ti":
		return "TIGRINYA"
	case "tk":
		return "🇹🇲"
	case "tl":
		return "TAGALOG"
	case "tlh":
		return "🎬" // Клингонский язык
	case "tn":
		return "TSWANA"
	case "to":
		return "🇹🇴"
	case "tr":
		return "🇹🇷"
	case "ts":
		return "TSONGA"
	case "tt":
		return "TATAR"
	case "ug":
		return "UIGHUR"
	case "uk":
		return "🇺🇦"
	case "ur":
		return "URDU"
	case "uz":
		return "🇺🇿"
	case "ve":
		return "VENDA"
	case "vi":
		return "🇻🇳"
	case "vo":
		return "🇺🇳" // VOLAPUK Волапюк
	case "war":
		return "🇵🇭" // Waray language
	case "wo":
		return "WOLOF"
	case "xh":
		return "🇿🇦" // XHOSA
	case "yi":
		return "🇮🇱" // YIDDISH Идиш
	case "yo":
		return "🇳🇬" // YORUBA
	case "za":
		return "🇨🇳" // ZHUANG
	case "zh":
		return "🇨🇳"
	case "zh-Hant":
		return "🇨🇳"
	case "zu":
		return "🇿🇦" // ZULU
	default:
		return "🏳"
	}
}

func DetectLanguage(value string) []LanguageDetection {
	detections, err := languageDetectionClient.Detect(value)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error detecting language:", err)
		return []LanguageDetection{}
	}

	result := make([]LanguageDetection, 0)
	for i, d := range detections {
		log.Printf("%d: %q [%f] (%t)", i+1, d.Language, d.Confidence, d.Reliable)
		detection := LanguageDetection{d.Language, int(d.Confidence), d.Reliable}
		result = append(result, detection)
	}

	return result
}

func GetEmojiByDigit(value int) string {
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
