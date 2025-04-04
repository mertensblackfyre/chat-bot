package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)
func Discord() {
	dg, err := discordgo.New("Bot " + DISCORD_KEY)

	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		SendMessages(s, m)
	})

	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = dg.Open()

	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	fmt.Println("--AI is online--")

	defer dg.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}

func SendMessages(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content != "" {
		message := Gemini(m.Content)
		s.ChannelMessageSend(m.ChannelID, message)
	}

}
