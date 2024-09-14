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
					Name:        "nama-permainan",
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
	if b.dg == nil {
		log.Panic("dg is nil")
	}
	//Delete Existing Command
	b.dg.ApplicationCommandDelete(b.dg.State.Application.ID, "", "")
	//Add new command
	_, err := b.dg.ApplicationCommandBulkOverwrite(b.dg.State.User.ID, "", commands)
	if err != nil {
		fmt.Printf("Error creating commands: %v", err)
	}
}
