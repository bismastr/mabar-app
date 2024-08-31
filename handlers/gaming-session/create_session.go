package gaming_session

import (
	"context"
	"time"

	"github.com/bismastr/discord-bot/components"
	"github.com/bismastr/discord-bot/db"
	"github.com/bismastr/discord-bot/model"
	"github.com/bwmarrin/discordgo"
)

func CreateSession(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient, ctx context.Context) {
	if IsHaveSession(s, i, db, ctx) {
		components.UnableCreateSession(s, i)
		return
	}

	session := model.GamingSession{
		CreatedAt: time.Now().String(),
		CreatedBy: &model.CreatedBy{
			Id:       i.Member.User.ID,
			Username: i.Member.Nick,
		},
		SessionEnd:   "", //Need to add session
		SessionStart: "",
		IsFinish:     false,
	}

	id, err := db.CreateGamingSession(ctx, session)
	if err != nil {
		panic(err)
	}

	// Getting option (slash-command parameter) value
	gameName := ""
	if optionValue := GetOptionValueByName(i, "nama-permainan"); optionValue != nil {
		gameName = optionValue.(string)
	}

	components.CreateSession(s, i, id, gameName)
}

func IsHaveSession(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient, ctx context.Context) bool {
	session, err := db.ReadGamingSessionByCreatedUserId(ctx, i.Member.User.ID)
	if err != nil {
		panic(err)
	}

	//If no session, then user doesnt have session active
	if session == nil {
		return false
	}
	//Need to check is session finished or not. If still have session active then user cannot create mabar
	if !session.IsFinish {
		return true
	} else {
		return false
	}
}
