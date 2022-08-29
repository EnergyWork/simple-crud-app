package errors

import "fmt"

type Error struct {
	Code int    `json:"Code"`
	Msg  string `json:"Message"`
}

func New() *Error {
	return &Error{}
}

func (e *Error) SetCode(code int) *Error {
	e.Code = code
	return e
}

func (e *Error) SetMsg(msg string, args ...interface{}) *Error {
	e.Msg = fmt.Sprintf(msg, args...)
	return e
}

func (e *Error) String() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Msg)
}

func (e *Error) Error() error {
	return fmt.Errorf("%d - %s", e.Code, e.Msg)
}
