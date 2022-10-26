package auth

import (
	"simple-crud-app/api/usecase"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/hash"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqAuthLogin struct {
	usecase.CustomHeader
	Login    string `json:"name"`
	Password string `json:"password"`
}

type RespAuthLogin struct {
	rest.Header
	AccessKey    string `json:",omitempty"` // user's access key
	SessionToken string `json:",omitempty"` // session token
	Deadline     int64  `json:",omitempty"` // unix format, token's deadline
}

func (obj *ReqAuthLogin) Validate() *errs.Error {
	if obj.Login == "" {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Login must be not empty")
	}
	if obj.Password == "" {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Password must be not empty")
	}
	return nil
}

func (obj *ReqAuthLogin) Execute() (rest.Response, *errs.Error) {
	l := logger.NewLogger().SetMethod("AuthLogin")

	out := &RespAuthLogin{}
	db := obj.GetDB() // definition of db connection

	l.Infof("Request: %+v", obj)
	defer l.Infof("Response: %+v", out)

	hashedPassword, err := hash.NewSHA256Hash(obj.Password)
	if err != nil {
		l.Errorf("error: %s", err)
		return nil, errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err.Error())
	}
	obj.Password = hashedPassword

	user, errApi := models.LoadUserByLogin(db, obj.Login)
	if errApi != nil {
		l.Errorf("unable to load user: %s", errApi)
		return nil, errApi
	}

	if user.Password != obj.Password {
		return nil, errs.New().SetCode(errs.ErrorForbidden).SetMsg("wrong password")
	}

	session, errApi := models.LoadSession(db, user.SessionID)
	if errApi != nil {
		l.Errorf("unable to load user's session: %s", errApi)
		return nil, errApi
	}

	session.UpdateTTL(db)

	out.AccessKey = user.AccessKey

	return out, nil
}
