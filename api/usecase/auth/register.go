package auth

import (
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/hash"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqAuthRegister struct {
	rest.Header
	Login    string `json:"name"`
	Password string `json:"password"`
}

type RespAuthRegister struct {
	rest.Header
	AccessKey string
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
	l := logger.NewLogger().SetMethod("Registration").SetID(obj.ReqID) // configure a logger

	out := &RespAuthRegister{} // definition of response struct
	db := obj.GetDB()          // definition of db connection

	l.Infof("Request: %+v", obj)
	defer l.Infof("Response: %+v", out)

	// getting password hash
	hashedPassword, errApi := hash.NewSHA256Hash(obj.Password)
	if errApi != nil {
		l.Errorf("error: %s", errApi)
		return nil, errApi
	}
	obj.Password = hashedPassword

	tx, _ := db.Begin()
	defer func() {
		_ = tx.Rollback() // tx.Commit will be earlier
	}()

	// create session
	session, errApi := models.NewSession(tx)
	if errApi != nil {
		l.Error(errApi)
		return nil, errApi
	}
	// create access key (Subject)
	accessKey, errApi := hash.NewAccessKey(obj.Password)
	if errApi != nil {
		l.Error(errApi)
		return nil, errApi
	}
	// create user with new session
	user := models.User{
		Login:     obj.Login,
		Password:  obj.Password,
		SessionID: session.ID,
		AccessKey: accessKey,
	}
	if errApi = user.Create(tx); errApi != nil {
		l.Error(errApi)
		return nil, errApi
	}

	_ = tx.Commit()

	out.AccessKey = user.AccessKey

	return out, nil
}
