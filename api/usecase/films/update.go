package films

import (
	"simple-crud-app/api/usecase"
	"simple-crud-app/internal/domain"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
	"time"
)

type ReqFilmUpdate struct {
	usecase.CustomHeader
	Film domain.Film `json:"updated_film"`
}

type RespFilmUpdate struct {
	rest.Header
}

func (obj *ReqFilmUpdate) Validate() *errs.Error {
	if obj.Film.ID == 0 {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Film.ID must be not null")
	}
	if obj.Film.Name == "" {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Film.Name must be not null")
	}
	if obj.Film.Score != nil && *obj.Film.Score > 100 {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Film.Score must be between 0 and 100")
	}
	return nil
}

func (obj *ReqFilmUpdate) Execute() (rest.Response, *errs.Error) {
	l := logger.NewLogger().SetID(obj.ReqID).SetMethod("UpdateFilm")
	out := &RespFilmUpdate{} // definition of response struct
	db := obj.GetDB()        // definition of db connection

	l.Infof("Request: %+v", obj)
	defer l.Infof("Response: %+v", out)

	film := models.Film{
		ID:       obj.Film.ID,
		Name:     obj.Film.Name,
		Duration: obj.Film.Duration,
		Score:    obj.Film.Score,
		Comment:  obj.Film.Comment,
		Watched:  obj.Film.Watched,
	}

	if obj.Film.ReleaseDate != nil {
		rd := time.Unix(*obj.Film.ReleaseDate, 0)
		film.ReleaseDate = &rd
	}

	if errApi := film.Update(db); errApi != nil {
		l.Error(errApi)
		return nil, errApi
	}

	l.Infof("film(id:%d) updated")

	return out, nil
}
