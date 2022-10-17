package handlers

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

// AuthLogin : user log in
func (s *Server) AuthLogin(w http.ResponseWriter, r *http.Request) {
	//* Setup //
	l := logger.NewLogger().SetMethod("AuthLogin")
	req := &ReqAuthLogin{}
	resp := &RespAuthLogin{}

	//* Form Request //

	// unmarshal input request into struct
	if err := rest.CreateRequest(r, &s.BaseServer, req, http.MethodPost); err != nil {
		rest.CreateResponseError(w, err)
		l.Errorf("error: unable create request - %s", err)
		return
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	hashedPassword, err := hash.NewSHA256Hash(req.Password)
	if err != nil {
		l.Errorf("error: %s", err)
		rest.CreateResponseError(w, err)
		return
	}
	req.Password = hashedPassword

	l.Debugf("Request: %+v", req)
	defer l.Infof("Response: %+v", resp)

	//* Business Logic //

	user, errApi := models.LoadUserByLogin(s.GetDB(), req.Login)
	if errApi != nil {
		l.Errorf("unable to load user: %s", errApi)
		rest.CreateResponseError(w, errApi)
		return
	}

	if user.Password != req.Password {
		errApi := errs.New().SetCode(errs.ErrorForbidden).SetMsg("wrong password")
		rest.CreateResponseError(w, errApi)
		l.Errorf("wrong password")
		return
	}

	session, errApi := models.LoadSession(s.GetDB(), user.SessionID)
	if errApi != nil {
		l.Errorf("unable to load user's session: %s", errApi)
		rest.CreateResponseError(w, errApi)
		return
	}

	session.UpdateTTL(s.GetDB())

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	resp.AccessKey = user.AccessKey
	rest.CreateResponse(w, resp)
}
