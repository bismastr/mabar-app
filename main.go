package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bismastr/discord-bot/internal/bot"
	"github.com/bismastr/discord-bot/internal/database"
	"github.com/bismastr/discord-bot/internal/server"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	ctx := context.Background()

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	//Start Database
	serverFirebaseClient, _ := database.NewFirebaseClient(ctx)
	//Start Bot
	bot := bot.NewBot(dg, serverFirebaseClient)
	bot.RegisterHandler()
	bot.Open()
	bot.AddAllCommand()
	//Start Server
	server := server.NewServer(gin.Default(), serverFirebaseClient, bot.Dg)
	server.RegisterRoutes()
	server.Start()

	exit(dg)
	defer serverFirebaseClient.Client.Close()
}

func exit(dg *discordgo.Session) {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
