package api

import (
	"net/http"
	"simple-crud-app/internal/lib/crypto"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/hash"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
	"time"
)

type ReqAuthRegister struct {
	rest.Header
	UserName     string `json:"name"`
	UserPassword string `json:"password"`
}

type RespAuthRegister struct {
	rest.Header
	PublicKey string // base64 crypto.PublicKey
	ExpiredAt int64
}

func (obj *ReqAuthRegister) Validate() *errs.Error {
	if obj.UserName == "" {
		return errs.New().SetCode(errs.ERROR_SYNTAX).SetMsg("UserName must be not empty")
	}
	if obj.UserPassword == "" {
		return errs.New().SetCode(errs.ERROR_SYNTAX).SetMsg("UserPassword must be not empty")
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

	// handle request method
	if r.Method != http.MethodPost {
		errApi := errs.New().SetCode(errs.ERROR_METHOD_NOT_ALLOWED).SetMsg("not allowed method - expected POST")
		rest.CreateResponseError(w, resp, errApi)
		l.Errorf("error: wrong request method - req: %s, but expected POST", r.Method)
		return
	}
	// unmarshal input request into struct
	if err := rest.CreateRequest(r, req); err != nil {
		rest.CreateResponseError(w, resp, err)
		l.Errorf("errro: unable create request - %s", err)
		return
	}

	//* Business Logic //

	hashedPassword, err := hash.NewSHA256Hash(req.UserName)
	if err != nil {
		l.Errorf("error: %s", err)
		rest.CreateResponseError(w, resp, err)
		return
	}
	req.UserPassword = hashedPassword

	l.Debugf("->REQ: %+v", req)

	a := models.UserAuth{
		UserName:     req.UserName,
		UserPassword: req.UserPassword,
	}

	expiredTime := time.Now().AddDate(0, 0, 1)

	privateKey, err := crypto.NewPrivateKey(1024)
	if err != nil {
		l.Error(err)
		rest.CreateResponseError(w, resp, err)
		return
	}

	a.PrivateKey = *privateKey
	a.ExpiredAt = expiredTime

	if err := a.UserRegister(s.GetDB()); err != nil {
		l.Errorf("error: %s", err)
		rest.CreateResponseError(w, resp, err)
		return
	}

	resp.PublicKey = privateKey.Public().GetBase64()
	resp.ExpiredAt = expiredTime.Unix()

	//* Response //
	rest.CreateResponse(w, resp)
	l.Debugf("<-RESP: %+v", resp)
}
