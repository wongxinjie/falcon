package apienum

import "encoding/json"

type ErrorCode struct {
	Err  int                    `json:"err"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data,omitempty"`
	Desc map[string][]int       `json:"desc,omitempty"`
}

func (e ErrorCode) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func (e ErrorCode) HasError() bool {
	return e.Err != 0
}

func (e *ErrorCode) WithData(data map[string]interface{}) *ErrorCode {
	e.Data = data
	return e
}

var (
	NIL                  = ErrorCode{Err: 0, Msg: "success"}
	ErrorInvalidArgument = ErrorCode{Err: 1, Msg: "invalid argument"}
	ErrorUnKnown         = ErrorCode{Err: 2, Msg: "invalid argument"}

	ErrorUnAuthorized = ErrorCode{Err: 401, Msg: "unAuthorized"}
	NotFoundError     = ErrorCode{Err: 404, Msg: "resource not found"}
	DBError           = ErrorCode{Err: 500, Msg: "db operation error"}
)

var (
	ErrorAccountOrPasswordNotMatch = ErrorCode{Err: 100001, Msg: "account not exists or password error"}
)
