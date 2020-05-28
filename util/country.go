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
	"ar": {"AR", "🇦🇷", "Argentina", "Аргентина", ""},
	"at": {"AT", "🇦🇹", "Austria", "Австрия", ""},
	"au": {"AU", "🇦🇺", "Australia", "Австралия", ""},
	"az": {"AZ", "🇦🇿", "Azerbaijan", "Азербайджан", ""},
	"bg": {"BG", "🇧🇬", "Bulgaria", "Болгария", ""},
	"br": {"BR", "🇧🇷", "Brazil", "Бразилия", ""},
	"by": {"BY", "🇧🇾", "Belarus", "Респу́блика Беларусь", ""},
	"ca": {"CA", "🇨🇦", "Canada", "Канада", ""},
	"ch": {"CH", "🇨🇭", "Switzerland", "Швейцария", ""},
	"cl": {"CL", "🇨🇱", "Chile", "Чили", ""},
	"cn": {"CN", "🇨🇳", "China", "Китай", ""},
	"co": {"CO", "🇨🇴", "Colombia", "Колумбия", ""},
	"cz": {"CZ", "🇨🇿", "Czech Republic", "Чешская Республика", ""},
	"de": {"DE", "🇩🇪", "Germany", "Германия", ""},
	"dk": {"DK", "🇩🇰", "Denmark", "Дания", ""},
	"ec": {"EC", "🇪🇨", "Ecuador", "Эквадор", ""},
	"es": {"ES", "🇪🇸", "Spain", "Испания", ""},
	"fi": {"FI", "🇫🇮", "Finland", "Финляндия", ""},
	"fr": {"FR", "🇫🇷", "France", "Франция", ""},
	"gb": {"GB", "🇬🇧", "United Kingdom", "Великобритания", ""},
	"ge": {"GE", "🇬🇪", "Georgia", "Грузия", ""},
	"gr": {"GR", "🇬🇷", "Greece", "Греция", ""},
	"hr": {"HR", "🇭🇷", "Croatia", "Хорватия", ""},
	"hu": {"HU", "🇭🇺", "Hungary", "Венгрия", ""},
	"id": {"ID", "🇮🇩", "Indonesia", "Индонезия", ""},
	"ie": {"IE", "🇮🇪", "Ireland", "Ирландия", ""},
	"in": {"IN", "🇮🇳", "India", "Индия", ""},
	"is": {"IS", "🇮🇸", "Iceland", "Исландия", ""},
	"it": {"IT", "🇮🇹", "Italy", "Италия", ""},
	"jp": {"JP", "🇯🇵", "Japan", "Япония", ""},
	"kz": {"KZ", "🇰🇿", "Kazakhstan", "Казахстан", ""},
	"nl": {"NL", "🇳🇱", "Netherlands", "Нидерланды", ""},
	"no": {"NO", "🇳🇴", "Norway", "Норвегия", ""},
	"nz": {"NZ", "🇳🇿", "New Zealand", "Новая Зеландия", ""},
	"pe": {"PE", "🇵🇪", "Peru", "Перу", ""},
	"ph": {"PH", "🇵🇭", "Philippines", "Филиппины", ""},
	"pl": {"PL", "🇵🇱", "Poland", "Польша", ""},
	"pt": {"PT", "🇵🇹", "Portugal", "Португалия", ""},
	"ro": {"RO", "🇷🇴", "Romania", "Румыния", ""},
	"ru": {"RU", "🇷🇺", "Russia", "Россия", ""},
	"se": {"SE", "🇸🇪", "Sweden", "Швеция", ""},
	"sk": {"SK", "🇸🇰", "Slovakia", "Словакия", ""},
	"ua": {"UA", "🇺🇦", "Ukraine", "Украина", ""},
	"us": {"US", "🇺🇸", "USA", "США", ""},
	"vn": {"VN", "🇻🇳", "Vietnam", "Вьетнам", ""},
	"za": {"ZA", "🇿🇦", "South Africa", "Южно-Африканская Республика", ""},
}
