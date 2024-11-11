package gamingSession

import "context"

type GamingSessionService struct {
	repository FirestoreRepositorySession
}

func NewGamingSessionService(repo FirestoreRepositorySession) *GamingSessionService {
	return &GamingSessionService{
		repository: repo,
	}
}

func (s *GamingSessionService) CreateGamingSession(ctx context.Context, model GamingSession) (string, error) {
	refId, err := s.repository.CreateGamingSession(ctx, model)
	if err != nil {
		return "", err
	}

	return refId, nil
}

func (s *GamingSessionService) UpdateGamingSessionByRefId(ctx context.Context, refId string, model GamingSession) error {
	err := s.repository.UpdateGamingSessionByRefId(ctx, refId, model)
	if err != nil {
		return err
	}

	return err
}

func (s *GamingSessionService) GetGamingSessionByRefId(ctx context.Context, refId string) (*GamingSession, error) {
	result, err := s.repository.GetGamingSessionByRefId(ctx, refId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
