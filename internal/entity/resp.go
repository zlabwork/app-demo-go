package entity

type RespData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Refer   interface{} `json:"refer,omitempty"`
}
