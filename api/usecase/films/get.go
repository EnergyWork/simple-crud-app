package films

import (
	"simple-crud-app/api/usecase"
	"simple-crud-app/internal/domain"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqFilmGet struct {
	usecase.CustomHeader
	ID uint64 `json:"id"`
}

type RespFilmGet struct {
	rest.Header
	Film *domain.Film `json:"film"`
}

func (obj *ReqFilmGet) Validate() *errs.Error {
	if obj.ID != 0 {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("ID must be not null")
	}
	return nil
}

func (obj *ReqFilmGet) Execute() (rest.Response, *errs.Error) {
	l := logger.NewLogger().SetID(obj.ReqID).SetMethod("GetFilm")
	out := &RespFilmGet{} // definition of response struct
	db := obj.GetDB()     // definition of db connection

	l.Infof("Request: %+v", obj)
	defer l.Infof("Response: %+v", out)

	film, err := models.LoadFilmByID(db, obj.ID)
	if err != nil {
		l.Error(err)
		return nil, err
	}

	out.Film = film.Convert()

	return out, nil
}
