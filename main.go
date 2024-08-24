package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bismastr/discord-bot/handlers/basic"
	gamingSessionHandler "github.com/bismastr/discord-bot/handlers/gaming-session"
	"github.com/bismastr/discord-bot/utils"
	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	//Handler
	dg.AddHandler(basic.PingPongMessage)
	dg.AddHandler(gamingSessionHandler.AddGamingSessionCommandData)

	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	//Socket Open
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	//Add all command
	utils.AddAllCommand(dg)

	//Exit channel
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
