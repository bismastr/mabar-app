package bot

import (
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
		{
			Name:        "buyar-sek",
			Description: "Hapus Sesi Mabar",
		},
		{
			Name:        "list-mabar",
			Description: "Melihat list of gaming Session ",
		},
	}
)

func (b *Bot) AddAllCommand() {
	if b.dg == nil {
		log.Panic("dg is nil")
	}

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := b.dg.ApplicationCommandCreate(b.dg.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}
