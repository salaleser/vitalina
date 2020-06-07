package command

import (
	"fmt"
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

		builder.WriteString("Список поддерживаемых команд (префикс `~`):\n")
		for title, description := range commands {
			builder.WriteString(fmt.Sprintf(""+
				"`%s` — %s\n",
				title,
				description,
			))
		}
		builder.WriteString("\n")

		text := "*Я переписываю бота на другой язык программирования. Я собираюсь " +
			"восстановить большую часть старых команд в ближайшее время. Вопросы можно " +
			"задать @salaleser#7570*\n\n" +
			builder.String() +
			"(Usage) Примеры использования (попробуйте просто отправить такое сообщение):\n" +
			"`143444`, `6014 CN`, `547702041 JP ja-jp`\n\n" +
			"(Aliases) Все префиксы (алиасы для скрытых команд):\n" +
			"`§` — search-apps (App Store, Google Play): `§майнкрафт`.\n" +
			"`?` — force command (verbose mode): `?RU`, `?hu-hu`.\n" +
			"`.` — detect language (by reactions): `.Աստղային ճանապարհ`\n" +
			"`,` — detect language (by reply message) `,HeghluʼmeH QaQ jajvam`.\n"

		util.SendInfo(s, m, text)
	}
}

// InitHelp initializes command. (FIXME)
func InitHelp(c map[string]string) {
	commands = c
}
