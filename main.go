package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bismastr/discord-bot/db"
	gamingSessionHandler "github.com/bismastr/discord-bot/handlers/gaming-session"
	"github.com/bismastr/discord-bot/internal/server"
	"github.com/bismastr/discord-bot/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

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

	server := server.NewServer(fireBaseClient)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
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
