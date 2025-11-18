package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/bismastr/discord-bot/internal/llm"
	"github.com/bwmarrin/discordgo"
)

type AIHandler struct {
	llmService *llm.LlmService
	ctx        context.Context
}

func NewAIHandler(llmService *llm.LlmService, ctx context.Context) *AIHandler {
	return &AIHandler{
		llmService: llmService,
		ctx:        ctx,
	}
}

func (h *AIHandler) Name() string {
	return "ask-ai"
}

func (h *AIHandler) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) error {
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
		return err
	}

	go func() {
		resp, err := h.llmService.GetGenerateResponse(h.ctx, question)
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
				log.Printf("Error sending response: %v", err)
				break
			}
		}
	}()

	return nil
}
