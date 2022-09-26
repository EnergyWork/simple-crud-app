package api

import (
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/hash"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqAuthLogin struct {
	CustomHeader
	Login    string `json:"name"`
	Password string `json:"password"`
}

type RespAuthLogin struct {
	rest.Header
	AccessKey    string // user's acccess key
	SessionToken string // session token
	Deadline     int64  // unix format, token's deadline
}

func (obj *ReqAuthLogin) Validate() *errs.Error {
	if obj.Login == "" {
		return errs.New().SetCode(errs.ERROR_SYNTAX).SetMsg("Login must be not empty")
	}
	if obj.Password == "" {
		return errs.New().SetCode(errs.ERROR_SYNTAX).SetMsg("Password must be not empty")
	}
	return nil
}

// AuthRegister : POST handler for user registration
func (s *Server) AuthLogin(w http.ResponseWriter, r *http.Request) {
	//* Setup //
	l := logger.NewLogger().SetMethod("AuthLogin")
	req := &ReqAuthLogin{}
	resp := &RespAuthLogin{}

	//* Form Request //

	// unmarshal input request into struct
	if err := rest.CreateRequest(r, req, http.MethodPost); err != nil {
		rest.CreateResponseError(w, resp, err)
		l.Errorf("errro: unable create request - %s", err)
		return
	}

	//? ///////////////////////////////////////////////////////////////////////////////////////////////////////////////

	hashedPassword, err := hash.NewSHA256Hash(req.Password)
	if err != nil {
		l.Errorf("error: %s", err)
		rest.CreateResponseError(w, resp, err)
		return
	}
	req.Password = hashedPassword

	l.Debugf("->REQ: %+v", req)

	//* Business Logic //

	user, errApi := models.LoadUserByLogin(s.GetDB(), req.Login)
	if errApi != nil {
		l.Errorf("unable to load user: %s", errApi)
		rest.CreateResponseError(w, resp, errApi)
		return
	}

	if user.Password != req.Password {
		errApi := errs.New().SetCode(errs.ERROR_FORBIDDEN).SetMsg("wrong password")
		rest.CreateResponseError(w, resp, errApi)
		l.Errorf("wrong password")
		return
	}

	session, errApi := models.LoadSession(s.GetDB(), user.SessionID)
	if errApi != nil {
		l.Errorf("unable to load user's session: %s", errApi)
		rest.CreateResponseError(w, resp, errApi)
		return
	}

	session.UpdateTTL(s.GetDB())

	//? ///////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//* Response //
	resp.AccessKey = user.AccessKey
	resp.SessionToken = session.Token
	resp.Deadline = session.Deadline.Unix()
	rest.CreateResponse(w, resp)
	l.Debugf("<-RESP: %+v", resp)
}
