package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Rahmans11/koda-b5-backend/internal/dto"
)

type AuthService struct {
	registeredUsers map[string]dto.AuthData
}

func NewAuthService() *AuthService {
	return &AuthService{
		registeredUsers: make(map[string]dto.AuthData),
	}
}

func (a *AuthService) InputFormatValidation(input *dto.AuthData) error {
	email := strings.TrimSpace(input.Email)
	if email == "" {
		return errors.New("email is required")
	}

	if !strings.Contains(email, "@") {
		return errors.New("invalid email format: must contain '@'")
	}

	if strings.HasPrefix(email, "@") || strings.HasSuffix(email, "@") {
		return errors.New("invalid email format: '@' cannot be at start or end")
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return errors.New("invalid email format: must have exactly one '@'")
	}

	localPart := parts[0]
	domain := parts[1]

	if localPart == "" {
		return errors.New("invalid email format: local part cannot be empty")
	}

	if !strings.Contains(domain, ".") || strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") {
		return errors.New("invalid email format: domain must contain '.' and not start/end with it")
	}

	if len(input.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, char := range input.Password {
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		} else if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func (a *AuthService) RegisterUser(input *dto.AuthData) error {
	if err := a.InputFormatValidation(input); err != nil {
		return err
	}

	if _, exists := a.registeredUsers[input.Email]; exists {
		return errors.New("email already registered")
	}

	a.registeredUsers[input.Email] = *input
	return nil
}

func (a *AuthService) LoginValidation(input *dto.AuthData) error {

	registeredUser, exists := a.registeredUsers[input.Email]
	if !exists {
		return errors.New("email not registered")
	}

	if input.Password != registeredUser.Password {
		return errors.New("invalid password")
	}

	return nil
}
