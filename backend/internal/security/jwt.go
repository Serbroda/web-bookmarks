package security

import (
	"backend/internal/db/sqlc"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Jwt = string

var (
	JwtAccessTokenSecret  = "JWT_ACCESS_TOKEN_SECRET_KEY"
	JwtRefreshTokenSecret = "JWT_REFRESH_TOKEN_SECRET_KEY"
	jwtAccessTokenExp     = 15
	jwtRefreshTokenExp    = 10080
)

type TokenPair struct {
	AccessToken            *Jwt      `json:"accessToken"`
	AccessTokenExpiration  time.Time `json:"-"`
	RefreshToken           *Jwt      `json:"-"`
	RefreshTokenExpiration time.Time `json:"-"`
}

type JwtCustomClaims struct {
	Name  string `json:"name,omitempty"`
	Roles string `json:"roles,omitempty"`
}

func GenerateJwtPair(user sqlc.User) (TokenPair, error) {
	accessTokenExpiration := time.Now().Add(time.Minute * time.Duration(jwtAccessTokenExp))
	accessToken, err := GenerateJwt(jwt.MapClaims{
		"sub":  user.ID,
		"exp":  accessTokenExpiration.Unix(),
		"iat":  time.Now().Unix(),
		"name": user.Username,
		//"roles": user.RolesAsStrings(),
	}, JwtAccessTokenSecret)
	if err != nil {
		return TokenPair{}, err
	}

	refreshTokenExpiration := time.Now().Add(time.Minute * time.Duration(jwtRefreshTokenExp))
	refreshToken, err := GenerateJwt(jwt.MapClaims{
		"sub": user.ID,
		"exp": refreshTokenExpiration.Unix(),
		"iat": time.Now().Unix(),
	}, JwtRefreshTokenSecret)
	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{
		AccessToken:            &accessToken,
		AccessTokenExpiration:  accessTokenExpiration,
		RefreshToken:           &refreshToken,
		RefreshTokenExpiration: refreshTokenExpiration,
	}, nil
}

func GenerateJwt(claims jwt.MapClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func VerifyAccessToken(token string) (*jwt.Token, error) {
	return parseToken(token, JwtAccessTokenSecret)
}

func VerifyRefreshToken(token string) (*jwt.Token, error) {
	return parseToken(token, JwtRefreshTokenSecret)
}

func parseToken(token string, secret string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	return t, err
}
