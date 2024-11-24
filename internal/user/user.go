package user

import (
	"context"
	"strconv"

	"github.com/bismastr/discord-bot/internal/repository"
	"github.com/markbates/goth"
)

type UserService struct {
	repository *repository.Queries
}

func NewUserService(repository *repository.Queries) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (u *UserService) Createuser(ctx context.Context, user *goth.User) error {
	discordId, _ := strconv.Atoi(user.UserID)
	userConverted := repository.InsertUserParams{
		Username:   user.Name,
		AvatarUrl:  user.AvatarURL,
		DiscordUid: int64(discordId),
	}

	err := u.repository.InsertUser(ctx, userConverted)
	if err != nil {
		return err
	}

	return nil
}
