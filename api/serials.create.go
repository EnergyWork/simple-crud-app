package api

import (
	"net/http"
	"simple-crud-app/internal/domain"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqCreateSerial struct {
	CustomHeader
	Serial  domain.Serial    `json:"serial"`
	Seasons *[]domain.Season `json:"seasons"`
}

type RespCreateSerial struct {
	rest.Header
}

func (obj *ReqCreateSerial) Validate() *errs.Error {
	if obj.Serial.Name == "" {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Serial.Name must be not empty")
	}
	if obj.Seasons != nil {
		for _, v := range *obj.Seasons {
			if v.Number == 0 || len(v.Series) == 0 {
				return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Season.Name must be not empty")
			}
		}
	}
	return nil
}

// CreateSerial :
func (s *Server) CreateSerial(w http.ResponseWriter, r *http.Request) {
	l := logger.NewLogger().SetMethod("CreateSerial")
	req := &ReqCreateSerial{}
	resp := &RespCreateSerial{}

	//unmarshal input request into struct
	if errApi := rest.CreateRequest(r, req, http.MethodPost); errApi != nil {
		rest.CreateResponseError(w, resp, errApi)
		l.Errorf("error: unable create request - %s", errApi)
		return
	}

	if errApi := rest.Prepare(s.GetDB(), req); errApi != nil {
		rest.CreateResponseError(w, resp, errApi)
		l.Errorf("error: %s", errApi)
		return
	}

	l.Debugf("->REQ: %+v", req)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// FIXME :CRITICAL: доработать бизнес логику

	tx, _ := s.GetDB().Begin()
	defer func() {
		_ = tx.Commit()
	}()

	// business logic
	serial := &models.Serial{
		Name:        req.Serial.Name,
		UserID:      req.user.ID,
		ReleaseDate: req.Serial.ReleaseDate,
		Score:       req.Serial.Score,
		Comment:     req.Serial.Comment,
	}
	if err := serial.Create(tx); err != nil {
		_ = tx.Rollback()
		rest.CreateResponseError(w, resp, err)
		l.Errorf("error: %s", err)
		return
	}

	for _, season := range *req.Seasons {
		seasonTmp := models.Season{
			SerialID: serial.ID,
			Number:   season.Number,
			Series:   season.Series,
		}
		if err := seasonTmp.Create(tx); err != nil {
			_ = tx.Rollback()
			rest.CreateResponseError(w, resp, err)
			l.Errorf("error: %s", err)
			return
		}
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// response
	rest.CreateResponse(w, resp)
	l.Debugf("<-RESP: %+v", resp)
	return
}
