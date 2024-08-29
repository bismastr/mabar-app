package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bismastr/discord-bot/db"
	gamingSessionHandler "github.com/bismastr/discord-bot/handlers/gaming-session"
	"github.com/bismastr/discord-bot/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	fireBaseClient := db.NewFirebaseClient(ctx)

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	//Handler
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		gamingSessionHandler.AddGamingSessionCommandData(s, i, fireBaseClient, ctx)
	})

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

	// Close when exit
	defer fireBaseClient.Client.Close()
	dg.Close()
}
