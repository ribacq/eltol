package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const (
	token     = "NDM5ODIwNDMzNzA1OTI2NjY2.DcYuZA.vz03YxlcLFM5M6f6hEWqK1KR6x4"
	botLandID = "436699658019143682"
)

func main() {
	// Load token
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// When a new message is sent
	bot.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// ignore my own messages
		if m.Author.ID == s.State.User.ID {
			return
		}

		// react to a message containing but not beginning with & with an mushroom emoji, because yeah!
		if !strings.HasPrefix(m.Content, "&") && strings.Contains(m.Content, "&") {
			err := s.MessageReactionAdd(m.ChannelID, m.ID, ":mushroom:")
			if err != nil {
				fmt.Println("Unable to react to &:", err)
			}
		}

		// commands
		args := strings.Split(m.Content, " ")
		if !strings.HasPrefix(args[0], "&") {
			return
		}

		switch args[0] {
		case "&ping":
			_, err := s.ChannelMessageSend(m.ChannelID, "Pong !")
			if err != nil {
				fmt.Println("Unable to send message:", err)
			}
		case "&amp;":
			_, err := s.ChannelMessageSend(m.ChannelID, m.Author.Username+" a tout compris.")
			if err != nil {
				fmt.Println("Unable to send message:", err)
			}
		}
	})

	// When joining a channel
	bot.AddHandler(func(s *discordgo.Session, c *discordgo.ChannelCreate) {
		if c.ID != botLandID {
			fmt.Println("Not #bot_land")
			return
		}

		_, err := s.ChannelMessageSend(c.ID, "Bordel de foutre-couille de merde !")
		if err != nil {
			fmt.Println("Unable to send message:", err)
		}
	})

	// When connecting to the Discord network
	bot.AddHandler(func(s *discordgo.Session, _ *discordgo.Connect) {
		fmt.Println("Connected.")
	})

	// Try opening connection
	err = bot.Open()
	if err != nil {
		fmt.Println("Open error:", err)
		return
	}

	// Wait for ^C
	fmt.Println("Eltol is flying!")
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// End of run
	bot.Close()
}
