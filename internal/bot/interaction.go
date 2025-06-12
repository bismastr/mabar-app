package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bismastr/discord-bot/internal/alert_cs_prices"
	"github.com/bismastr/discord-bot/internal/bot/components/message_components"
	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/bismastr/discord-bot/internal/llm"
	"github.com/bismastr/discord-bot/internal/repository"
	"github.com/bismastr/discord-bot/internal/user"
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ActionHandlerCtrl struct {
	userService    *user.UserService
	gamingSession  *gaming_session.GamingSessionService
	BotService     *BotService
	llmService     *llm.LlmService
	alertCsService *alert_cs_prices.AlertPriceSertvice
	ctx            context.Context
}

func NewActionHandlerCtrl(
	userService *user.UserService,
	gamingSession *gaming_session.GamingSessionService,
	botService *BotService,
	llmService *llm.LlmService,
	alertCsService *alert_cs_prices.AlertPriceSertvice,
	ctx context.Context) *ActionHandlerCtrl {
	return &ActionHandlerCtrl{
		userService:    userService,
		gamingSession:  gamingSession,
		BotService:     botService,
		ctx:            ctx,
		llmService:     llmService,
		alertCsService: alertCsService,
	}
}

func (a *ActionHandlerCtrl) DailyScheduleSummary() (func(), error) {
	log.Println("Daily report summary ")
	msgs, close, err := a.alertCsService.DailyReportSummary()
	if err != nil {
		log.Println("Error daily report")
		return nil, err
	}
	go func() {
		for d := range msgs {
			var dailySummary alert_cs_prices.NotificationPriceSummary
			err := json.Unmarshal(d.Body, &dailySummary)
			if err != nil {
				log.Println("Error daily report")
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

			a.BotService.SendMessageToChannel("1276782792876888075", report)
		}
	}()

	return close, nil
}

func (a *ActionHandlerCtrl) CreateSchedulerCsItems(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userId, _ := strconv.ParseInt(i.Member.User.ID, 10, 64)

	// Convert the string value to an integer
	itemIDStr := i.ApplicationCommandData().Options[0].StringValue()
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		log.Println("Invalid item ID")
		return
	}

	content := "## Succesfully add daily schedule summary"
	err = a.alertCsService.AddDailySchedule(a.ctx, repository.InsertAlertDailyScheduleParams{
		ItemID: pgtype.Int4{
			Int32: int32(itemID),
			Valid: true,
		},
		DiscordID: pgtype.Int8{
			Int64: userId,
			Valid: true,
		},
	})
	if err != nil {
		log.Println("Cannot insert alert")
	}

	message_components.SendMessage(s, i, content)
}

func (a *ActionHandlerCtrl) CsItemsAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	var focusedOption *discordgo.ApplicationCommandInteractionDataOption
	for _, option := range data.Options {
		if option.Focused {
			focusedOption = option
			break
		}
	}

	searchQuery := ""
	if focusedOption != nil {
		searchQuery = focusedOption.StringValue()
	}

	itemsPtr, err := a.alertCsService.GetItemsContainsName(a.ctx, searchQuery)
	if err != nil {
		log.Printf("error getting: %d", err)
	}

	if itemsPtr == nil {
		return
	}

	items := *itemsPtr

	choices := make([]*discordgo.ApplicationCommandOptionChoice, 0, len(items))
	for _, item := range items {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  item.Name,
			Value: fmt.Sprintf("%d", item.ID),
		})
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: choices,
		},
	})

	if err != nil {
		log.Printf("Error responding to autocomplete: %v", err)
	}

}

func (a *ActionHandlerCtrl) GenerateContent(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	username := i.Member.User.Username
	var question string
	for _, option := range options {
		if option.Name == "question" {
			question = option.StringValue()
			break
		}
	}

	content := fmt.Sprintf("@%s Asking: %s", username, question)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		log.Printf("Error getting response: %v", err)
	}

	go func() {
		resp, err := a.llmService.GetGenerateResponse(a.ctx, question)
		if err != nil {
			log.Printf("Error getting response: %v", err)
			s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: "Unable to generate LLM response",
			})
			return
		}
		s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: content,
		})

		for _, part := range resp {
			_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: part,
			})
			if err != nil {
				log.Printf("Error getting response: %v", err)
				break
			}
		}
	}()
}

func (a *ActionHandlerCtrl) JoinGamingSessionV2(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userId, _ := strconv.ParseInt(i.Member.User.ID, 10, 64)
	customId := i.MessageComponentData().CustomID
	split := strings.Split(customId, "_")
	id, _ := strconv.ParseInt(split[2], 10, 64)

	user, err := a.userService.GetUserByDiscordUID(a.ctx, userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			message_components.NeedLoginMessage(s, i)
		} else {
			message_components.ErrorMessage(s, i)
		}
		return
	}

	err = a.gamingSession.InsertUserJoinSession(a.ctx, user.ID, id)
	if err != nil {
		message_components.ErrorMessage(s, i)
	}

	response, err := a.gamingSession.GetGamingSessionById(a.ctx, id)
	if err != nil {
		message_components.ErrorMessage(s, i)
	}

	message_components.JoinSessionV2(s, i, userId, response)
}

func (a *ActionHandlerCtrl) DeclineGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userid := i.Member.User.ID
	noJoin := fmt.Sprintf("<@%v> tidak join duls, kecewaaaa sangat berat!", userid)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: noJoin,
		},
	})
	if err != nil {
		panic(err)
	}
}

func (a *ActionHandlerCtrl) CreateMabar(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userId, _ := strconv.ParseInt(i.Member.User.ID, 10, 64)

	user, err := a.userService.GetUserByDiscordUID(a.ctx, userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			message_components.NeedLoginMessage(s, i)
		} else {
			message_components.ErrorMessage(s, i)
		}
		return
	}

	createSession, err := a.gamingSession.CreateGamingSession(a.ctx, &gaming_session.CreateGamingSessionRequest{
		IsFinish: pgtype.Bool{
			Bool:  false,
			Valid: true,
		},
		CreatedBy: user.ID,
		GameID:    i.ApplicationCommandData().Options[0].IntValue(),
	})
	if err != nil {
		message_components.ErrorMessage(s, i)
	}

	session, err := a.gamingSession.GetGamingSessionById(a.ctx, createSession.ID)
	if err != nil {
		message_components.ErrorMessage(s, i)
	}

	_, err = a.BotService.CreateGamingSession(session, i.ChannelID)
	if err != nil {
		message_components.ErrorMessage(s, i)
	}
}
