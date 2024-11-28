package user

import (
	"context"

	"github.com/bismastr/discord-bot/internal/repository"
)

type UserService struct {
	repository *repository.Queries
}

func NewUserService(repository *repository.Queries) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (u *UserService) Createuser(ctx context.Context, user repository.InsertUserParams) error {
	err := u.repository.InsertUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) GetUserByDiscordUID(ctx context.Context, dicord_uid int64) (*repository.User, error) {
	result, err := u.repository.GetUserByDiscordUID(ctx, dicord_uid)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
