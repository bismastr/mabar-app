package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Dg *discordgo.Session
}

func NewBot(dg *discordgo.Session) *Bot {
	return &Bot{
		Dg: dg,
	}
}

func (b *Bot) Open() {
	b.Dg.Identify.Intents = discordgo.IntentsAll
	err := b.Dg.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
}
