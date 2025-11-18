package bot

import (
	"context"
	"log"
	"strings"

	"github.com/bismastr/discord-bot/internal/bot/handlers"
	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/bismastr/discord-bot/internal/llm"
	"github.com/bismastr/discord-bot/internal/user"
	"github.com/bwmarrin/discordgo"
)

type HandlerRegistry struct {
	commands      map[string]CommandHandler
	components    map[string]ComponentHandler
	autocompletes map[string]AutocompleteHandler
}

func NewHandlerRegistry() *HandlerRegistry {
	return &HandlerRegistry{
		commands:      make(map[string]CommandHandler),
		components:    make(map[string]ComponentHandler),
		autocompletes: make(map[string]AutocompleteHandler),
	}
}

func (r *HandlerRegistry) RegisterCommand(handler CommandHandler) {
	r.commands[handler.Name()] = handler
}

func (r *HandlerRegistry) RegisterComponent(handler ComponentHandler) {
	r.components[handler.CustomIDPrefix()] = handler
}

func (r *HandlerRegistry) RegisterAutocomplete(handler AutocompleteHandler) {
	r.autocompletes[handler.Name()] = handler
}

func (b *Bot) RegisterHandlers(registry *HandlerRegistry) {
	b.Dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		b.handleInteraction(registry, s, i)
	})
}

func (b *Bot) handleInteraction(registry *HandlerRegistry, s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		commandName := i.ApplicationCommandData().Name
		if handler, ok := registry.commands[commandName]; ok {
			if err := handler.Handle(s, i); err != nil {
				log.Printf("Error handling command %s: %v", commandName, err)
			}
		}
	case discordgo.InteractionMessageComponent:
		customID := i.MessageComponentData().CustomID
		prefix := extractPrefix(customID)
		if handler, ok := registry.components[prefix]; ok {
			if err := handler.Handle(s, i); err != nil {
				log.Printf("Error handling component %s: %v", prefix, err)
			}
		}
	case discordgo.InteractionApplicationCommandAutocomplete:
		commandName := i.ApplicationCommandData().Name
		if handler, ok := registry.autocompletes[commandName]; ok {
			if err := handler.Handle(s, i); err != nil {
				log.Printf("Error handling autocomplete %s: %v", commandName, err)
			}
		}
	}
}

func SetupHandlers(
	bot *Bot,
	userService *user.UserService,
	sessionService *gaming_session.GamingSessionService,
	botService *BotService,
	llmService *llm.LlmService,
	ctx context.Context,
) {
	registry := NewHandlerRegistry()

	registry.RegisterCommand(handlers.NewCreateMabarHandler(userService, sessionService, botService, ctx))
	registry.RegisterCommand(handlers.NewAIHandler(llmService, ctx))
	// registry.RegisterCommand(handlers.NewCSAlertHandler(alertService, ctx))

	registry.RegisterComponent(handlers.NewJoinSessionHandler(userService, sessionService, ctx))
	registry.RegisterComponent(handlers.NewDeclineSessionHandler())

	// registry.RegisterAutocomplete(handlers.NewCSAutocompleteHandler(alertService, ctx))

	bot.RegisterHandlers(registry)
}

func extractPrefix(customID string) string {
	parts := strings.Split(customID, "_")
	if len(parts) >= 2 {
		return parts[0] + "_" + parts[1]
	}
	return customID
}
