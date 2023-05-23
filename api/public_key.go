package api

import (
	"app"
	"encoding/base64"
	"encoding/pem"
	"github.com/labstack/echo/v4"
	"io"
	"os"
)

func PublicKey(c echo.Context) error {

	// 1.
	env := os.Getenv("APP_ENV")
	f, err := os.Open("config/public." + env + ".pem")
	if err != nil {
		return err
	}
	bs, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	// 2.
	block, _ := pem.Decode(bs)
	key := base64.StdEncoding.EncodeToString(block.Bytes)

	type publicKey struct {
		Format    string `json:"format,omitempty"`
		PublicKey string `json:"public_key,omitempty"`
	}
	return app.RespData(c, &publicKey{
		Format:    "pkcs1",
		PublicKey: key,
	})
}
