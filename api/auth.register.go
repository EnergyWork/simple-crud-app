package api

import (
	"net/http"

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
}

func (obj *ReqAuthRegister) Validate() *errs.Error {
	if obj.Login == "" {
		return errs.New().SetCode(errs.ERROR_SYNTAX).SetMsg("Login must be not empty")
	}
	if obj.Password == "" {
		return errs.New().SetCode(errs.ERROR_SYNTAX).SetMsg("Password must be not empty")
	}
	return nil
}

// AuthRegister : POST handler for user registration
func (s *Server) AuthRegister(w http.ResponseWriter, r *http.Request) {
	//* Setup //
	l := logger.NewLogger().SetMethod("AuthRegister")
	req := &ReqAuthRegister{}
	resp := &RespAuthRegister{}

	//* Form Request //

	// unmarshal input request into struct
	if errApi := rest.CreateRequest(r, req, http.MethodPost); errApi != nil {
		rest.CreateResponseError(w, resp, errApi)
		l.Errorf("errro: unable create request - %s", errApi)
		return
	}

	//? ///////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//* Business Logic //

	// getting password hash
	hashedPassword, errApi := hash.NewSHA256Hash(req.Password)
	if errApi != nil {
		l.Errorf("error: %s", errApi)
		rest.CreateResponseError(w, resp, errApi)
		return
	}
	req.Password = hashedPassword

	l.Debugf("->REQ: %+v", req)

	tx, err := s.GetDB().Begin()
	if err != nil {
		l.Error("unable to create sql trx")
		rest.CreateResponseError(w, resp, errs.New().SetCode(errs.ERROR_INTERNAL))
		return
	}
	defer tx.Rollback() // tx.Commit will be earlier

	// create session
	session, errApi := models.NewSession(tx)
	if errApi != nil {
		l.Error(errApi)
		rest.CreateResponseError(w, resp, errs.New().SetCode(errs.ERROR_INTERNAL))
		return
	}
	// create access key (Subject)
	accessKey, errApi := hash.NewAccessKey(req.Password)
	if errApi != nil {
		l.Error(errApi)
		rest.CreateResponseError(w, resp, errs.New().SetCode(errs.ERROR_INTERNAL))
		return
	}
	// create user with new session
	user := models.User{
		Login:     req.Login,
		Password:  req.Password,
		SessionID: session.ID,
		AccessKey: accessKey,
	}
	if errApi := user.Create(tx); errApi != nil {
		l.Error(errApi)
		rest.CreateResponseError(w, resp, errs.New().SetCode(errs.ERROR_INTERNAL))
		return
	}

	//? ///////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//* Response //
	tx.Commit()
	rest.CreateResponse(w, resp)
	l.Debugf("<-RESP: %+v", resp)
}
