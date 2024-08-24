package gaming_session

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	membersSession []string
)

func CreateGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "# Info Info Info Mabar dulu ga sih? @here",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Emoji: &discordgo.ComponentEmoji{
								Name: "🔥",
							},
							Label:    "Gas Join!",
							Style:    discordgo.PrimaryButton,
							CustomID: "mabar_yes",
						},
						discordgo.Button{
							Emoji: &discordgo.ComponentEmoji{
								Name: "❌",
							},
							Label:    "Skip duls",
							Style:    discordgo.SecondaryButton,
							CustomID: "mabar_no",
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

func JoinGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if CheckJoin(i.Member.User.ID) {
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

	userId := i.Member.User.ID
	messageContent := fmt.Sprintf("<@%v>Join abang quh 🥳\n\nArek-arek sing join 👥:%v  ", userId, GenerateMemberMention())
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
	alrJoin := fmt.Sprintf("Hei <@%v>, kalo udah join ga boleh sekip bang :( Wajib Ikut", userid, GenerateMemberMention())
	noJoin := fmt.Sprintf("<@%v> tidak join duls, kecewaaaa sangat berat! Join sini lahh :(", userid, GenerateMemberMention())

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

func GenerateMemberMention() string {
	result := ""
	for _, s := range membersSession {
		result += fmt.Sprintf("<@%v>", s)
	}
	return result
}

func CheckJoin(userId string) bool {
	for _, u := range membersSession {
		if u == userId {
			return true
		}
	}
	return false
}
