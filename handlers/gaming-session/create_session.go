package gaming_session

import (
	"context"
	"fmt"
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

	err := db.CreateGamingSession(ctx, session)
	if err != nil {
		panic(err)
	}

	components.CreateSession(s, i)
}

func IsHaveSession(s *discordgo.Session, i *discordgo.InteractionCreate, db *db.DbClient, ctx context.Context) bool {
	session, err := db.ReadGamingSessionByCreatedUserId(ctx, i.Member.User.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(session)

	if session == nil {
		return false
	}

	if !session.IsFinish {
		return true
	} else {
		return false
	}
}
