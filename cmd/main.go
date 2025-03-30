package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bismastr/discord-bot/internal/alert_cs_prices"
	"github.com/bismastr/discord-bot/internal/auth"
	"github.com/bismastr/discord-bot/internal/bot"
	"github.com/bismastr/discord-bot/internal/config"
	"github.com/bismastr/discord-bot/internal/db/cs_prices_db"
	"github.com/bismastr/discord-bot/internal/db/mabar_db"
	"github.com/bismastr/discord-bot/internal/firebase"
	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/bismastr/discord-bot/internal/handler"
	"github.com/bismastr/discord-bot/internal/llm"
	"github.com/bismastr/discord-bot/internal/messaging"
	"github.com/bismastr/discord-bot/internal/notification"
	"github.com/bismastr/discord-bot/internal/repository"
	"github.com/bismastr/discord-bot/internal/server"
	"github.com/bismastr/discord-bot/internal/user"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := mabar_db.NewDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Conn.Close()

	csPricesDb, err := cs_prices_db.NewDatabase()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	firebase, err := firebase.NewFirebaseClient(ctx)
	if err != nil {
		panic(err)
	}
	csRepository := repository.New(csPricesDb.Conn)
	repository := repository.New(db.Conn)

	dg, _ := discordgo.New(config.Envs.DiscordBotToken)
	discordBot := bot.NewBot(dg) //Discord bot init

	sessionStore := auth.NewSessionStore(auth.SessionOptions{
		CookiesKey: config.Envs.CookiesAuthSecret,
		MaxAge:     config.Envs.CookiesAuthAgeInSeconds,
		Secure:     config.Envs.CookiesAuthIsSecure,
		HttpOnly:   config.Envs.CookiesAuthIsHttpOnly,
	}) //Session for auth

	//Init all service
	gaming_session := gaming_session.NewGamingSessionService(repository)
	authService := auth.NewAuthService(sessionStore)
	botService := bot.NewBotService(discordBot.Dg)
	userService := user.NewUserService(repository)
	notificationService := notification.NewNotificationClient(firebase.Messaging)

	consumer, err := messaging.NewConsumer(config.Envs.RmqUrl)
	if err != nil {
		log.Fatal("Unable create consumer")
	}
	alertPriceService, err := alert_cs_prices.NewAlertPriceServcie(consumer, csRepository)
	if err != nil {
		log.Fatal("Unable create price service")
	}

	gemini := llm.NewGeminiClient(ctx)
	llmService := llm.NewLlmService(gemini)

	//Start Discord
	botHandler := bot.NewActionHandlerCtrl(userService, gaming_session, botService, llmService, alertPriceService, ctx)

	discordBot.RegisterHandler(botHandler)
	discordBot.Open()
	discordBot.AddAllCommand()
	close, _ := botHandler.DailyScheduleSummary()

	//Start server
	handler := handler.NewHandler(botService, authService, userService, gaming_session, notificationService)
	server := server.NewServer(gin.Default(), discordBot.Dg)
	server.RegisterRoutes(handler)
	server.Start()

	defer close()
	exit(dg)
	defer dg.Close()
}

func exit(dg *discordgo.Session) {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
