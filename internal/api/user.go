package api

import (
	"app/internal/boot"
	"app/internal/msg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Me(c echo.Context) error {
	token := boot.ParseToken(c)

	return c.JSON(http.StatusOK, msg.DataWrap{
		Status:  msg.StatusSuccess,
		Message: "success",
		Data:    token.UserId,
	})
}
