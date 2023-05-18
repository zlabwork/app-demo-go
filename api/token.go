package api

import (
	"app"
	"app/entity"
	"app/msg"
	"app/service/cache"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

const (
	accessTokenTimeout  = time.Hour * 72
	refreshTokenTimeout = time.Hour * 24 * 20
	refreshTokenKey     = "_RT:%d"
)

func Token(c echo.Context) error {

	token, err := genToken(123456)
	if err != nil {
		return app.RespFailed(c, msg.ErrProcess)
	}

	return app.RespData(c, token)
}

func RefreshToken(c echo.Context) error {

	type handler struct {
		UserId       int64  `json:"user_id,omitempty"`
		RefreshToken string `json:"refresh_token,omitempty"`
	}
	req := handler{}
	if json.NewDecoder(c.Request().Body).Decode(&req) != nil {
		return app.RespFailed(c, msg.ErrParameter)
	}

	// get refresh_token
	key := fmt.Sprintf(refreshTokenKey, req.UserId)
	bs, err := cache.Get(c.Request().Context(), key)
	if err != nil {
		return app.RespFailed(c, msg.ErrProcess)
	}

	// check
	if string(bs) != req.RefreshToken {
		return app.RespFailedWithMessage(c, msg.ErrAccess, "error refresh_token")
	}

	// new access_token
	token, err := genToken(req.UserId)
	if err != nil {
		return app.RespFailed(c, msg.ErrProcess)
	}

	return app.RespData(c, token)
}

func genToken(userId int64) (*entity.Token, error) {

	userName := ""
	expiresAt := time.Now().Add(accessTokenTimeout)
	claims := &entity.TokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("APP_NAME"),
			Subject:   "jwt",
			Audience:  []string{userName},
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(time.Now().Add(-5 * time.Minute)),
			ID:        uuid.New().String(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	tk, err := token.SignedString([]byte(os.Getenv("APP_KEY")))
	if err != nil {
		return nil, err
	}

	// Refresh Token
	rt := uuid.New().String()
	key := fmt.Sprintf(refreshTokenKey, userId)
	cache.Set(context.TODO(), key, []byte(rt), refreshTokenTimeout)

	return &entity.Token{
		UserId:       userId,
		AccessToken:  tk,
		RefreshToken: rt,
		ExpiresAt:    expiresAt,
	}, nil
}
