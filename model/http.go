package model

type HTTPRespResult struct {
	Code   int    `json:"-"`
	Err    error  `json:"-"`
	Msg    string `json:"message"`
	Status string `json:"status"`
	Data   string `json:"data,omitempty"`
}

const (
	GinResponseKey = "response"
)
