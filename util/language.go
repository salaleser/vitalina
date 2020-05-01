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
	// case "hmn":
	// 	return Language{"hmn", "", "HMONG", "", ""}
	case "hr":
		return Language{"hr", "🇭🇷", "Croatian", "Хорва́тский язы́к", "hrvatski jezik"}
	case "ht":
		return Language{"ht", "🇭🇹", "Haitian", "Гаитянский креольский язык", "Kreyòl ayisyen"}
	case "hu":
		return Language{"hu", "🇭🇺", "Hungarian", "Венге́рский язы́к", "magyar"}
	case "hy":
		return Language{"hy", "🇦🇲", "Armenian", "Армя́нский язы́к", "Հայերեն"}
	case "ia":
		return Language{"ia", "🇺🇳", "Interlingua", "Интерли́нгва", "Interlingua"}
	case "id":
		return Language{"id", "🇮🇩", "Indonesian", "Индонези́йский язы́к", "Bahasa Indonesia"}
	case "ie":
		return Language{"ie", "🇺🇳", "Interlingue", "Интерли́нгве", "Interlingue"}
	case "ig":
		return Language{"ig", "🇳🇬", "Igbo", "И́гбо", "Asụsụ Igbo"}
	case "ik":
		return Language{"ik", "🇨🇦", "Inupiaq", "Инупиак", "Iñupiaq"}
	case "is":
		return Language{"is", "🇮🇸", "Icelandic", "Исла́ндский язы́к", "Íslenska"}
	case "it":
		return Language{"it", "🇮🇹", "Italian", "Италья́нский язы́к", "Italiano"}
	case "iu":
		return Language{"iu", "🇨🇦", "Inuktitut", "Инуктитут", "ᐃᓄᒃᑎᑐᑦ"}
	case "iw":
		return Language{"he", "🇮🇱", "Hebrew", "Иври́т", "עברית"}
	case "ja":
		return Language{"ja", "🇯🇵", "Japanese", "Япо́нский язы́к", "日本語"}
	case "jw":
		return Language{"jv", "🇮🇩", "Javanese", "Ява́нский язы́к", "ꦧꦱꦗꦮ"}
	case "ka":
		return Language{"ka", "🇬🇪", "Georgian", "Грузи́нский язы́к", "ქართული"}
	// case "kha":
	// 	return Language{"kha", "", "KHASI", "", ""}
	case "kk":
		return Language{"kk", "🇰🇿", "Kazakh", "Каза́хский язы́к", "қазақ тілі"}
	case "kl":
		return Language{"kl", "🇬🇱", "Kalaallisut", "Гренла́ндский язы́к", "kalaallisut"}
	case "km":
		return Language{"km", "🇰🇭", "Central Khmer", "Кхмерский язык", "ភាសាខ្មែរ"}
	case "kn":
		return Language{"kn", "🇮🇳", "Kannada", "Ка́ннада", "ಕನ್ನಡ"}
	case "ko":
		return Language{"ko", "🇰🇷", "Korean", "Коре́йский язы́к", "한국어"}
	case "ks":
		return Language{"ks", "🇮🇳", "Kashmiri", "Кашми́рский язы́к", "कश्मीरी"}
	case "ku":
		return Language{"ku", "🇹🇷", "Kurdish", "Ку́рдские языки́", "Kurdî"} // для упрощения Турции флаг
	case "ky":
		return Language{"ky", "🇰🇬", "Kyrgyz", "Киргизский", "Кыргызча"}
	case "la":
		return Language{"la", "🇻🇦", "Latin", "Лати́нский язык", "latine"}
	case "lb":
		return Language{"lb", "🇱🇺", "Luxembourgish", "Люксембу́ргский язы́к", "Lëtzebuergesch"}
	case "lg":
		return Language{"lg", "🇺🇬", "Ganda", "Луга́нда", "Luganda"}
	case "lif":
		return Language{"li", "🇳🇱", "Limburgan", "Ли́мбургский язык", "Limburgs"}
	case "ln":
		return Language{"ln", "🇨🇩", "Lingala", "Линга́ла", "Lingála"}
	case "lo":
		return Language{"lo", "🇱🇦", "Lao", "Лао́сский язык", "ພາສາລາວ"}
	case "lt":
		return Language{"lt", "🇱🇹", "Lithuanian", "Лито́вский язы́к", "lietuvių kalba"}
	case "lv":
		return Language{"lv", "🇱🇻", "Latvian", "Латы́шский язы́к", "latviešu valoda"}
	// case "mfe":
	// 	return Language{"mfe", "", "MAURITIAN_CREOLE", "", ""}
	case "mg":
		return Language{"mg", "🇲🇬", "Malagasy", "Малагаси́йский язы́к", "fiteny malagasy"}
	case "mi":
		return Language{"mi", "🇳🇿", "Maori", "Ма́ори", "te reo Māori"}
	case "mk":
		return Language{"mk", "🇲🇰", "Macedonian", "Македо́нский язы́к", "македонски јазик"}
	case "ml":
		return Language{"ml", "🇮🇳", "Malayalam", "Малая́лам", "മലയാളം"}
	case "mn":
		return Language{"mn", "🇲🇳", "Mongolian", "Монго́льский язы́к", "Монгол хэл"}
	case "mr":
		return Language{"mr", "🇮🇳", "Marathi", "Мара́тхи", "मराठी"}
	case "ms":
		return Language{"ms", "🇲🇾", "Malay", "Мала́йский", "Bahasa Melayu"}
	case "mt":
		return Language{"mt", "🇲🇹", "Maltese", "Мальти́йский язы́к", "Malti"}
	case "my":
		return Language{"my", "🇲🇲", "Burmese", "Бирма́нский язы́к", "ဗမာစာ"}
	case "na":
		return Language{"na", "🇳🇷", "Nauru", "Науруанский язык", "Dorerin Naoero"}
	case "ne":
		return Language{"ne", "🇳🇵", "Nepali", "Непа́льский язык", "नेपाली"}
	case "nl":
		return Language{"nl", "🇳🇱", "Dutch", "Нидерла́ндский язы́к", "Nederlands"}
	case "no":
		return Language{"no", "🇳🇴", "Norwegian", "Норве́жский язык", "Norsk"}
	case "nr":
		return Language{"nr", "🇿🇦", "South Ndebele", "Южный Ндебеле", "isiNdebele"}
	// case "nso":
	// 	return Language{"nso", "🇿🇦", "PEDI", "", ""}
	case "ny":
		return Language{"ny", "🇲🇼", "Chichewa", "Чичева", "chiCheŵa"}
	case "oc":
		return Language{"oc", "🇫🇷", "Occitan", "Оксита́нский язы́к", "occitan"}
	case "om":
		return Language{"om", "🇪🇹", "Oromo", "Оромо", "Afaan Oromoo"}
	case "or":
		return Language{"or", "🇮🇳", "Oriya", "Ори́я", "ଓଡ଼ିଆ"}
	case "pa":
		return Language{"pa", "🇮🇳", "Punjabi", "Панджа́би", "ਪੰਜਾਬੀ"}
	case "pl":
		return Language{"pl", "🇵🇱", "Polish", "По́льский язы́к", "język polski"}
	case "ps":
		return Language{"ps", "🇦🇫", "Pashto", "Пушту́", "پښتو"}
	case "pt":
		return Language{"pt", "🇵🇹", "Portuguese", "Португа́льский язы́к", "Português"}
	case "qu":
		return Language{"qu", "🇧🇴", "Quechua", "Ке́чуа", "Runa Simi"}
	case "rm":
		return Language{"rm", "🇨🇭", "Romansh", "Рома́ншский язы́к", "Rumantsch Grischun"}
	case "rn":
		return Language{"rn", "🇧🇮", "Rundi", "Рунди", "Ikirundi"}
	case "ro":
		return Language{"ro", "🇷🇴", "Romanian", "Румы́нский язы́к", "Română"}
	case "ru":
		return Language{"ru", "🇷🇺", "Russian", "Ру́сский язы́к", "русский"}
	case "rw":
		return Language{"rw", "🇷🇼", "Kinyarwanda", "Руанда", "Ikinyarwanda"}
	case "sa":
		return Language{"sa", "🇮🇳", "Sanskrit", "Санскри́т", "संस्कृतम्"}
	// case "sco":
	// 	return Language{"sco", "🏴󠁧󠁢󠁳󠁣󠁴󠁿", "Scots", "", ""}
	case "sd":
		return Language{"sd", "🇵🇰", "Sindhi", "Си́ндхи", "सिन्धी"}
	case "sg":
		return Language{"sg", "🇨🇫", "Sango", "Санго", "yângâ tî sängö"}
	case "si":
		return Language{"si", "🇱🇰", "Sinhala", "Синга́льский язы́к", "සිංහල"}
	case "sk":
		return Language{"sk", "🇸🇰", "Slovak", "Слова́цкий язы́к", "Slovenčina"}
	case "sl":
		return Language{"sl", "🇸🇮", "Slovenian", "Слове́нский язы́к", "Slovenski Jezik"}
	case "sm":
		return Language{"sm", "🇼🇸", "Samoan", "Самоа́нский язы́к", "gagana fa'a Samoa"}
	case "sn":
		return Language{"sn", "🇿🇼", "Shona", "Шо́на", "chiShona"}
	case "so":
		return Language{"so", "🇸🇴", "Somali", "Сомали́йский язык", "Soomaaliga"}
	case "sq":
		return Language{"sq", "🇦🇱", "Albanian", "Алба́нский язы́к", "Shqip"}
	case "sr":
		return Language{"sr", "🇷🇸", "Serbian", "Се́рбский язы́к", "српски језик"}
	case "ss":
		return Language{"ss", "🇸🇿", "Swati", "Сва́ти", "SiSwati"}
	case "st":
		return Language{"st", "🇱🇸", "Sotho", "Сесо́то", "Sesotho"}
	case "su":
		return Language{"su", "🇮🇩", "Sundanese", "Сунданский язык", "Basa Sunda"}
	case "sv":
		return Language{"sv", "🇸🇪", "Swedish", "Шве́дский язы́к", "Svenska"}
	case "sw":
		return Language{"sw", "🇹🇿", "Swahili", "Суахи́ли", "Kiswahili"}
	// case "syr":
	// 	return Language{"syr", "", "SYRIAC", "", ""}
	case "ta":
		return Language{"ta", "🇮🇳", "Tamil", "Тами́льский язы́к", "தமிழ்"}
	case "te":
		return Language{"te", "🇮🇳", "Telugu", "Язык те́лугу", "తెలుగు"}
	case "tg":
		return Language{"tg", "🇹🇯", "Tajik", "Таджи́кский язы́к", "тоҷикӣ"}
	case "th":
		return Language{"th", "🇹🇭", "Thai", "Та́йский язык", "ไทย"}
	case "ti":
		return Language{"ti", "🇪🇷", "Tigrinya", "Тигринья", "ትግርኛ"}
	case "tk":
		return Language{"tk", "🇹🇲", "Turkmen", "Туркме́нский язы́к", "Türkmen"}
	case "tl":
		return Language{"tl", "🇵🇭", "Tagalog", "Тага́льский язык", "Wikang Tagalog"}
	case "tlh":
		return Language{"tlh", "🎬", "tlh", "Клингонский", ""}
	case "tn":
		return Language{"tn", "🇿🇦", "Tswana", "", "Setswana"}
	case "to":
		return Language{"to", "🇹🇴", "Tonga", "Тонганский язык", "Faka Tonga"}
	case "tr":
		return Language{"tr", "🇹🇷", "Turkish", "Туре́цкий язы́к", "Türkçe"}
	case "ts":
		return Language{"ts", "🇿🇦", "Tsonga", "Тсо́нга", "Xitsonga"}
	case "tt":
		return Language{"tt", "🇷🇺", "Tatar", "Тата́рский язы́к", "татар теле"}
	case "ug":
		return Language{"ug", "🇨🇳", "Uighur", "Уйгу́рский язы́к", "ئۇيغۇرچە‎"}
	case "uk":
		return Language{"uk", "🇺🇦", "Ukrainian", "Украи́нский язы́к", "Українська"}
	case "ur":
		return Language{"ur", "🇵🇰", "Urdu", "Урду́", "اردو"}
	case "uz":
		return Language{"uz", "🇺🇿", "Uzbek", "Узбе́кский язы́к", "Oʻzbek"}
	case "ve":
		return Language{"ve", "🇿🇦", "Venda", "Венда", "Tshivenḓa"}
	case "vi":
		return Language{"vi", "🇻🇳", "Vietnamese", "Вьетна́мский язы́к", "Tiếng Việt"}
	case "vo":
		return Language{"vo", "🇺🇳", "Volapük", "Волапю́к", "Volapük"}
	// case "war":
	// 	return Language{"war", "🇵🇭", "Waray", "", ""}
	case "wo":
		return Language{"wo", "🇸🇳", "Wolof", "Воло́ф", "Wollof"}
	case "xh":
		return Language{"xh", "🇿🇦", "Xhosa", "Ко́са", "isiXhosa"}
	case "yi":
		return Language{"yi", "🇮🇱", "Yiddish", "И́диш", "ייִדיש"}
	case "yo":
		return Language{"yo", "🇳🇬", "Yoruba", "Йо́руба", "Yorùbá"}
	case "za":
		return Language{"za", "🇨🇳", "Zhuang", "Чжуанский язык", "Saɯ cueŋƅ"}
	case "zh":
		return Language{"zh", "🇨🇳", "Chinese", "Кита́йский язы́к", "中文"}
	// case "zh-Hant":
	// 	return Language{"zh-Hant", "🇨🇳", "ZH-HANT", "", ""}
	case "zu":
		return Language{"zu", "🇿🇦", "Zulu", "isiZulu", "Зу́лу"}
	default:
		return Language{"??", "🏳", "Unknown", "Неизвестный", ""}
	}
}
