package service

import (
	"context"
	"encoding/base32"
	"github.com/radityarestan/ecom-authentication/internal/entity"
	"github.com/radityarestan/ecom-authentication/internal/repository"
	"github.com/radityarestan/ecom-authentication/internal/shared"
	"github.com/radityarestan/ecom-authentication/internal/shared/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	FrontEndURLVerify = "https://ceritanya-front-end.com/verify/"
)

type (
	Auth interface {
		SignIn(ctx context.Context, req *dto.SignInRequest) (*dto.SignInResponse, error)
		SignUp(ctx context.Context, req *dto.SignUpRequest) (*dto.SignUpResponse, error)
		VerifyEmail(ctx context.Context, code string) (*dto.SignInResponse, error)
	}

	authService struct {
		deps shared.Deps
		repo repository.Auth
	}
)

func (a *authService) SignIn(ctx context.Context, req *dto.SignInRequest) (*dto.SignInResponse, error) {
	user, err := a.repo.FindUserByEmail(a.deps.Database.WithContext(ctx), req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			a.deps.Logger.Errorf("User not found: %v", err)
			return nil, dto.ErrFindUserNotFound
		}

		a.deps.Logger.Errorf("Error when find user by email: %v", err)
		return nil, dto.ErrFindUserFailed
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		a.deps.Logger.Errorf("Error when compare password: %v", err)
		return nil, dto.ErrPasswordNotMatch
	}

	if !user.IsVerified {
		a.deps.Logger.Errorf("User not verified")
		return nil, dto.ErrUserNotVerified
	}

	token, err := a.generateToken(user.ID)
	if err != nil {
		a.deps.Logger.Errorf("Error when generate token: %v", err)
		return nil, dto.ErrGenerateTokenFailed
	}

	return &dto.SignInResponse{
		Token: token,
	}, nil

}

func (a *authService) SignUp(ctx context.Context, req *dto.SignUpRequest) (*dto.SignUpResponse, error) {
	var orm = a.deps.Database.WithContext(ctx)

	oldUser, err := a.repo.FindUserByEmail(orm, req.Email)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			a.deps.Logger.Errorf("Error when find user by email: %v", err)
			return nil, dto.ErrFindUserFailed
		}
	}

	if oldUser != nil {
		a.deps.Logger.Errorf("User already exists")
		return nil, dto.ErrUserAlreadyExists
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	hashedPasswordString := string(hashedPassword)
	hashedEmailString := base32.StdEncoding.EncodeToString([]byte(req.Email))

	var user = &entity.User{
		Name:             req.Name,
		Email:            req.Email,
		Password:         hashedPasswordString,
		Photo:            "default.jpg",
		VerificationCode: hashedEmailString,
		IsVerified:       false,
	}

	userCreated, err := a.repo.CreateUser(orm, user)
	if err != nil {
		a.deps.Logger.Errorf("Failed to create user: %v", err)
		return nil, dto.ErrCreateUserFailed
	}

	go func() {
		err := a.deps.NSQProducer.Publish([]byte(userCreated.VerificationCode))
		if err != nil {
			a.deps.Logger.Errorf("Failed to publish to NSQ: %v", err)
			return
		}
	}()

	return &dto.SignUpResponse{
		ID:           userCreated.ID,
		Email:        userCreated.Email,
		Verification: FrontEndURLVerify + userCreated.VerificationCode,
	}, nil

}

func (a *authService) VerifyEmail(ctx context.Context, code string) (*dto.SignInResponse, error) {
	var orm = a.deps.Database.WithContext(ctx)

	user, err := a.repo.UpdateUserVerified(orm, code)
	if err != nil {
		a.deps.Logger.Errorf("Failed to update user verified: %v", err)
		return nil, dto.ErrUpdateUserFailed
	}

	token, err := a.generateToken(user.ID)
	if err != nil {
		a.deps.Logger.Errorf("Error when generate token: %v", err)
		return nil, dto.ErrGenerateTokenFailed
	}

	return &dto.SignInResponse{
		Token: token,
	}, nil
}

func NewAuth(deps shared.Deps, repo repository.Auth) (Auth, error) {
	return &authService{deps, repo}, nil
}
