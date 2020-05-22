package util

// Country is a coutnry structure.
type Country struct {
	Code    string
	Emoji   string
	English string
	Russian string
	X       string
}

// Countries contains supported countries.
var Countries = map[string]Country{
	"am": {"am", "🇦🇲", "Armenia", "Арме́ния", ""},
	"ar": {"ar", "🇦🇷", "Argentina", "Аргенти́на", ""},
	"at": {"at", "🇦🇹", "Austria", "А́встрия", ""},
	"au": {"au", "🇦🇺", "Australia", "Австра́лия", ""},
	"az": {"az", "🇦🇿", "Azerbaijan", "Азербайджа́н", ""},
	"bg": {"bg", "🇧🇬", "Bulgaria", "Болга́рия", ""},
	"br": {"br", "🇧🇷", "Brazil", "Брази́лия", ""},
	"by": {"by", "🇧🇾", "Belarus", "Респу́блика Белару́сь", ""},
	"ca": {"ca", "🇨🇦", "Canada", "Кана́да", ""},
	"cn": {"cn", "🇨🇳", "China", "Кита́й", ""},
	"co": {"co", "🇨🇴", "Colombia", "Колу́мбия", ""},
	"cz": {"cz", "🇨🇿", "Czech Republic", "Че́шская Респу́блика", ""},
	"de": {"de", "🇩🇪", "Germany", "Герма́ния", ""},
	"dk": {"dk", "🇩🇰", "Denmark", "Да́ния", ""},
	"es": {"es", "🇪🇸", "Spain", "Испа́ния", ""},
	"fi": {"fi", "🇫🇮", "Finland", "Финля́ндия", ""},
	"fr": {"fr", "🇫🇷", "France", "Фра́нция", ""},
	"gb": {"gb", "🇬🇧", "United Kingdom", "Великобрита́ния", ""},
	"ge": {"ge", "🇬🇪", "Georgia", "Гру́зия", ""},
	"gr": {"gr", "🇬🇷", "Greece", "Гре́ция", ""},
	"hr": {"hr", "🇭🇷", "Croatia", "Хорва́тия", ""},
	"hu": {"hu", "🇭🇺", "Hungary", "Ве́нгрия", ""},
	"id": {"id", "🇮🇩", "Indonesia", "Индоне́зия", ""},
	"ie": {"ie", "🇮🇪", "Ireland", "Ирла́ндия", ""},
	"in": {"in", "🇮🇳", "India", "И́ндия", ""},
	"is": {"is", "🇮🇸", "Iceland", "Исла́ндия", ""},
	"it": {"it", "🇮🇹", "Italy", "Ита́лия", ""},
	"jp": {"jp", "🇯🇵", "Japan", "Япо́ния", ""},
	"kz": {"kz", "🇰🇿", "Kazakhstan", "Казахста́н", ""},
	"nl": {"nl", "🇳🇱", "Netherlands", "Нидерла́нды", ""},
	"no": {"no", "🇳🇴", "Norway", "Норве́гия", ""},
	"nz": {"nz", "🇳🇿", "New Zealand", "Но́вая Зела́ндия", ""},
	"pl": {"pl", "🇵🇱", "Poland", "По́льша", ""},
	"pt": {"pt", "🇵🇹", "Portugal", "Португа́лия", ""},
	"ro": {"ro", "🇷🇴", "Romania", "Румы́ния", ""},
	"ru": {"ru", "🇷🇺", "Russia", "Росси́я", ""},
	"se": {"se", "🇸🇪", "Sweden", "Шве́ция", ""},
	"ch": {"ch", "🇨🇭", "Switzerland", "Швейца́рия", ""},
	"ua": {"ua", "🇺🇦", "Ukraine", "Украи́на", ""},
	"us": {"us", "🇺🇸", "USA", "США", ""},
	"vn": {"vn", "🇻🇳", "Vietnam", "Вьетна́м", ""},
	"za": {"za", "🇿🇦", "South Africa", "Ю́жно-Африка́нская Респу́блика", ""},
}
