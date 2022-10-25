package repository

import (
	"github.com/radityarestan/ecom-authentication/internal/entity"
	"gorm.io/gorm"
)

type (
	Auth interface {
		FindUserByEmail(orm *gorm.DB, email string) (*entity.User, error)
		CreateUser(orm *gorm.DB, user *entity.User) (*entity.User, error)
		UpdateUserVerified(orm *gorm.DB, code string) error
	}

	authRepo struct{}
)

func (a *authRepo) FindUserByEmail(orm *gorm.DB, email string) (*entity.User, error) {
	var user = &entity.User{}
	if err := orm.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (a *authRepo) CreateUser(orm *gorm.DB, user *entity.User) (*entity.User, error) {
	if err := orm.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (a *authRepo) UpdateUserVerified(orm *gorm.DB, code string) error {
	if err := orm.Model(&entity.User{}).Where("verification_code = ?", code).Update("is_verified", true).Error; err != nil {
		return err
	}

	return nil
}

func NewAuth() (Auth, error) {
	return &authRepo{}, nil
}
