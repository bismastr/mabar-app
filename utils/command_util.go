package utils

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "create-mabar",
			Description: "Buat Sesi Mabar Baru",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "game-name",
					Description: "Nama Game",
					Required:    false,
				},
			},
		},
		{
			Name:        "buyar-sek",
			Description: "Hapus Sesi Mabar",
		},
	}
)

func AddAllCommand(dg *discordgo.Session) {
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.Application.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}
