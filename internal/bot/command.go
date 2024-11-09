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
			Description: "Buat sesi mabar baru. Kamu bisa tambahkan nama permainan/game (opsional)",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "game-name-1",
					Description: "Nama Permainan/Game",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "game-name-2",
					Description: "Nama Permainan/Game",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "game-name-3",
					Description: "Nama Permainan/Game",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "game-name-4",
					Description: "Nama Permainan/Game",
					Required:    false,
				},
			},
		},

		// {
		// 	Name:        "buyar-sek",
		// 	Description: "Hapus Sesi Mabar",
		// },
		// {
		// 	Name:        "list-mabar",
		// 	Description: "Melihat list of gaming Session ",
		// },
	}
)

func (b *Bot) AddAllCommand() {
	if b.Dg == nil {
		log.Panic("dg is nil")
	}
	//Delete Existing Command
	err := b.Dg.ApplicationCommandDelete(b.Dg.State.Application.ID, "", "")
	if err != nil {
		fmt.Printf("Error creating commands: %v", err)
	}
	//Add new command
	_, err = b.Dg.ApplicationCommandBulkOverwrite(b.Dg.State.Application.ID, "", commands)
	if err != nil {
		fmt.Printf("Error creating commands: %v", err)
	}
}
