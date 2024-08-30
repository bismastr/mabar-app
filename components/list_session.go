package components

import (
	"fmt"

	"github.com/bismastr/discord-bot/model"
	"github.com/bismastr/discord-bot/utils"
	"github.com/bwmarrin/discordgo"
)

func ListSession(s *discordgo.Session, i *discordgo.InteractionCreate, listSessions []model.GamingSession) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		panic(err)
	}

	for j, session := range listSessions {
		fmt.Println(j)
		messageContent := singleSession(session)

		_, err := s.FollowupMessageCreate(i.Interaction, true, &messageContent)
		if err != nil {
			panic(err)
		}
	}
}

func singleSession(session model.GamingSession) discordgo.WebhookParams {
	messageContent := fmt.Sprintf("Session by <@%v>\n\nGame: <Feature on Development>\nStart: <Feature on Development>\nMembers: <%v>", session.CreatedBy.Id, utils.GenerateMemberMention(session.MembersSession))

	i := discordgo.WebhookParams{
		Content: messageContent,
	}

	return i
}
