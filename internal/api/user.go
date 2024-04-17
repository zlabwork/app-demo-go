package api

import (
	"app/internal/help"
	"app/internal/msg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Me(c echo.Context) error {
	token := help.ParseToken(c)

	return c.JSON(http.StatusOK, msg.DataWrap{
		Status:  msg.StatusSuccess,
		Message: "success",
		Data:    token.UserId,
	})
}
