package api

import (
	"app/internal/boot"
	"app/internal/entity"
	"app/internal/msg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Me(c echo.Context) error {
	token := boot.ParseToken(c)

	return c.JSON(http.StatusOK, entity.RespData{
		Code:    msg.OK,
		Message: "success",
		Data:    token.UserId,
	})
}
