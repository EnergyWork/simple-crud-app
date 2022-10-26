package films

import (
	"simple-crud-app/api/usecase"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqFilmList struct {
	usecase.CustomHeader
	// TODO SortParameters
	Offset uint64
	Limit  uint64
}

type RespFilmList struct {
	rest.Header
	Films []models.Film
	Total uint64
}

func (obj *ReqFilmList) Execute() (rest.Response, *errs.Error) {
	l := logger.NewLogger().SetID(obj.ReqID).SetMethod("FilmsList")
	out := &RespFilmList{} // definition of response struct
	db := obj.GetDB()      // definition of db connection

	l.Infof("Request: %+v", obj)
	defer l.Infof("Response: %+v", out)

	// default limit
	if obj.Limit == 0 {
		obj.Limit = 5
	}

	filmList := models.FilmList{
		UserID: obj.User.ID,
		Offset: obj.Offset,
		Limit:  obj.Limit,
	}
	list, total, errApi := filmList.GetList(db)
	if errApi != nil {
		l.Error(errApi)
		return nil, errApi
	}

	out.Films = list
	out.Total = total

	return out, nil
}
