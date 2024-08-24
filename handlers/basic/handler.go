package basic

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func PingPongMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "hello" {
		fmt.Println("Send msg in server")
		s.ChannelMessageSend(m.ChannelID, "World!")
	}

	if m.Content == "pong" {
		fmt.Println("Send msg in server")
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
