package boot

import (
	"app/internal/consts"
	"app/internal/entity"
	"github.com/bwmarrin/snowflake"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strconv"
)

var (
	Dir  = &directory{}
	Env  = &environment{}
	Libs *libraries
)

type environment struct {
	IsDev  bool
	IsProd bool
}

type directory struct {
	Root   string
	Config string
	Data   string
}

type libraries struct {
	Snow *snowflake.Node
}

func (d *directory) Setting() {
	// TODO :: FIXME
	root := "./"
	d.Root = root
	d.Config = root + "config/"
}

func (e *environment) Setting() {
	if os.Getenv("APP_ENV") == "dev" {
		e.IsDev = true
	} else if os.Getenv("APP_ENV") == "prod" {
		e.IsProd = true
	}
}

func init() {

	// 1. setting
	Dir.Setting()
	Env.Setting()

	// 2. check env
	if len(os.Getenv("APP_ENV")) == 0 {
		log.Fatal("env is missing, try to execute 'export `cat .env`' before run the app")
	}

	// 3. config
	bs, err := os.ReadFile(Dir.Config + "app.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if yaml.Unmarshal(bs, entity.Yaml) != nil {
		log.Fatal(err)
	}

	// 4. libs
	Libs = NewLibs()
}

func NewLibs() *libraries {

	i, _ := strconv.ParseInt(os.Getenv("APP_NODE"), 10, 64)
	snowflake.Epoch = 1498612200000 // 2017-06-28 09:10:00
	snowflake.NodeBits = 8
	snowflake.StepBits = 14
	sn, _ := snowflake.NewNode(i)
	return &libraries{
		Snow: sn,
	}
}

func ParseToken(c echo.Context) *entity.TokenClaims {
	t := c.Get(consts.ContextKey).(*jwt.Token)
	return t.Claims.(*entity.TokenClaims)
}
