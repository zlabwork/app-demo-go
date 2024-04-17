package help

import (
	"app/internal/consts"
	"app/internal/entity"
	"github.com/bwmarrin/snowflake"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"os"
	"strconv"
)

func ParseToken(c echo.Context) *entity.TokenClaims {
	t := c.Get(consts.ContextKey).(*jwt.Token)
	return t.Claims.(*entity.TokenClaims)
}

var (
	Libs *libraries
)

type libraries struct {
	Snow *snowflake.Node
}

func newLibs() *libraries {

	i, _ := strconv.ParseInt(os.Getenv("APP_NODE"), 10, 64)
	snowflake.Epoch = 1498612200000 // 2017-06-28 09:10:00
	snowflake.NodeBits = 8
	snowflake.StepBits = 14
	sn, _ := snowflake.NewNode(i)
	return &libraries{
		Snow: sn,
	}
}

func init() {
	// 4. libs
	Libs = newLibs()
}
