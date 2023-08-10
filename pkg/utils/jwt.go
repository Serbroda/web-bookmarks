package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Serbroda/ragbag/pkg/user"
	"github.com/golang-jwt/jwt"
)

var (
	jwtSecretKey       = "JWT_SECRET_KEY"
	jwtAccessTokenExp  = 15
	jwtRefreshTokenExp = 10080
)

type TokenPair struct {
	AccessToken  *string `json:"accessToken"`
	RefreshToken *string `json:"refreshToken"`
}

type JwtCustomClaims struct {
	Name  string `json:"name,omitempty"`
	Roles string `json:"roles,omitempty"`
	jwt.StandardClaims
}

func GenerateTokenPair(user *user.User) (TokenPair, error) {
	userIdStr := strconv.FormatInt(user.ID, 10)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaims{
		Name: user.Username,
		StandardClaims: jwt.StandardClaims{
			Subject:   userIdStr,
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(jwtAccessTokenExp)).Unix(),
		},
	})
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return TokenPair{}, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userIdStr,
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(jwtRefreshTokenExp)).Unix(),
		},
	})
	rt, err := refreshToken.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{
		AccessToken:  &t,
		RefreshToken: &rt,
	}, nil
}

func ParseToken(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})
	return t, err
}
