package gaming_session

import (
	"context"
	"fmt"
	"strings"

	"github.com/bismastr/discord-bot/components"
	"github.com/bismastr/discord-bot/db"
	"github.com/bismastr/discord-bot/utils"
	"github.com/bwmarrin/discordgo"
)

func JoinGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient, ctx context.Context) {
	userid := i.Member.User.ID
	customId := i.MessageComponentData().CustomID

	split := strings.Split(customId, "_")
	refId := split[2]

	result, err := db.AddMemberToSession(ctx, refId, userid)
	if err != nil {
		panic(err)
	}

	if result == nil {
		components.AlreadyInSession(s, i)
	} else {
		components.JoinSession(s, i, userid, utils.GenerateMemberMention(result.MembersSession))
	}
}

func DeclineGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient, ctx context.Context) {
	userid := i.Member.User.ID
	noJoin := fmt.Sprintf("<@%v> tidak join duls, kecewaaaa sangat berat!", userid)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: noJoin,
		},
	})
	if err != nil {
		panic(err)
	}
}
