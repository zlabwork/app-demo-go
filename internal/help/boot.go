package help

import (
	"app/internal/entity"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var (
	Dir = &directory{}
	Env = &environment{}
)

type environment struct {
	IsLocal bool
	IsDev   bool
	IsProd  bool
}

type directory struct {
	Root   string
	Config string
	Data   string
}

func (d *directory) setting() {
	// TODO :: FIXME
	root := "./"
	d.Root = root
	d.Config = root + "config/"
	d.Data = root + "data/"
}

func (e *environment) setting() {
	if os.Getenv("APP_ENV") == "prod" {
		e.IsProd = true
	} else if os.Getenv("APP_ENV") == "dev" {
		e.IsDev = true
	} else if os.Getenv("APP_ENV") == "local" {
		e.IsLocal = true
	} else {
		e.IsProd = true
	}
}

func init() {

	// 1. setting
	Dir.setting()
	Env.setting()

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
}
