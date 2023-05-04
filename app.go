package app

import (
	"app/entity"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

const (
	ContextKey = "_token"
)

func GetUser(c echo.Context) *entity.TokenClaims {
	user := c.Get(ContextKey).(*jwt.Token)
	return user.Claims.(*entity.TokenClaims)
}
