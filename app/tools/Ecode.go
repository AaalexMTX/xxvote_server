package tools

var (
	OK = ECode{Code: 0}
)

type ECode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
