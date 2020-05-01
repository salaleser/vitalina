package util

type Language struct {
	Code    string
	Emoji   string
	English string
	Russian string
	Native  string
}

// GetLanguageByCode returns language structure (code from ISO 639-1)
func GetLanguageByCode(languageCode string) Language {
	switch languageCode {
	case "aa":
		return Language{"aa", "🇪🇹", "Afar", "Афарский язык", "Afaraf"}
	case "ab":
		return Language{"ab", "🌏", "Abkhazian", "Абха́зский язы́к", "аҧсшәа"}
	case "af":
		return Language{"af", "🇿🇦", "Afrikaans", "Африкаанс", "Afrikaans"}
	case "ak":
		return Language{"ak", "🇬🇭", "Akan", "Ака́н", "Akan"}
	case "am":
		return Language{"am", "🌏", "Amharic", "Амхарский", "አማርኛ"}
	case "ar":
		return Language{"ar", "🌍", "Arabic", "Ара́бский", "العربية"}
	case "as":
		return Language{"as", "🇮🇳", "Assamese", "Асса́мский", "অসমীয়া"}
	case "ay":
		return Language{"ay", "🇧🇴", "Aymara", "Аймара́", "aymar aru"}
	case "az":
		return Language{"az", "🇦🇿", "Azerbaijani", "Азербайджа́нский язы́к", "azərbaycan dili"}
	case "ba":
		return Language{"ba", "🇷🇺", "Bashkir", "Башкирский язык", "башҡорт теле"}
	case "be":
		return Language{"be", "🇧🇾", "Belarusian", "Белору́сский язы́к", "беларуская мова"}
	case "bg":
		return Language{"bg", "🇧🇬", "Bulgarian", "Болга́рский язы́к", "български език"}
	case "bh":
		return Language{"bh", "🇮🇳", "Bihari languages", "Биха́рские языки", "भोजपुरी"}
	case "bi":
		return Language{"bi", "🇻🇺", "Bislama", "Бисла́ма", "Bislama"}
	case "bn":
		return Language{"bn", "🇧🇩", "BN", "", ""}
	case "bo":
		return Language{"bo", "🇨🇳", "Tibetan", "", ""}
	case "br":
		return Language{"br", "🇫🇷", "Breton", "Брето́нский язы́к", "brezhoneg"}
	case "bs":
		return Language{"bs", "🇧🇦", "Bosnian", "Босни́йский язы́к", "bosanski jezik"}
	// case "bug":
	// 	return Language{"bug", "", "BUGINESE", "", ""}
	case "ca":
		return Language{"ca", "🇪🇸", "Catalan", "Катала́нский язы́к", "català"}
	// case "ceb":
	// 	return Language{"ceb", "", "CEBUANO", "", ""}
	// case "chr":
	// 	return Language{"chr", "🌎", "CHEROKEE", "", ""}
	case "co":
		return Language{"co", "🇫🇷", "Corsican", "Корсика́нский язы́к", "corsu"}
	// case "crs":
	// 	return Language{"crs", "", "SESELWA", "", ""}
	case "cs":
		return Language{"cs", "🇨🇿", "Czech", "Че́шский язы́к", "čeština"}
	case "cy":
		return Language{"cy", "🏴󠁧󠁢󠁷󠁬󠁳󠁿", "Welsh", "Валли́йский язы́к", "Cymraeg"}
	case "da":
		return Language{"da", "🇩🇰", "Danish", "Датский язы́к", "dansk"}
	case "de":
		return Language{"de", "🇩🇪", "German", "Неме́цкий язык", "Deutsch"}
	case "dv":
		return Language{"dv", "🇲🇻", "Maldivian", "Мальди́вский язы́к", "ދިވެހި"}
	case "dz":
		return Language{"dz", "🇮🇳", "Dzongkha", "Дзонг-кэ", "རྫོང་ཁ"}
	// case "egy":
	// 	return Language{"egy", "🇪🇬", "EGY", "", ""}
	case "el":
		return Language{"el", "🇬🇷", "Greek", "Гре́ческий язы́к", "ελληνικά"}
	case "en":
		return Language{"en", "🇺🇸", "English", "Англи́йский язы́к", "English"}
	case "eo":
		return Language{"eo", "🇺🇳", "Esperanto", "Эспера́нто", "Esperanto"}
	case "es":
		return Language{"es", "🇪🇸", "Spanish", "Испа́нский язык", "Español"}
	case "et":
		return Language{"et", "🇪🇪", "Estonian", "Эсто́нский язы́к", "eesti"}
	case "eu":
		return Language{"eu", "🇪🇸", "Basque", "Ба́скский язы́к", "euskara"}
	case "fa":
		return Language{"fa", "🇮🇷", "Persian", "Перси́дский язы́к", "فارسی"}
	case "fi":
		return Language{"fi", "🇫🇮", "Finnish", "Фи́нский язы́к", "suomi"}
	case "fj":
		return Language{"fj", "🇫🇯", "Fijian", "Фиджийский язык", "vosa Vakaviti"}
	case "fo":
		return Language{"fo", "🇩🇰", "Faroese", "Фаре́рский язык", "føroyskt"}
	case "fr":
		return Language{"fr", "🇫🇷", "French", "Францу́зский язы́к", "français"}
	case "fy":
		return Language{"fy", "🇳🇱", "Western Frisian", "Западнофризский язык", "Frysk"}
	case "ga":
		return Language{"ga", "🇮🇪", "Irish", "Ирла́ндский язы́к", "Gaeilge"}
	case "gd":
		return Language{"gd", "🏴󠁧󠁢󠁳󠁣󠁴󠁿", "Scottish Gaelic", "Шотла́ндский язы́к", "Gàidhlig"}
	case "gl":
		return Language{"gl", "🇪🇸", "Galician", "Галиси́йский язык", "Galego"}
	case "gn":
		return Language{"gn", "🌎", "Guarani", "Гуарани́", "Avañe'ẽ"}
	// case "got":
	// 	return Language{"got", "", "GOTHIC", "", ""}
	case "gu":
		return Language{"gu", "🇮🇳", "Gujarati", "Гуджара́ти", "ગુજરાતી"}
	case "gv":
		return Language{"gv", "🇮🇲", "Manx", "Мэ́нский язык", "Gaelg"}
	case "ha":
		return Language{"ha", "🌍", "Hausa", "Ха́уса", "هَوُسَ"}
	// case "haw":
	// 	return Language{"haw", "🌏", "HAW", "Гавайский", ""}
	case "hi":
		return Language{"hi", "🇮🇳", "Hindi", "Хинди", "हिन्दी"}
	case "hmn":
		return Language{"hmn", "", "HMONG", "", ""}
	case "hr":
		return Language{"hr", "🇭🇷", "HR", "", ""}
	case "ht":
		return Language{"ht", "🇭🇹", "HT", "", ""}
	case "hu":
		return Language{"hu", "🇭🇺", "HU", "", ""}
	case "hy":
		return Language{"hy", "🇦🇲", "HY", "", ""}
	case "ia":
		return Language{"ia", "", "INTERLINGUA", "", ""}
	case "id":
		return Language{"id", "🇮🇩", "ID", "", ""}
	case "ie":
		return Language{"ie", "", "INTERLINGUE", "", ""}
	case "ig":
		return Language{"ig", "", "IGBO", "", ""}
	case "ik":
		return Language{"ik", "", "INUPIAK", "", ""}
	case "is":
		return Language{"is", "🇮🇸", "IS", "", ""}
	case "it":
		return Language{"it", "🇮🇹", "Italian", "Итальянский", ""}
	case "iu":
		return Language{"iu", "", "INUKTITUT", "", ""}
	case "iw":
		return Language{"iw", "🇮🇱", "HEBREW", "", ""}
	case "ja":
		return Language{"ja", "🇯🇵", "JA", "Японский", ""}
	case "jw":
		return Language{"jw", "🌏", "JAVANESE", "", ""}
	case "ka":
		return Language{"ka", "🇬🇪", "KA", "", ""}
	case "kha":
		return Language{"kha", "", "KHASI", "", ""}
	case "kk":
		return Language{"kk", "🇰🇿", "KK", "Казахский", ""}
	case "kl":
		return Language{"kl", "🇬🇱", "KL", "", ""}
	case "km":
		return Language{"km", "", "KHMER", "", ""}
	case "kn":
		return Language{"kn", "", "KANNADA", "", ""}
	case "ko":
		return Language{"ko", "🇰🇷", "KOREAN", "", ""}
	case "ks":
		return Language{"ks", "🇮🇳", "KASHMIRI", "", ""}
	case "ku":
		return Language{"ku", "🇹🇷", "KURDISH", "", ""} // для упрощения Турции флаг
	case "ky":
		return Language{"ky", "🇰🇬", "KY", "Киргизский", ""}
	case "la":
		return Language{"la", "", "LATIN", "", ""}
	case "lb":
		return Language{"lb", "🇱🇺", "LB", "", ""}
	case "lg":
		return Language{"lg", "", "GANDA", "", ""}
	case "lif":
		return Language{"lif", "", "LIMBU", "", ""}
	case "ln":
		return Language{"ln", "", "LINGALA", "", ""}
	case "lo":
		return Language{"lo", "", "LAOTHIAN", "", ""}
	case "lt":
		return Language{"lt", "🇱🇹", "LT", "", ""}
	case "lv":
		return Language{"lv", "🇱🇻", "LV", "", ""}
	case "mfe":
		return Language{"mfe", "", "MAURITIAN_CREOLE", "", ""}
	case "mg":
		return Language{"mg", "🇲🇬", "MG", "", ""}
	case "mi":
		return Language{"mi", "🌏", "MAORI", "", ""}
	case "mk":
		return Language{"mk", "🇲🇰", "MK", "", ""}
	case "ml":
		return Language{"ml", "", "MALAYALAM", "", ""}
	case "mn":
		return Language{"mn", "🇲🇳", "MN", "", ""}
	case "mr":
		return Language{"mr", "", "MARATHI", "", ""}
	case "ms":
		return Language{"ms", "🇲🇾", "MS", "", ""}
	case "mt":
		return Language{"mt", "🇲🇹", "MT", "", ""}
	case "my":
		return Language{"my", "🇲🇲", "MY", "", ""}
	case "na":
		return Language{"na", "🇳🇷", "NA", "", ""}
	case "ne":
		return Language{"ne", "🇳🇵", "NE", "", ""}
	case "nl":
		return Language{"nl", "🇳🇱", "NL", "", ""}
	case "no":
		return Language{"no", "🇳🇴", "NO", "Норвежский", ""}
	case "nr":
		return Language{"nr", "🇿🇦", "Ndebele", "", ""}
	case "nso":
		return Language{"nso", "🇿🇦", "PEDI", "", ""}
	case "ny":
		return Language{"ny", "🇲🇼", "NYANJA", "", ""}
	case "oc":
		return Language{"oc", "🇫🇷", "OC", "Окситанский", ""}
	case "om":
		return Language{"om", "🇸🇴", "OROMO", "", ""}
	case "or":
		return Language{"or", "🇮🇳", "ORIYA", "", ""}
	case "pa":
		return Language{"pa", "🇮🇳", "PUNJABI", "Панджа́би", ""}
	case "pl":
		return Language{"pl", "🇵🇱", "Polish", "Польский", ""}
	case "ps":
		return Language{"ps", "🇦🇫", "PASHTO", "Пушту", ""}
	case "pt":
		return Language{"pt", "🇵🇹", "PT", "Португальский", ""}
	case "qu":
		return Language{"qu", "🌎", "Quechuan", "Кечуа", ""}
	case "rm":
		return Language{"rm", "", "RHAETO_ROMANCE", "", ""}
	case "rn":
		return Language{"rn", "🇧🇮", "Kirundi", "", ""}
	case "ro":
		return Language{"ro", "🇷🇴", "RO", "", ""}
	case "ru":
		return Language{"ru", "🇷🇺", "Russian", "Русский", ""}
	case "rw":
		return Language{"rw", "🇷🇼", "Kinyarwanda", "", ""}
	case "sa":
		return Language{"sa", "🇮🇳", "SANSKRIT", "", ""}
	case "sco":
		return Language{"sco", "🏴󠁧󠁢󠁳󠁣󠁴󠁿", "SCO", "", ""}
	case "sd":
		return Language{"sd", "", "SANGO", "", ""}
	case "sg":
		return Language{"sg", "", "SANGO", "", ""}
	case "si":
		return Language{"si", "", "SINHALESE", "", ""}
	case "sk":
		return Language{"sk", "🇸🇰", "SK", "", ""}
	case "sl":
		return Language{"sl", "🇸🇮", "SL", "", ""}
	case "sm":
		return Language{"sm", "🇼🇸", "SM", "", ""}
	case "sn":
		return Language{"sn", "", "SHONA", "", ""}
	case "so":
		return Language{"so", "🇸🇴", "SO", "", ""}
	case "sq":
		return Language{"sq", "🇦🇱", "", "", ""}
	case "sr":
		return Language{"sr", "🇷🇸", "", "", ""}
	case "ss":
		return Language{"ss", "", "SISWANT", "", ""}
	case "st":
		return Language{"st", "", "SESOTHO", "", ""}
	case "su":
		return Language{"su", "", "SUNDANESE", "", ""}
	case "sv":
		return Language{"sv", "🇸🇪", "", "", ""}
	case "sw":
		return Language{"sw", "", "SWAHILI", "", ""}
	case "syr":
		return Language{"syr", "", "SYRIAC", "", ""}
	case "ta":
		return Language{"ta", "", "TAMIL", "", ""}
	case "te":
		return Language{"te", "", "TELUGU", "", ""}
	case "tg":
		return Language{"tg", "🇹🇯", "", "", ""}
	case "th":
		return Language{"th", "🇹🇭", "", "", ""}
	case "ti":
		return Language{"ti", "", "TIGRINYA", "", ""}
	case "tk":
		return Language{"tk", "🇹🇲", "", "", ""}
	case "tl":
		return Language{"tl", "", "TAGALOG", "", ""}
	case "tlh":
		return Language{"tlh", "🎬", "tlh", "Клингонский", ""}
	case "tn":
		return Language{"tn", "", "TSWANA", "", ""}
	case "to":
		return Language{"to", "🇹🇴", "", "", ""}
	case "tr":
		return Language{"tr", "🇹🇷", "", "", ""}
	case "ts":
		return Language{"ts", "", "TSONGA", "", ""}
	case "tt":
		return Language{"tt", "", "TATAR", "", ""}
	case "ug":
		return Language{"ug", "", "UIGHUR", "", ""}
	case "uk":
		return Language{"uk", "🇺🇦", "", "", ""}
	case "ur":
		return Language{"ur", "", "URDU", "", ""}
	case "uz":
		return Language{"uz", "🇺🇿", "", "", ""}
	case "ve":
		return Language{"ve", "", "VENDA", "", ""}
	case "vi":
		return Language{"vi", "🇻🇳", "", "", ""}
	case "vo":
		return Language{"vo", "🇺🇳", "VOLAPUK", "Волапюк", ""}
	case "war":
		return Language{"war", "🇵🇭", "Waray", "", ""}
	case "wo":
		return Language{"wo", "", "WOLOF", "", ""}
	case "xh":
		return Language{"xh", "🇿🇦", "XHOSA", "", ""}
	case "yi":
		return Language{"yi", "🇮🇱", "YIDDISH", "Идиш", ""}
	case "yo":
		return Language{"yo", "🇳🇬", "YORUBA", "", ""}
	case "za":
		return Language{"za", "🇨🇳", "ZHUANG", "", ""}
	case "zh":
		return Language{"zh", "🇨🇳", "ZH", "", ""}
	case "zh-Hant":
		return Language{"zh-Hant", "🇨🇳", "ZH-HANT", "", ""}
	case "zu":
		return Language{"zu", "🇿🇦", "ZULU", "", ""}
	default:
		return Language{"??", "🏳", "Unknown", "Неизвестный", ""}
	}
}
