package dto

import "errors"

const (
	CreateUserSuccess  = "user created successfully"
	VerifyEmailSuccess = "email verified successfully"
	SignInSuccess      = "sign in successfully"
)

var (
	ErrFindUserFailed      = errors.New("error when find user by email")
	ErrFindUserNotFound    = errors.New("user not found")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrCreateUserFailed    = errors.New("failed to create user")
	ErrUpdateUserFailed    = errors.New("failed to update user")
	ErrInvalidCode         = errors.New("invalid verification code")
	ErrPasswordNotMatch    = errors.New("password not match")
	ErrUserNotVerified     = errors.New("user not verified")
	ErrGenerateTokenFailed = errors.New("failed to generate jwt token")
)

type (
	SignUpRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,password"`
	}

	SignUpResponse struct {
		ID           int64  `json:"id" validate:"required"`
		Email        string `json:"email" validate:"required,email"`
		Verification string `json:"verification" validate:"required"`
	}

	SignInRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,password"`
	}

	SignInResponse struct {
		Token string `json:"token" validate:"required"`
	}
)
