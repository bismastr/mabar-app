package gaming_session

import (
	"context"
	"fmt"
	"strings"

	"github.com/bismastr/discord-bot/components"
	"github.com/bismastr/discord-bot/db"
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

	components.JoinSession(s, i, userid, GenerateMemberMention(result.MembersSession))
}

func DeclineGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient) {
	userid := i.Member.User.ID
	alrJoin := fmt.Sprintf("Hei <@%v>, kalo udah join ga boleh sekip bang :( Wajib Ikut", userid)
	noJoin := fmt.Sprintf("<@%v> tidak join duls, kecewaaaa sangat berat! Join sini lahh :(", userid)

	if CheckJoin(i.Member.User.ID) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: alrJoin,
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			panic(err)
		}

		return
	}
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
