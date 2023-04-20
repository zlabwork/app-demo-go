package middleware

import (
	"app/entity"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"os"
)

func Jwt() echo.MiddlewareFunc {

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entity.TokenClaims)
		},
		SigningKey: []byte(os.Getenv("APP_KEY")),
		ContextKey: "_token",
	}

	return echojwt.WithConfig(config)

}
