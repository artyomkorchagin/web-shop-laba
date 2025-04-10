package user

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	signingKey = "qrkj#%@FNSAzpZ!@M<24FjH" // i need to change it to secret
	tokenTTL   = 12 * time.Hour
)

func (s *Service) GenerateToken(ctx context.Context, email, password string) (string, error) {

	user, err := s.repo.GetUser(ctx, email)
	if err != nil {
		return "", err
	}

	if !VerifyPassword(password, user.PasswordHash) {
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   user.Email,
		Issuer:    "kursach-app",
		Audience:  user.Role,
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	})

	return token.SignedString([]byte(signingKey))
}
