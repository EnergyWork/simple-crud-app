package errors

import "fmt"

type Error struct {
	Code int    `json:"Errno"`
	Msg  string `json:"Error"`
}

func New() *Error {
	return &Error{}
}

func (e *Error) SetCode(code int) *Error {
	e.Code = code
	return e
}

func (e *Error) SetMsg(msg string) *Error {
	e.Msg = msg
	return e
}

func (e *Error) String() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Msg)
}

func (e *Error) Error() error {
	return fmt.Errorf("%d - %s", e.Code, e.Msg)
}
