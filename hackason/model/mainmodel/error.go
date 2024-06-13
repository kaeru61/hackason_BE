package mainmodel

import "log"

type Error struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func MakeError(code int, detail string) Error {
	if code != 0 {
		log.Printf("error code %v: %v", code, detail)
	}
	return Error{
		Code:   code,
		Detail: detail,
	}
}

func (err *Error) UpdateError(code int, detail string) {
	err.Code = code
	err.Detail += detail
}

var NilError = MakeError(0, "")
