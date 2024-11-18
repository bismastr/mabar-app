package gaming_session

import (
	"github.com/bismastr/discord-bot/internal/db"
	"github.com/bismastr/discord-bot/internal/model"
)

type GamingSessionRepositoryImpl struct {
	db *db.Db
}

type GamingSessionRepository interface {
	CreateGaming(req *model.Session) (*model.Session, error)
	GetAllGaming() (*[]model.Session, error)
}

func NewGamingSessionRepository(db *db.Db) *GamingSessionRepositoryImpl {
	return &GamingSessionRepositoryImpl{
		db: db,
	}
}

func (r *GamingSessionRepositoryImpl) CreateGaming(session *model.Session) (*model.Session, error) {
	result := r.db.Client.Create(session)

	if result.Error != nil {
		return nil, result.Error
	}

	return session, nil
}

func (r *GamingSessionRepositoryImpl) GetAllGaming() (*[]model.Session, error) {
	var result []model.Session
	err := r.db.Client.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
