package gaming_session

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	membersSession []string
	mabarSession   bool
)

func CreateGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !mabarSession {
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
		} else {
			mabarSession = true
		}
	} else {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Sesi mabar sudah ada, kawanku",
			},
		})

		if err != nil {
			panic(err)
		}
	}

}

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
	messageContent := fmt.Sprintf("<@%v>Join abang quh 🥳\n\nArek-arek sing join 👥:%v  ", userid, GenerateMemberMention())
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

func DeleteGamingSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if mabarSession {
		messageContent := "Buyar dulu kawanku, sampai jumpa di mabar berikutnya!👋"
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: messageContent,
			},
		})

		mabarSession = false
		membersSession = nil

		if err != nil {
			panic(err)
		}

	} else {
		messageContent := "❌ Tidak ada sesi mabar yang berjalan, kawan"
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
