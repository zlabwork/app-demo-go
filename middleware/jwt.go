package middleware

import (
	"app"
	"app/entity"
	"app/msg"
	"github.com/golang-jwt/jwt/v5"
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
		ContextKey: app.ContextKey,
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(msg.OK, entity.RespData{
				Code:    msg.ErrAccess,
				Message: "error access",
			})
		},
	}

	return echojwt.WithConfig(config)

}
