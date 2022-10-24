package repository

import (
	"github.com/radityarestan/ecom-authentication/internal/entity"
	"gorm.io/gorm"
)

type (
	Auth interface {
		CreateUser(orm *gorm.DB, user *entity.User) (*entity.User, error)
	}

	authRepo struct{}
)

func (a *authRepo) CreateUser(orm *gorm.DB, user *entity.User) (*entity.User, error) {
	if err := orm.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func NewAuth() (Auth, error) {
	return &authRepo{}, nil
}
