package api

import (
	"app/internal/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RespFailed(c echo.Context, code int) error {
	return c.JSON(http.StatusOK, entity.RespData{
		Code:    code,
		Message: http.StatusText(code),
	})
}

func RespFailedWithMessage(c echo.Context, code int, message string) error {
	return c.JSON(http.StatusOK, entity.RespData{
		Code:    code,
		Message: message,
	})
}

func RespSuccess(c echo.Context) error {
	return c.JSON(http.StatusOK, entity.RespData{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	})
}

func RespData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, entity.RespData{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	})
}
