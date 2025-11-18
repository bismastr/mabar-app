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

	gemini := llm.NewGeminiClient(ctx)
	llmService := llm.NewLlmService(gemini)

	//Start Discord
	bot.SetupHandlers(discordBot, userService, gaming_session, botService, llmService, ctx)
	discordBot.Open()
	discordBot.AddAllCommand()

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
			var dailySummary alert_cs_prices.NotificationPriceSummary
			err := json.Unmarshal(d.Body, &dailySummary)
			if err != nil {
				log.Println("Error unmarshaling daily report:", err)
				continue
			}

			report := fmt.Sprintf("📊 **DAILY SUMMARY** <@%d> 📊 **FOR %s** \n", dailySummary.DiscordId, dailySummary.ItemName)
			report += "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"
			report += fmt.Sprintf("🟢 **Open**:   $%.2f\n", dailySummary.OpeningPrice/100)
			report += fmt.Sprintf("🔴 **Close**:  $%.2f\n", dailySummary.ClosingPrice/100)
			report += fmt.Sprintf("🔺 **High**:    $%.2f\n", dailySummary.MaxPrice/100)
			report += fmt.Sprintf("🔻 **Low**:     $%.2f\n", dailySummary.MinPrice/100)
			report += fmt.Sprintf("📌 **Avg**:     $%.2f\n", dailySummary.AvgPrice/100)
			report += fmt.Sprintf("📈 **Change**: %.2f%%\n", dailySummary.ChangePct)
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
