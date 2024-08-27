package components

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func JoinSession(s *discordgo.Session, i *discordgo.InteractionCreate, userId string, memberMentioned string) {
	messageContent := fmt.Sprintf("<@%v> join abang quh 🥳\n\nArek-arek sing join 👥:%v  ", userId, memberMentioned)
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
