package app

import (
	"app/entity"
	"github.com/bwmarrin/snowflake"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"os"
	"strconv"
)

const (
	ContextKey = "_token"
)

var Libs *libs

type libs struct {
	Snow *snowflake.Node
}

func init() {
	Libs = NewLibs()
}

func NewLibs() *libs {

	i, _ := strconv.ParseInt(os.Getenv("APP_NODE"), 10, 64)
	snowflake.Epoch = 1498612200000 // 2017-06-28 09:10:00
	snowflake.NodeBits = 8
	snowflake.StepBits = 14
	sn, _ := snowflake.NewNode(i)
	return &libs{
		Snow: sn,
	}
}

func GetUser(c echo.Context) *entity.TokenClaims {
	user := c.Get(ContextKey).(*jwt.Token)
	return user.Claims.(*entity.TokenClaims)
}
