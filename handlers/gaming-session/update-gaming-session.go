package gaming_session

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func JoinGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userid := i.Member.User.ID

	if CheckJoin(userid) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Kamu udah join bang :(",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			panic(err)
		}

		return
	}

	membersSession = append(membersSession, i.Member.User.ID)
	messageContent := fmt.Sprintf("<@%v> join abang quh 🥳\n\nArek-arek sing join 👥:%v  ", userid, GenerateMemberMention())
	go func() {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: messageContent,
			},
		})
		if err != nil {
			panic(err)
		}
	}()
}

func DeclineGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
