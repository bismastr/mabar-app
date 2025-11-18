package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type DeclineSessionHandler struct{}

func NewDeclineSessionHandler() *DeclineSessionHandler {
	return &DeclineSessionHandler{}
}

func (h *DeclineSessionHandler) CustomIDPrefix() string {
	return "mabar_no"
}

func (h *DeclineSessionHandler) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	userid := i.Member.User.ID
	noJoin := fmt.Sprintf("<@%v> tidak join duls, kecewaaaa sangat berat!", userid)

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: noJoin,
		},
	})
}
