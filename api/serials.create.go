package api

import (
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
	"time"
)

type ReqCreateSerial struct { // TODO ещё подумать над структурой запроса
	CustomHeader

	Serial struct {
		Name        string     `json:"name"`
		ReleaseDate *time.Time `json:"release_date"`
		Score       *uint64    `json:"score"`
		Comment     *string    `json:"comment"`
	} `json:"serial"`

	Seasons *[]struct {
		Number int               `json:"number"`
		Series map[string]string `json:"series"`
	} `json:"seasons"`
}

type RespCreateSerial struct {
	rest.Header
}

func (obj *ReqCreateSerial) Validate() *errs.Error {
	if obj.Serial.Name == "" {
		return errs.New().SetCode(errs.ERROR_SYNTAX).SetMsg("Name must be not empty")
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
		l.Errorf("errro: unable create request - %s", errApi)
		return
	}

	if errApi := rest.Prepare(s.GetDB(), req); errApi != nil {
		rest.CreateResponseError(w, resp, errApi)
		l.Errorf("error: %s", errApi)
		return
	}

	l.Debugf("->REQ: %+v", req)

	//? ///////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// business logic
	serial := &models.Serial{
		Name:        req.Serial.Name,
		UserID:      req.user.ID,
		ReleaseDate: req.Serial.ReleaseDate,
		Score:       req.Serial.Score,
		Comment:     req.Serial.Comment,
	}
	if err := serial.Create(s.GetDB()); err != nil {
		rest.CreateResponseError(w, resp, err)
		l.Errorf("error: %s", err)
		return
	}

	//? ///////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// response
	rest.CreateResponse(w, resp)
	l.Debugf("<-RESP: %+v", resp)
}
