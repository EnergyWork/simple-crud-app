package api

import (
	"github.com/google/uuid"
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
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Login must be not empty")
	}
	if obj.Password == "" {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Password must be not empty")
	}
	return nil
}

// AuthRegister : POST handler for user registration
func (s *Server) AuthRegister(w http.ResponseWriter, r *http.Request) {
	//* Setup //
	l := logger.NewLogger().SetMethod("AuthRegister").SetID(uuid.NewString())
	req := &ReqAuthRegister{}
	resp := &RespAuthRegister{}

	//* Form Request //

	// unmarshal input request into struct
	if errApi := rest.CreateRequest(r, req, http.MethodPost); errApi != nil {
		rest.CreateResponseError(w, resp, errApi)
		l.Errorf("error: unable create request - %s", errApi)
		return
	}

	//? ///////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// getting password hash
	hashedPassword, errApi := hash.NewSHA256Hash(req.Password)
	if errApi != nil {
		l.Errorf("error: %s", errApi)
		rest.CreateResponseError(w, resp, errApi)
		return
	}
	req.Password = hashedPassword

	l.Debugf("Request: %+v", req)
	defer l.Debugf("Response: %+v", resp)

	tx, _ := s.GetDB().Begin()

	defer func() {
		_ = tx.Rollback() // tx.Commit will be earlier
	}()

	// create session
	session, errApi := models.NewSession(tx)
	if errApi != nil {
		l.Error(errApi)
		rest.CreateResponseError(w, resp, errApi)
		return
	}
	// create access key (Subject)
	accessKey, errApi := hash.NewAccessKey(req.Password)
	if errApi != nil {
		l.Error(errApi)
		rest.CreateResponseError(w, resp, errApi)
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
		rest.CreateResponseError(w, resp, errApi)
		return
	}

	//? ///////////////////////////////////////////////////////////////////////////////////////////////////////////////
	_ = tx.Commit()
	rest.CreateResponse(w, resp)
}
