package entity

type RespData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	TraceId interface{} `json:"trace_id,omitempty"`
}
