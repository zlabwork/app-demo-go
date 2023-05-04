package api

import (
	"app"
	"app/entity"
	"app/msg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Me(c echo.Context) error {
	user := app.GetUser(c)

	return c.JSON(http.StatusOK, entity.RespData{
		Code:    msg.OK,
		Message: "success",
		Data:    user.UserId,
	})
}
