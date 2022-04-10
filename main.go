package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Meta502/go-discord-boilerplate/src/constants"
	"github.com/Meta502/go-discord-boilerplate/src/env"
	"github.com/bwmarrin/discordgo"
)

func createMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := env.GetEnvironmentVariable("prefix")
	if m.Author.ID == s.State.User.ID {
		return
	}

	messagePrefix := m.Content[0:len(prefix)]
	message := m.Content[len(prefix):]
	args := strings.Split(message, " ")

	if messagePrefix != prefix {
		return
	}

	for _, element := range constants.Commands {
		command := args[0]

		_, ok := element.Alias[command]
		if ok || element.Name == command {
			element.Handler(s, m)
			break
		}
	}
}

func main() {
	token := env.GetEnvironmentVariable("token")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error initializing bot instance.")
	}

	dg.AddHandler(createMessageHandler)

	dg.Identify.Intents = discordgo.IntentDirectMessages | discordgo.IntentsGuildMessages

	err = dg.Open()

	if err != nil {
		log.Fatal(fmt.Errorf("Error opening connection to discord, %w", err))
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
