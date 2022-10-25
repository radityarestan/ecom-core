package repository

import (
	"github.com/radityarestan/ecom-authentication/internal/entity"
	"gorm.io/gorm"
)

type (
	Auth interface {
		FindUserByEmail(orm *gorm.DB, email string) (*entity.User, error)
		CreateUser(orm *gorm.DB, user *entity.User) (*entity.User, error)
		UpdateUserVerified(orm *gorm.DB, code string) (*entity.User, error)
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

func (a *authRepo) UpdateUserVerified(orm *gorm.DB, code string) (*entity.User, error) {
	var user = &entity.User{}
	if err := orm.Where("verification_code = ?", code).First(user).Error; err != nil {
		return nil, err
	}

	user.IsVerified = true
	if err := orm.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func NewAuth() (Auth, error) {
	return &authRepo{}, nil
}
