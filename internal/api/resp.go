package api

import (
	"app/internal/msg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RespFailed(c echo.Context, status string) error {
	return c.JSON(http.StatusOK, msg.DataWrap{
		Status:  status,
		Message: status,
	})
}

func RespFailedWithMessage(c echo.Context, status string, message string) error {
	return c.JSON(http.StatusOK, msg.DataWrap{
		Status:  status,
		Message: message,
	})
}

func RespSuccess(c echo.Context) error {
	return c.JSON(http.StatusOK, msg.DataWrap{
		Status:  msg.StatusSuccess,
		Message: msg.StatusSuccess,
	})
}

func RespData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, msg.DataWrap{
		Status:  msg.StatusSuccess,
		Message: msg.StatusSuccess,
		Data:    data,
	})
}
