package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/bismastr/discord-bot/internal/alert_cs_prices"
	"github.com/bwmarrin/discordgo"
)

type CSAutocompleteHandler struct {
	alertService *alert_cs_prices.AlertPriceSertvice
	ctx          context.Context
}

func NewCSAutocompleteHandler(alertService *alert_cs_prices.AlertPriceSertvice, ctx context.Context) *CSAutocompleteHandler {
	return &CSAutocompleteHandler{
		alertService: alertService,
		ctx:          ctx,
	}
}

func (h *CSAutocompleteHandler) Name() string {
	return "create-daily-cs-alert"
}

func (h *CSAutocompleteHandler) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) error {
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

	itemsPtr, err := h.alertService.GetItemsContainsName(h.ctx, searchQuery)
	if err != nil {
		log.Printf("error getting items: %v", err)
		return err
	}

	if itemsPtr == nil {
		return nil
	}

	items := *itemsPtr
	choices := make([]*discordgo.ApplicationCommandOptionChoice, 0, len(items))
	for _, item := range items {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  item.Name,
			Value: fmt.Sprintf("%d", item.ID),
		})
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: choices,
		},
	})
}
