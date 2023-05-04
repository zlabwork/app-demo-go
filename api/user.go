package api

import (
	"app"
	"app/entity"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Me(c echo.Context) error {
	user := c.Get(app.ContextKey).(*jwt.Token)
	claims := user.Claims.(*entity.TokenClaims)
	return c.String(http.StatusOK, "Welcome "+claims.UserId+"!")
}
