package message_components

import (
	"fmt"

	"github.com/bismastr/discord-bot/internal/gaming_session"
	"github.com/bwmarrin/discordgo"
)

func JoinSessionV2(s *discordgo.Session, i *discordgo.InteractionCreate, userId int64, gamingSession *gaming_session.GetGamingSessionResponse) {
	var memberMentioned string
	for _, user := range gamingSession.Users {
		memberMentioned += fmt.Sprintf("<@%v>", user.DiscordUid.Int64)
	}

	messageContent := fmt.Sprintf("<@%d> join abang quh 🥳\n\nPlayers  👥:%v  ", userId, memberMentioned)
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: messageContent,
		},
	})
	if err != nil {
		panic(err)
	}
}
