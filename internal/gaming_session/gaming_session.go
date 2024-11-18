package gaming_session

import "github.com/bismastr/discord-bot/internal/model"

type GamingSessionService struct {
	repository GamingSessionRepository
}

func NewGamingService(repository GamingSessionRepository) *GamingSessionService {
	return &GamingSessionService{
		repository: repository,
	}
}

func (g *GamingSessionService) CreateGaming(session *model.Session) (*model.Session, error) {
	result, err := g.repository.CreateGaming(session)
	if err != nil {
		return nil, err
	}

	return result, err
}
