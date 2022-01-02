package auth

import (
	authEntities "auth-service/internal/core/entities/auth"
	authPort "auth-service/internal/core/ports/auth"
)

type AuthenticateService interface {
	authPort.Authenticate
}

type authenticateService struct {
	repo authPort.AuthenticationRepository
}

func NewAuthenticateService(repo authPort.AuthenticationRepository) AuthenticateService {
	return authenticateService{
		repo: repo,
	}
}

func (s authenticateService) Login(username, password string) (*authEntities.TokenPair, error) {
	user, err := s.repo.Exists(username, password)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
