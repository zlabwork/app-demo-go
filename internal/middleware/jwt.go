package middleware

import (
	"app/internal/consts"
	"app/internal/entity"
	"app/internal/msg"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func Jwt() echo.MiddlewareFunc {

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entity.TokenClaims)
		},
		SigningKey: []byte(os.Getenv("APP_KEY")),
		ContextKey: consts.ContextKey,
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusOK, msg.DataWrap{
				Status:  msg.StatusUnauthorized,
				Message: "error access",
			})
		},
		SuccessHandler: func(c echo.Context) {
			t := c.Get(consts.ContextKey).(*jwt.Token)
			uid := t.Claims.(*entity.TokenClaims).UserId
			c.Set(consts.UserId, uid)
		},
	}

	return echojwt.WithConfig(config)

}
