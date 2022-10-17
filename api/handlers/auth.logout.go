package handlers

import (
	"net/http"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqAuthLogout struct {
	CustomHeader
}

type RespAuthLogout struct {
	rest.Header
}

// AuthLogout ...
func (s *Server) AuthLogout(w http.ResponseWriter, r *http.Request) {
	//* Setup //
	l := logger.NewLogger().SetMethod("AuthLogout")
	req := &ReqAuthLogout{}
	resp := &RespAuthLogout{}

	//* Form Request //

	// unmarshal input request into struct
	if err := rest.CreateRequest(r, &s.BaseServer, req, http.MethodGet); err != nil {
		rest.CreateResponseError(w, err)
		l.Errorf("errro: unable create request - %s", err)
		return
	}

	if errApi := rest.Prepare(s.GetDB(), req); errApi != nil {
		rest.CreateResponseError(w, errApi)
		l.Errorf("error: %s", errApi)
		return
	}

	l.Debugf("->REQ: %+v", req)

	//? ///////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//* Business Logic //

	user, errApi := models.LoadUserByAccessKey(s.GetDB(), req.AccessKey)
	if errApi != nil {
		rest.CreateResponseError(w, errApi)
		l.Errorf("%s", errApi)
		return
	}

	if errApi := user.CloseSession(s.GetDB()); errApi != nil {
		rest.CreateResponseError(w, errApi)
		l.Errorf("%s", errApi)
		return
	}

	//* Response //
	rest.CreateResponse(w, resp)
	l.Debugf("<-RESP: %+v", resp)
}
