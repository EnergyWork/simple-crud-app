package api

import (
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqSerialList struct {
	CustomHeader
	Offset uint64
	Limit  uint64
}

type RespSerialList struct {
	rest.Header
	Serials []models.Serial
	Total   uint64
}

func (obj *ReqSerialList) Validate() *errs.Error {
	return nil
}

// SerialList :
func (s *Server) SerialList(w http.ResponseWriter, r *http.Request) {
	l := logger.NewLogger().SetMethod("SerialList")
	req := &ReqSerialList{}
	resp := &RespSerialList{}

	//unmarshal input request into struct
	if err := rest.CreateRequest(r, req, http.MethodPost); err != nil {
		rest.CreateResponseError(w, resp, err)
		l.Errorf("errro: unable create request - %s", err)
		return
	}

	if errApi := rest.Prepare(s.GetDB(), req); errApi != nil {
		rest.CreateResponseError(w, resp, errApi)
		l.Errorf("error: %s", errApi)
		return
	}

	l.Debugf("->REQ: %+v", req)

	//? business logic //////////////////////////////////////////////////////////////////////////////////////////////////////////

	serialList := models.SerialList{
		UserID: req.user.ID,
		Offset: req.Offset,
		Limit:  req.Limit,
	}
	list, total, errApi := serialList.GetList(s.GetDB())
	if errApi != nil {
		rest.CreateResponseError(w, resp, errApi)
		l.Error(errApi)
		return
	}

	//? response /////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	resp.Serials = list
	resp.Total = total
	rest.CreateResponse(w, resp)
	l.Debugf("<-RESP: %+v", resp)
}