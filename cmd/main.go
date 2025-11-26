package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bismastr/discord-bot/internal/alert_cs_prices"
	"github.com/bismastr/discord-bot/internal/auth"
	"github.com/bismastr/discord-bot/internal/bot"
	"github.com/bismastr/discord-bot/internal/config"
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

	ctx := context.Background()

	firebase, err := firebase.NewFirebaseClient(ctx)
	if err != nil {
		panic(err)
	}

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
	messagingConsumer, _ := messaging.NewConsumer(config.Envs.RmqUrl)
	alertService, _ := alert_cs_prices.NewAlertPriceServcie(messagingConsumer, repository)

	gemini := llm.NewGeminiClient(ctx)
	llmService := llm.NewLlmService(gemini)

	//Start Discord
	bot.SetupHandlers(discordBot, userService, gaming_session, botService, llmService, ctx)
	discordBot.Open()
	discordBot.AddAllCommand()
	setupDailySummary(alertService, botService)

	//Start server
	handler := handler.NewHandler(botService, authService, userService, gaming_session, notificationService)
	server := server.NewServer(gin.Default(), discordBot.Dg)
	server.RegisterRoutes(handler)
	server.Start()
	exit(dg)
	defer dg.Close()
}

func setupDailySummary(alertService *alert_cs_prices.AlertPriceSertvice, botService *bot.BotService) (func(), error) {
	log.Println("Setting up daily report summary")
	msgs, close, err := alertService.DailyReportSummary()
	if err != nil {
		log.Println("Error setting up daily report")
		return nil, err
	}

	go func() {
		for d := range msgs {
			log.Printf("Received message body: %s", string(d.Body))

			var dailySummary alert_cs_prices.NotificationPriceSummary
			err := json.Unmarshal(d.Body, &dailySummary)
			if err != nil {
				log.Printf("Error unmarshaling daily report: %v | Raw body: %s", err, string(d.Body))
				continue
			}

			var alertEmoji, trendEmoji, colorCode string
			if dailySummary.AlertType == "INCREASE" {
				alertEmoji = "📈"
				trendEmoji = "🟢"
				colorCode = "```ansi\n\u001b[1;32m" // Green
			} else {
				alertEmoji = "📉"
				trendEmoji = "🔴"
				colorCode = "```ansi\n\u001b[1;31m" // Red
			}

			report := fmt.Sprintf("%s **24H PRICE ALERT** - %s %s\n", trendEmoji, dailySummary.Name, trendEmoji)
			report += "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"
			report += colorCode
			report += fmt.Sprintf("%s Trend: %s\n", alertEmoji, dailySummary.AlertType)
			report += fmt.Sprintf("💰 Change: %.2f%%\n", dailySummary.ChangePct)
			report += fmt.Sprintf("🆕 Latest Price: $%.2f	\n", dailySummary.LatestSellPrice)
			report += fmt.Sprintf("🕰️ Old Price: $%.2f\n", dailySummary.OldSellPrice)
			report += "\u001b[0m```\n"
			report += "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"

			botService.SendMessageToChannel("1348197711773564949", report)
		}
	}()

	return close, nil
}

func exit(dg *discordgo.Session) {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
