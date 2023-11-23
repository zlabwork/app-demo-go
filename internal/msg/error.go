package msg

const (
	OK          = 200
	Err         = 400
	ErrNotFound = 404
)

const (
	ErrDefault = iota + 20000
	ErrTimeout
	ErrSignature
	ErrAccess
	ErrEncoding
	ErrHeader
	ErrParameter
	ErrProcess
	ErrWrite
)

var statusText = map[int]string{
	OK:           "success",
	Err:          "error",
	ErrNotFound:  "page not found",
	ErrTimeout:   "error request time",
	ErrSignature: "error request signature",
	ErrAccess:    "error access",
	ErrEncoding:  "error encoding",
	ErrHeader:    "error request headers",
	ErrParameter: "error parameter",
	ErrProcess:   "error in process",
	ErrWrite:     "failed when write data",
}

func Text(code int) string {
	return statusText[code]
}
