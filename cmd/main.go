package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bismastr/discord-bot/internal/auth"
	"github.com/bismastr/discord-bot/internal/bot"
	"github.com/bismastr/discord-bot/internal/config"
	"github.com/bismastr/discord-bot/internal/database"
	"github.com/bismastr/discord-bot/internal/gamingSession"
	"github.com/bismastr/discord-bot/internal/handler"
	"github.com/bismastr/discord-bot/internal/server"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	serverFirebaseClient, _ := database.NewFirebaseClient(ctx) //Database init

	dg, _ := discordgo.New(config.Envs.DiscordBotToken)
	discordBot := bot.NewBot(dg, serverFirebaseClient) //Discord bot init
	discordBot.RegisterHandler()
	discordBot.Open()
	discordBot.AddAllCommand()

	sessionStore := auth.NewSessionStore(auth.SessionOptions{
		CookiesKey: config.Envs.CookiesAuthSecret,
		MaxAge:     config.Envs.CookiesAuthAgeInSeconds,
		Secure:     config.Envs.CookiesAuthIsSecure,
		HttpOnly:   config.Envs.CookiesAuthIsHttpOnly,
	}) //Session for auth

	authService := auth.NewAuthService(sessionStore)                                                                     //Auth service
	botService := bot.NewBotGamingSessionService(gamingSession.NewRepositoryImpl(serverFirebaseClient), discordBot.Dg)   //Bot service
	gamingSessionService := gamingSession.NewGamingSessionService(gamingSession.NewRepositoryImpl(serverFirebaseClient)) //gaming session service

	handler := handler.NewHandler(botService, gamingSessionService, authService)
	server := server.NewServer(gin.Default(), serverFirebaseClient, discordBot.Dg) //Server
	server.RegisterRoutes(handler)
	server.Start()

	exit(dg)
	defer serverFirebaseClient.Client.Close()
	defer dg.Close()
}

func exit(dg *discordgo.Session) {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
