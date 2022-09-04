package api

/*
import (
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
	"time"
)

type ReqCreateFilm struct {
	rest.Header
	Name        string     `json:"name"`
	ReleaseDate *time.Time `json:"release_date"`
	Duration    *string    `json:"duration"`
	Score       *uint64    `json:"score"`
	Comment     *string    `json:"comment"`
}

type RespCreateFilm struct {
	rest.Header
}

func (obj *ReqCreateFilm) Validate() *errs.Error {
	if obj.Name == "" {
		return errs.New().SetCode(errs.ERROR_SYNTAX).SetMsg("Name must be not empty")
	}
	return nil
}

// CreateFilm : POST handler for create film requests
func (s *Server) CreateFilm(w http.ResponseWriter, r *http.Request) {
	l := logger.NewLogger().SetMethod("CreateFilm")
	req := &ReqCreateFilm{}
	resp := &RespCreateFilm{}

	// handle request method
	if r.Method != http.MethodPost {
		errApi := errs.New().SetCode(errs.ERROR_METHOD_NOT_ALLOWED).SetMsg("not allowed method - expected POST")
		rest.CreateResponseError(w, resp, errApi)
		l.Errorf("error: wrong request method - req: %s, but expected POST", r.Method)
		return
	}

	//unmarshal input request into struct
	if err := rest.CreateRequest(r, req); err != nil {
		rest.CreateResponseError(w, resp, err)
		l.Errorf("errro: unable create request - %s", err)
		return
	}

	l.Debugf("->REQ: %+v", req)

	// authorization
	if err := req.Authorize(s.GetDB()); err != nil {
		rest.CreateResponseError(w, resp, err)
		l.Errorf("error: unauthorized request - %s", err)
		return
	}

	// business logic
	film := &models.Film{
		Name:        req.Name,
		ReleaseDate: req.ReleaseDate,
		Duration:    req.Duration,
		Score:       req.Score,
		Comment:     req.Comment,
	}
	if err := film.Create(s.GetDB()); err != nil {
		rest.CreateResponseError(w, resp, err)
		l.Errorf("error: %s", err)
		return
	}

	// response
	rest.CreateResponse(w, resp)
	l.Debugf("<-RESP: %+v", resp)
}
*/
