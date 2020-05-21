package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/salaleser/vitalina/util"
)

var commands map[string]string

// Help sends user manual.
func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content[1:], " ")

	if len(args) == 1 {
		var builder strings.Builder

		builder.WriteString("Список поддерживаемых комад (префикс `~`):\n")
		for title, description := range commands {
			builder.WriteString("`" + title + "`")
			builder.WriteString(" — ")
			builder.WriteString(description)
			builder.WriteByte('\n')
		}

		text := "*Я переписываю бота на другой язык программирования. Я собираюсь " +
			"восстановить большую часть старых команд в ближайшее время. Вопросы можно " +
			"задать @salaleser#7570*\n\n" +
			builder.String() +
			"\nКроме того читает ID приложений App Store и Google Play (Steam будет " +
			"добавлен) в тексте сообщений. А также текст, начинающийся с " +
			"восклицательного знака воспринимает как команду для поиска в магазине " +
			"приложений по ключевому запросу. Например: `!майнкрафт`." +
			"\nВсе префиксы (алиасы для скрытых команд):" +
			"\n`!` — search-apps (App Store, Google Play)" +
			"\n`?` — force command (verbose mode)" +
			"\n`.` — detect language (by reactions)" +
			"\n`,` — detect language (by reply message)."

		util.SendInfo(s, m, text)
	}
}

// InitHelp initializes command. (FIXME)
func InitHelp(c map[string]string) {
	commands = c
}
