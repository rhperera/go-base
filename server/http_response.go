package server

type HttpResponse struct {
	Data  interface{} `json:"data"`
	Error Error       `json:"error"`
}

func (resp *HttpResponse) SetError(code int, msg string) {
	resp.Error.Code = code
	resp.Error.Message = msg
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
