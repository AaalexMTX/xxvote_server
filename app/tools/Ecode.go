package tools

var (
	OK             = ECode{Code: 0, Message: "Success"}
	UserErr        = ECode{Code: 1010, Message: "User Not Exist"}
	CookieErr      = ECode{Code: 1011, Message: "Cookie Not Exist"}
	UniversalECode = ECode{Code: 1110, Message: "Error"}
)

type ECode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
