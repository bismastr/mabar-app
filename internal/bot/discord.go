package bot

import (
	"fmt"

	"github.com/bismastr/discord-bot/internal/database"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Dg       *discordgo.Session
	database *database.DbClient
}

func NewBot(dg *discordgo.Session, database *database.DbClient) *Bot {
	return &Bot{
		Dg:       dg,
		database: database,
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
