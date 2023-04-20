package api

import (
	"app/entity"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

func Token(c echo.Context) error {

	userId := "123456"
	userName := "XieShan"
	claims := &entity.TokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("APP_NAME"),
			Subject:   "jwt",
			Audience:  []string{userId, userName},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			NotBefore: jwt.NewNumericDate(time.Now().Add(-5 * time.Minute)),
			ID:        uuid.New().String(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("APP_KEY")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func Restricted(c echo.Context) error {
	user := c.Get("_token").(*jwt.Token)
	claims := user.Claims.(*entity.TokenClaims)
	return c.String(http.StatusOK, "Welcome "+claims.UserId+"!")
}
