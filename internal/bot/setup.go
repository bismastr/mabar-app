package bot

import (
	"context"

	"github.com/bismastr/discord-bot/internal/bot/handlers"
	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/bismastr/discord-bot/internal/llm"
	"github.com/bismastr/discord-bot/internal/user"
)

// SetupHandlers initializes and registers all bot handlers
func SetupHandlers(
	bot *Bot,
	userService *user.UserService,
	sessionService *gaming_session.GamingSessionService,
	botService *BotService,
	llmService *llm.LlmService,
	ctx context.Context,
) {
	registry := NewHandlerRegistry()

	// Register command handlers
	registry.RegisterCommand(handlers.NewCreateMabarHandler(userService, sessionService, botService, ctx))
	registry.RegisterCommand(handlers.NewAIHandler(llmService, ctx))
	// registry.RegisterCommand(handlers.NewCSAlertHandler(alertService, ctx))

	// Register component handlers
	registry.RegisterComponent(handlers.NewJoinSessionHandler(userService, sessionService, ctx))
	registry.RegisterComponent(handlers.NewDeclineSessionHandler())

	// Register autocomplete handlers
	// registry.RegisterAutocomplete(handlers.NewCSAutocompleteHandler(alertService, ctx))

	bot.RegisterHandlers(registry)
}
