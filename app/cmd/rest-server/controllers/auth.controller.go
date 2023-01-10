package controllers

import (
	utils2 "github.com/Serbroda/ragbag/app/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"

	. "github.com/Serbroda/ragbag/app/pkg/models"
	. "github.com/Serbroda/ragbag/app/pkg/services"
)

var (
	jwtSecretKey       = utils2.MustGetEnv("JWT_SECRET_KEY")
	jwtAccessTokenExp  = utils2.MustParseInt64(utils2.GetEnvFallback("JWT_ACCESS_EXPIRE_MINUTES", "15"))
	jwtRefreshTokenExp = utils2.MustParseInt64(utils2.GetEnvFallback("JWT_REFRESH_EXPIRE_MINUTES", "10080"))
	baseUrl            = utils2.MustGetEnv("SERVER_BASE_URL")
)

type AuthController struct {
	UserService UserService
}

type Jwt = string

type JwtCustomClaims struct {
	Name  string `json:"name,omitempty"`
	Roles string `json:"roles,omitempty"`
	jwt.StandardClaims
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  *Jwt `json:"access_token,omitempty"`
	RefreshToken *Jwt `json:"refresh_token,omitempty"`
}

type RegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	Token Jwt `json:"refresh_token"`
}

type ActivateParams struct {
	Code string `form:"code" json:"code"`
}

func RegisterAuthController(e *echo.Echo, c AuthController, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.POST(baseUrl+"/auth/login", c.Login, middlewares...)
	e.POST(baseUrl+"/auth/register", c.Register, middlewares...)
	e.POST(baseUrl+"/auth/refresh_token", c.RefreshToken, middlewares...)
}

// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "login body"
// @Success 200 {object} LoginResponse
// @Failure 400
// @Router /auth/login [post]
func (c *AuthController) Login(ctx echo.Context) error {
	var payload LoginRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.Username == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	user, err := c.UserService.FindOneByUsername(payload.Username)
	if err != nil || user.ID < 1 || !utils2.CheckBcryptHash(payload.Password, user.Password) {
		return ctx.String(http.StatusBadRequest, "invalid login")
	}

	if !user.Active {
		return ctx.String(http.StatusBadRequest, "user not active")
	}
	t, rt, err := generateTokenPair(&user)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate token")
	}

	return ctx.JSON(http.StatusOK, LoginResponse{
		AccessToken:  &t,
		RefreshToken: &rt,
	})
}

// @Tags auth
// @Accept json
// @Produce json
// @Param registration body RegistrationRequest true "registration body"
// @Success 200
// @Failure 400
// @Router /auth/register [post]
func (c *AuthController) Register(ctx echo.Context) error {
	var payload RegistrationRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.Email == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	_, err = c.UserService.FindOneByUsername(payload.Email)

	if err != nil {
		return ctx.String(http.StatusConflict, "user already exists")
	}

	hashedPassword, _ := utils2.HashBcrypt(payload.Password)

	_, err = c.UserService.CreateUser(User{
		Username: payload.Email,
		Password: hashedPassword,
		Email:    payload.Email,
		Active:   false,
	})
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	/*_, err = services.Service.CreateActivationToken(ctx.Request().Context(), u.ID)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = utils2.SendMailTemplate(utils2.MailWithTemplate{
		Mail: utils2.Mail{
			To:      []string{payload.Email},
			Subject: "Verify your email address",
		},
		Template: "resources/templates/email/email-verification.html",
		Data: struct {
			Name string
			Link string
		}{
			Name: payload.FirstName,
			Link: fmt.Sprintf("%s/auth/activate?code=%s", baseUrl, token),
		},
	})

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}*/

	return ctx.JSON(http.StatusCreated, struct {
	}{})
}

// @Tags auth
// @Accept json
// @Produce json
// @Param refreshtoken body RefreshTokenRequest true "refresh token body"
// @Success 200
// @Failure 400
// @Router /auth/refresh_token [post]
func (c *AuthController) RefreshToken(ctx echo.Context) error {
	return ctx.String(http.StatusInternalServerError, "not implemented")
}

// @Id activateUser
// @Tags auth
// @Param code query string true "code"
// @Success 200
// @Failure 400
// @Router /auth/activate [get]
func (c *AuthController) ActivateUser(ctx echo.Context) error {
	return ctx.String(http.StatusInternalServerError, "not implemented")
}

func generateTokenPair(user *User) (string, string, error) {
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
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userIdStr,
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(jwtRefreshTokenExp)).Unix(),
		},
	})
	rt, err := refreshToken.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", "", err
	}
	return t, rt, nil
}
