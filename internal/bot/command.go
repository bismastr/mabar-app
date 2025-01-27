package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "create-mabar",
			Description: "Buat sesi mabar baru",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "game_name",
					Description: "Choose your game",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "Counter-Strike 2",
							Value: 1,
						},
						{
							Name:  "Deadlock",
							Value: 2,
						},
						{
							Name:  "Valorant",
							Value: 3,
						},
						{
							Name:  "GTA V",
							Value: 4,
						},
					},
				},
			},
		},
	}
)

func (b *Bot) AddAllCommand() {
	if b.Dg == nil {
		log.Panic("dg is nil")
	}

	err := b.Dg.ApplicationCommandDelete(b.Dg.State.Application.ID, "", "")
	if err != nil {
		fmt.Printf("Error creating commands: %v", err)
	}

	_, err = b.Dg.ApplicationCommandBulkOverwrite(b.Dg.State.Application.ID, "", commands)
	if err != nil {
		fmt.Printf("Error creating commands: %v", err)
	}
}
