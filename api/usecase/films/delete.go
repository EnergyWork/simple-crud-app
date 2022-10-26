package films

import (
	"simple-crud-app/api/usecase"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqFilmDelete struct {
	usecase.CustomHeader
	ID uint64 `json:"film_id"`
}

type RespFilmDelete struct {
	rest.Header
}

func (obj *ReqFilmDelete) Validate() *errs.Error {
	if obj.ID != 0 {
		return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("ID must be not null")
	}
	return nil
}

func (obj *ReqFilmDelete) Execute() (rest.Response, *errs.Error) {
	l := logger.NewLogger().SetID(obj.ReqID).SetMethod("DeleteFilm")
	out := &RespFilmDelete{} // definition of response struct
	db := obj.GetDB()        // definition of db connection

	l.Infof("Request: %+v", obj)
	defer l.Infof("Response: %+v", out)

	if errApi := models.DeleteFilmByID(db, obj.User.ID, obj.ID); errApi != nil {
		l.Error(errApi)
		return nil, errApi
	}

	return out, nil
}
