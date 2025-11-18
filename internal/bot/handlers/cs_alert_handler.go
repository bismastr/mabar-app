package handlers

import (
	"context"
	"log"
	"strconv"

	"github.com/bismastr/discord-bot/internal/alert_cs_prices"
	"github.com/bismastr/discord-bot/internal/bot/components/message_components"
	"github.com/bismastr/discord-bot/internal/repository"
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5/pgtype"
)

type CSAlertHandler struct {
	alertService *alert_cs_prices.AlertPriceSertvice
	ctx          context.Context
}

func NewCSAlertHandler(alertService *alert_cs_prices.AlertPriceSertvice, ctx context.Context) *CSAlertHandler {
	return &CSAlertHandler{
		alertService: alertService,
		ctx:          ctx,
	}
}

func (h *CSAlertHandler) Name() string {
	return "create-daily-cs-alert"
}

func (h *CSAlertHandler) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	userId, _ := strconv.ParseInt(i.Member.User.ID, 10, 64)

	itemIDStr := i.ApplicationCommandData().Options[0].StringValue()
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		log.Println("Invalid item ID")
		return err
	}

	err = h.alertService.AddDailySchedule(h.ctx, repository.InsertAlertDailyScheduleParams{
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
		return err
	}

	content := "## Successfully added daily schedule summary"
	message_components.SendMessage(s, i, content)
	return nil
}
