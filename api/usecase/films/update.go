package films

import (
	"simple-crud-app/api/usecase"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqFilmUpdate struct {
	usecase.CustomHeader
	Film models.Film `json:"updated_film"`
}

type RespFilmUpdate struct {
	rest.Header
}

func (obj *ReqFilmUpdate) Validate() *errs.Error {
	//todo: расширить валидацию
	if obj.Film.ID == 0 {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Film.ID must be not null")
	}
	return nil
}

func (obj *ReqFilmUpdate) Execute() (rest.Response, *errs.Error) {
	l := logger.NewLogger().SetID(obj.ReqID).SetMethod("UpdateFilm")
	out := &RespFilmUpdate{} // definition of response struct
	db := obj.GetDB()        // definition of db connection

	l.Infof("Request: %+v", obj)
	defer l.Infof("Response: %+v", out)

	if errApi := obj.Film.Update(db); errApi != nil {
		l.Error(errApi)
		return nil, errApi
	}

	return out, nil
}
