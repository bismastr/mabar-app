package user

import (
	"strconv"

	"github.com/bismastr/discord-bot/internal/model"
	"github.com/markbates/goth"
)

type UserService struct {
	userRepository *UserRepositoryImpl
}

func NewUserService(userRepository *UserRepositoryImpl) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Createuser(user *goth.User) (*model.User, error) {
	discordId, _ := strconv.Atoi(user.UserID)
	userConverted := &model.User{
		DiscordUID: discordId,
		Username:   user.Name,
		AvatarURL:  user.AvatarURL,
	}

	result, err := u.userRepository.CreateUser(userConverted)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UserService) GetUserByDiscordId(discordId int) (*model.User, error) {
	result, err := u.userRepository.GetUserByDiscordId(discordId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
