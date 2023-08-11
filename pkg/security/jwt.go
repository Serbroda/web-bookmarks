package security

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Serbroda/ragbag/pkg/dtos"
	"github.com/golang-jwt/jwt/v5"
)

var (
	JwtSecretKey       = "JWT_SECRET_KEY"
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
}

func GenerateJwtPair(user *dtos.User) (TokenPair, error) {
	userIdStr := strconv.FormatInt(user.ID, 10)

	accessToken, err := GenerateJwt(jwt.MapClaims{
		"sub":   userIdStr,
		"exp":   time.Now().Add(time.Minute * time.Duration(jwtAccessTokenExp)).Unix(),
		"iat":   time.Now().Unix(),
		"name":  user.Username,
		"roles": user.RolesAsStrings(),
	})
	if err != nil {
		return TokenPair{}, err
	}

	refreshToken, err := GenerateJwt(jwt.MapClaims{
		"sub": userIdStr,
		"exp": time.Now().Add(time.Minute * time.Duration(jwtRefreshTokenExp)).Unix(),
		"iat": time.Now().Unix(),
	})
	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{
		AccessToken:  &accessToken,
		RefreshToken: &refreshToken,
	}, nil
}

func GenerateJwt(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ParseJwt(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JwtSecretKey), nil
	})
	return t, err
}
