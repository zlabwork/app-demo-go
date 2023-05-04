package api

import (
	"app"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Me(c echo.Context) error {
	user := app.GetUser(c)
	return c.JSON(http.StatusOK, echo.Map{
		"user_id": user.UserId,
	})
}
