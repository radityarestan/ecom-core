package service

import (
	"context"
	"github.com/radityarestan/ecom-authentication/internal/entity"
	"github.com/radityarestan/ecom-authentication/internal/repository"
	"github.com/radityarestan/ecom-authentication/internal/shared"
	"github.com/radityarestan/ecom-authentication/internal/shared/dto"
)

type (
	Auth interface {
		CreateUser(ctx context.Context, req *dto.TestRequest) (*dto.TestResponse, error)
	}

	authService struct {
		deps shared.Deps
		repo repository.Auth
	}
)

func (a *authService) CreateUser(ctx context.Context, req *dto.TestRequest) (*dto.TestResponse, error) {
	var orm = a.deps.Database.WithContext(ctx)
	userCreated, err := a.repo.CreateUser(orm, &entity.User{Name: req.Name})

	if err != nil {
		a.deps.Logger.Errorf("Failed to create user: %v", err)
		return nil, err
	}

	return &dto.TestResponse{
		ID:   userCreated.ID,
		Name: userCreated.Name,
	}, nil

}

func NewAuth(deps shared.Deps, repo repository.Auth) (Auth, error) {
	return &authService{deps, repo}, nil
}
