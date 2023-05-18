package api

import (
	"app"
	"app/entity"
	"app/service/cache"
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
	refreshTokenKey     = "_RT:%s"
)

func Token(c echo.Context) error {

	userId := "123456"
	userName := "XieShan"
	expiresAt := time.Now().Add(accessTokenTimeout)
	claims := &entity.TokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("APP_NAME"),
			Subject:   "jwt",
			Audience:  []string{userId, userName},
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
		return err
	}

	// Refresh Token
	rt := uuid.New().String()
	key := fmt.Sprintf(refreshTokenKey, userId)
	cache.Set(c.Request().Context(), key, []byte(rt), refreshTokenTimeout)

	return app.RespData(c, entity.Token{
		UserId:       userId,
		AccessToken:  tk,
		RefreshToken: rt,
		ExpiresAt:    expiresAt,
	})
}

func RefreshToken(c echo.Context) error {

	return app.RespSuccess(c)
}
