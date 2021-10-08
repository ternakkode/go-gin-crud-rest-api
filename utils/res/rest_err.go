package res

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"errors"`
}

func NewRestErr(code int, message string, err string) *Err {
	return &Err{
		Code:    code,
		Message: message,
		Error:   err,
	}
}
