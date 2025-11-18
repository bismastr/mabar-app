package bot

import "github.com/bwmarrin/discordgo"

// InteractionHandler defines the interface for handling Discord interactions
type InteractionHandler interface {
	Handle(s *discordgo.Session, i *discordgo.InteractionCreate) error
}

// CommandHandler handles slash commands
type CommandHandler interface {
	InteractionHandler
	Name() string
}

// ComponentHandler handles message components (buttons, select menus)
type ComponentHandler interface {
	InteractionHandler
	CustomIDPrefix() string
}

// AutocompleteHandler handles autocomplete interactions
type AutocompleteHandler interface {
	InteractionHandler
	Name() string
}
