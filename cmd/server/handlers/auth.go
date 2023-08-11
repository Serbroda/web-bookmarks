package handlers

import (
	"net/http"

	"github.com/Serbroda/ragbag/pkg/security"
	"github.com/Serbroda/ragbag/pkg/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//var baseUrl = "SERVER_BASE_URL"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type AuthHandler struct {
	UserService services.UserService
}

func RegisterAuthHandlers(e *echo.Echo, c AuthHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.POST(baseUrl+"/auth/login", c.Login, middlewares...)
	e.POST(baseUrl+"/auth/refresh_token", c.RefreshToken, middlewares...)
}

func (si *AuthHandler) Login(ctx echo.Context) error {
	var payload LoginRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.Username == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	user, err := si.UserService.FindByUsername(ctx.Request().Context(), payload.Username)
	if err != nil || user.ID < 1 || !user.Active || !security.CheckBcryptHash(payload.Password, user.Password) {
		return ctx.String(http.StatusBadRequest, "invalid login")
	}

	tokenPair, err := security.GenerateJwtPair(&user)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate token")
	}
	return ctx.JSON(http.StatusOK, tokenPair)
}

func (si *AuthHandler) RefreshToken(ctx echo.Context) error {
	var payload RefreshTokenRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.RefreshToken == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	token, err := security.ParseJwt(payload.RefreshToken)
	if err != nil || !token.Valid {
		return middleware.ErrJWTInvalid
	}
	auth, err := security.ParseToken(token)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	user, err := si.UserService.FindOne(ctx.Request().Context(), auth.Subject)
	if err != nil || user.ID < 1 || !user.Active {
		return echo.ErrUnauthorized
	}

	newTokenPair, err := security.GenerateJwtPair(&user)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, newTokenPair)
}

/*func (si *AuthHandler) Register(ctx echo.Context) error {
	var payload RegistrationRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.Username == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	if si.Services.ExistsUser(ctx.Request().Context(), payload.Username) {
		return ctx.String(http.StatusConflict, "user already exists")
	}

	hashedPassword, _ := utils.HashBcrypt(payload.Password)

	user, err := si.Services.CreateUser(ctx.Request().Context(), sqlc.CreateUserParams{
		Username: strings.ToLower(payload.Username),
		Password: hashedPassword,
		Email:    payload.Email,
		Active:   false,
	})
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	token, err := services.Service.CreateActivationToken(ctx.Request().Context(), user.ID)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = mail.SendMailTemplate(mail.MailWithTemplate{
		Mail: mail.Mail{
			To:      []string{payload.Email},
			Subject: "Verify your email address",
		},
		Template: "resources/templates/email/email-verification.html",
		Data: struct {
			Link string
		}{
			Link: fmt.Sprintf("%s/auth/activate?code=%s", baseUrl, token),
		},
	})

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, &user.User{
		Id:       &user.ID,
		Username: &user.Username,
	})
}

func (si *AuthHandler) Activate(ctx echo.Context, params public.ActivateParams) error {
	err := si.Services.ActivateUser(ctx.Request().Context(), params.Code)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return ctx.String(http.StatusOK, "user activated")
}

func (si *AuthHandler) RequestPasswordReset(ctx echo.Context) error {
	var payload public.RequestPasswordResetJSONRequestBody
	err := ctx.Bind(&payload)
	if err != nil || payload.Email == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	user, err := si.Queries.FindByEmail(ctx.Request().Context(), payload.Email)
	if err != nil {
		return ctx.String(http.StatusNoContent, "password reset mail sent")
	}

	token, err := si.Services.CreatePasswordResetToken(ctx.Request().Context(), user.ID)
	if err == nil {
		mail.SendMailTemplate(mail.MailWithTemplate{
			Mail: mail.Mail{
				To:      []string{payload.Email},
				Subject: "Password reset",
			},
			Template: "resources/templates/email/password-reset.html",
			Data: struct {
				Name     string
				Username string
				Link     string
			}{
				Name:     user.FirstName,
				Username: user.Username,
				Link:     fmt.Sprintf("%s/api/v1/password_reset?code=%s", baseUrl, token),
			},
		})
	}

	return ctx.String(http.StatusNoContent, "password reset mail sent")
}

func (si *AuthHandler) StartPasswordReset(ctx echo.Context, params public.StartPasswordResetParams) error {
	ctx.Response().Header().Set("password_reset_code", params.Code)
	return ctx.Redirect(http.StatusFound, fmt.Sprintf("%s/password_reset", baseUrl))
}

func (si *AuthHandler) ResetPassword(ctx echo.Context) error {
	var payload public.ResetPasswordJSONRequestBody
	err := ctx.Bind(&payload)
	if err != nil || payload.Email == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	prt, err := si.Queries.FindPasswordResetCodeByEmailAndToken(ctx.Request().Context(), sqlc.FindPasswordResetCodeByEmailAndTokenParams{
		Email:     payload.Email,
		TokenHash: utils.HashSha3256(payload.Code),
	})

	if err != nil {
		return err
	}

	if prt.ExpiresAt.Before(time.Now()) || !prt.UserActive {
		return ctx.String(http.StatusBadRequest, "password reset code expired or user not active")
	}

	err = si.Services.ChangePassword(ctx.Request().Context(), prt.UserID, payload.Password)
	if err != nil {
		return err
	}
	return ctx.String(http.StatusNoContent, "password reset")
}
*/
