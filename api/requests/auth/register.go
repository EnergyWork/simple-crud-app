package auth

import (
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/rest"
)

type ReqAuthRegister struct {
	rest.Header
	Login    string `json:"name"`
	Password string `json:"password"`
}

type RespAuthRegister struct {
	rest.Header
}

func (obj *ReqAuthRegister) Validate() *errs.Error {
	if obj.Login == "" {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Login must be not empty")
	}
	if obj.Password == "" {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Password must be not empty")
	}
	return nil
}

func (obj *ReqAuthRegister) Execute() (rest.Response, *errs.Error) {
	// TODO
	return nil, errs.New().SetCode(errs.ErrorInternal)
}
