package usecase

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"lg/internal/entity"
	"time"
)

type JwtUseCase struct {
	u      UserContract
	secret string
}

func NewJwtUseCase(u UserContract, secret string) *JwtUseCase {
	return &JwtUseCase{
		u:      u,
		secret: secret,
	}
}

func (j *JwtUseCase) CompareUserPassword(ctx context.Context, us entity.User) error {
	userFromDb, err := j.u.GetUser(ctx, us.Email)
	if err != nil {
		return fmt.Errorf("cannot get user from db: %w", err)
	}
	return bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(us.Password))
}

func (j *JwtUseCase) GenerateToken(uuid string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(120 * time.Minute).Unix(),
		Subject:   uuid,
	})
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", fmt.Errorf("cannot generate jwt key %w", err)
	}
	return tokenString, nil
}

func (j *JwtUseCase) CheckToken(token string) (string, error) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse token: %v", err)
	}
	claims, ok := tkn.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token")
	}
	return claims["sub"].(string), nil
}
