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
	"github.com/bismastr/discord-bot/internal/db"
	"github.com/bismastr/discord-bot/internal/gamingSession"
	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/bismastr/discord-bot/internal/handler"
	"github.com/bismastr/discord-bot/internal/repository"
	"github.com/bismastr/discord-bot/internal/server"
	"github.com/bismastr/discord-bot/internal/user"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	db, err := db.NewDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Conn.Close(ctx)
	repository := repository.New(db.Conn)

	serverFirebaseClient, _ := database.NewFirebaseClient(ctx) //Database init

	dg, _ := discordgo.New(config.Envs.DiscordBotToken)
	discordBot := bot.NewBot(dg, serverFirebaseClient) //Discord bot init

	sessionStore := auth.NewSessionStore(auth.SessionOptions{
		CookiesKey: config.Envs.CookiesAuthSecret,
		MaxAge:     config.Envs.CookiesAuthAgeInSeconds,
		Secure:     config.Envs.CookiesAuthIsSecure,
		HttpOnly:   config.Envs.CookiesAuthIsHttpOnly,
	}) //Session for auth

	//Init all service
	gaming_session := gaming_session.NewGamingSessionService(repository)
	authService := auth.NewAuthService(sessionStore)                                                                     //Auth service
	botService := bot.NewBotService(discordBot.Dg)                                                                       //Bot service
	gamingSessionService := gamingSession.NewGamingSessionService(gamingSession.NewRepositoryImpl(serverFirebaseClient)) //gaming session service
	userService := user.NewUserService(repository)

	//Start Discord
	botHandler := bot.NewActionHandlerCtrl(gamingSessionService, userService, gaming_session, context.Background())
	discordBot.RegisterHandler(botHandler)
	discordBot.Open()
	discordBot.AddAllCommand()

	//Start server
	handler := handler.NewHandler(botService, gamingSessionService, authService, userService, gaming_session)
	server := server.NewServer(gin.Default(), serverFirebaseClient, discordBot.Dg)
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
