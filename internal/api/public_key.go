package api

import (
	"app/internal/msg"
	"encoding/base64"
	"encoding/pem"
	"github.com/labstack/echo/v4"
	"io"
	"os"
)

var _publicKey []byte

func PublicKey(c echo.Context) error {

	if _publicKey == nil {

		// 1.
		env := os.Getenv("APP_ENV")
		f, err := os.Open("config/public." + env + ".pem")
		if err != nil {
			return RespFailed(c, msg.StatusNotFound)
		}
		bs, err := io.ReadAll(f)
		if err != nil {
			return RespFailed(c, msg.StatusServerError)
		}

		// 2.
		block, _ := pem.Decode(bs)
		_publicKey = block.Bytes
	}

	key := base64.StdEncoding.EncodeToString(_publicKey)
	type publicKey struct {
		Format    string `json:"format,omitempty"`
		PublicKey string `json:"public_key,omitempty"`
	}
	return RespData(c, &publicKey{
		Format:    "pkcs1",
		PublicKey: key,
	})
}
