package errorenum

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
	DBError              = ErrorCode{Err: 500, Msg: "dbu operation error"}
	NotFoundError        = ErrorCode{Err: 404, Msg: "resource not found"}
	UnknownError         = ErrorCode{Err: 500, Msg: "unknown error"}
)
