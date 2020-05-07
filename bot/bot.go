package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/salaleser/vitalina/command"
	"github.com/salaleser/vitalina/util"
)

type handler func(s *discordgo.Session, m *discordgo.MessageCreate)

// Command is a command
type Command struct {
	Hanlder     handler
	Description string
}

var Commands = map[string]Command{
	"help": {command.Help, "Помощь"},
	"play": {command.Play, "(Play)"},
	"tts":  {command.Tts, "(TTS)"},
}

// Start starts the bot
func Start(token string) {
	dg, err := discordgo.New(token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		os.Exit(1)
	}

	dg.Debug = false

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		os.Exit(1)
	}

	fmt.Println("Vitalina is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	guilds, _ := s.UserGuilds(100, "", "")
	fmt.Printf("Connected to %d guilds:\n", len(guilds))
	// for _, g := range guilds {
	// 	fmt.Printf("%s (%s)\n", g.Name, g.ID)
	// }

	c := make(map[string]string)
	for k, v := range Commands {
		c[k] = v.Description
	}
	command.InitHelp(c) // FIXME workaround cycle imports

	s.UpdateStatus(0, "~help")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	args := strings.Split(m.Content, " ")

	guild, err := s.Guild(m.GuildID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get guild: %v\n", err)
	}

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get channel: %v\n", err)
	}

	if err == nil {
		log.Printf("*%s #%s @%s: %q", guild.Name, channel.Name, m.Author.Username, m.Message.Content)
	}

	if !strings.HasPrefix(m.Content, "~") {
		go command.Vitalina(s, m)
		return
	}

	if cmd, ok := Commands[args[0][1:]]; ok {
		go cmd.Hanlder(s, m)
	} else {
		util.SendInfo(s, m, fmt.Sprintf("Команда `%s` не поддерживается", args[0]))
	}
}
