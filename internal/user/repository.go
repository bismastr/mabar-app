package user

import (
	"github.com/bismastr/discord-bot/internal/db"
	"github.com/bismastr/discord-bot/internal/model"
	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct {
	db *db.Db
}

type UserRepository interface {
	GetUserByDiscordId(id int) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
}

func NewUserRepositoryImpl(db *db.Db) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) GetUserByDiscordId(id int) (*model.User, error) {
	var user model.User
	result := u.db.Client.First(&user, model.User{DiscordUID: id})

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u *UserRepositoryImpl) CreateUser(user *model.User) (*model.User, error) {
	result := u.db.Client.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "discord_uid"}},
		DoNothing: true,
	}).Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
