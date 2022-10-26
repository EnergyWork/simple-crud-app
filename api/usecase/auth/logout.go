package auth

import (
	"simple-crud-app/api/usecase"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqAuthLogout struct {
	usecase.CustomHeader
}

type RespAuthLogout struct {
	rest.Header
}

func (obj *ReqAuthLogout) Execute() (rest.Response, *errs.Error) {
	l := logger.NewLogger().SetMethod("AuthLogout")

	out := &RespAuthLogout{}
	db := obj.GetDB() // definition of db connection

	l.Infof("Request: %+v", obj)
	defer l.Infof("Response: %+v", out)

	user, errApi := models.LoadUserByAccessKey(db, obj.AccessKey)
	if errApi != nil {
		l.Errorf("%s", errApi)
		return nil, errApi
	}

	if errApi = user.CloseSession(db); errApi != nil {
		l.Errorf("%s", errApi)
		return nil, errApi
	}

	return out, nil
}
