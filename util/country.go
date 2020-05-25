package util

// Country is a coutnry structure.
type Country struct {
	// Cc is App Store country code
	Cc      string
	Emoji   string
	Title   string
	TitleRu string
	// Code is a ISO 3166-1 alpha-2 code
	Code string
}

// Translate translates country title
func (c Country) Translate(language string) string {
	switch language {
	case "en-us", "en-gb", "en-au":
		return c.Title
	case "ru-ru":
		return c.TitleRu
	}

	return c.Title
}

// Countries contains supported countries.
// TODO добавить страны
var Countries = map[string]Country{
	"am": {"AM", "🇦🇲", "Armenia", "Арме́ния", ""},
	"ar": {"AR", "🇦🇷", "Argentina", "Аргенти́на", ""},
	"at": {"AT", "🇦🇹", "Austria", "А́встрия", ""},
	"au": {"AU", "🇦🇺", "Australia", "Австра́лия", ""},
	"az": {"AZ", "🇦🇿", "Azerbaijan", "Азербайджа́н", ""},
	"bg": {"BG", "🇧🇬", "Bulgaria", "Болга́рия", ""},
	"br": {"BR", "🇧🇷", "Brazil", "Брази́лия", ""},
	"by": {"BY", "🇧🇾", "Belarus", "Респу́блика Белару́сь", ""},
	"ca": {"CA", "🇨🇦", "Canada", "Кана́да", ""},
	"cn": {"CN", "🇨🇳", "China", "Кита́й", ""},
	"co": {"CO", "🇨🇴", "Colombia", "Колу́мбия", ""},
	"cz": {"CZ", "🇨🇿", "Czech Republic", "Че́шская Респу́блика", ""},
	"de": {"DE", "🇩🇪", "Germany", "Герма́ния", ""},
	"dk": {"DK", "🇩🇰", "Denmark", "Да́ния", ""},
	"es": {"ES", "🇪🇸", "Spain", "Испа́ния", ""},
	"fi": {"FI", "🇫🇮", "Finland", "Финля́ндия", ""},
	"fr": {"FR", "🇫🇷", "France", "Фра́нция", ""},
	"gb": {"GB", "🇬🇧", "United Kingdom", "Великобрита́ния", ""},
	"ge": {"GE", "🇬🇪", "Georgia", "Гру́зия", ""},
	"gr": {"GR", "🇬🇷", "Greece", "Гре́ция", ""},
	"hr": {"HR", "🇭🇷", "Croatia", "Хорва́тия", ""},
	"hu": {"HU", "🇭🇺", "Hungary", "Ве́нгрия", ""},
	"id": {"ID", "🇮🇩", "Indonesia", "Индоне́зия", ""},
	"ie": {"IE", "🇮🇪", "Ireland", "Ирла́ндия", ""},
	"in": {"IN", "🇮🇳", "India", "И́ндия", ""},
	"is": {"IS", "🇮🇸", "Iceland", "Исла́ндия", ""},
	"it": {"IT", "🇮🇹", "Italy", "Ита́лия", ""},
	"jp": {"JP", "🇯🇵", "Japan", "Япо́ния", ""},
	"kz": {"KZ", "🇰🇿", "Kazakhstan", "Казахста́н", ""},
	"nl": {"NL", "🇳🇱", "Netherlands", "Нидерла́нды", ""},
	"no": {"NO", "🇳🇴", "Norway", "Норве́гия", ""},
	"nz": {"NZ", "🇳🇿", "New Zealand", "Но́вая Зела́ндия", ""},
	"pl": {"PL", "🇵🇱", "Poland", "По́льша", ""},
	"pt": {"PT", "🇵🇹", "Portugal", "Португа́лия", ""},
	"ro": {"RO", "🇷🇴", "Romania", "Румы́ния", ""},
	"ru": {"RU", "🇷🇺", "Russia", "Росси́я", ""},
	"se": {"SE", "🇸🇪", "Sweden", "Шве́ция", ""},
	"ch": {"CH", "🇨🇭", "Switzerland", "Швейца́рия", ""},
	"ua": {"UA", "🇺🇦", "Ukraine", "Украи́на", ""},
	"us": {"US", "🇺🇸", "USA", "США", ""},
	"vn": {"VN", "🇻🇳", "Vietnam", "Вьетна́м", ""},
	"za": {"ZA", "🇿🇦", "South Africa", "Ю́жно-Африка́нская Респу́блика", ""},
}
