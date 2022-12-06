package films

import (
	"time"

	"simple-crud-app/api/usecase"
	"simple-crud-app/internal"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"

	"github.com/go-playground/validator/v10"
)

type ReqCreateFilm struct {
	usecase.CustomHeader
	Name        string  `json:"name" validate:"required"`
	ReleaseDate *string `json:"release_date"`
	Duration    *string `json:"duration"`
	Score       *uint64 `json:"score"`
	Comment     *string `json:"comment"`
	Watched     bool    `json:"watched"`
}

type RespCreateFilm struct {
	rest.Header
	// todo : возвращать в ответе сам фильм?
}

func (obj *ReqCreateFilm) Validate() *errs.Error {
	err := validator.New().Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("Validation Error: failed in %s parameter", err.Field())
		}
	}
	return nil
}

func (obj *ReqCreateFilm) Execute() (rest.Response, *errs.Error) {
	l := logger.NewLogger().SetID(obj.ReqID).SetMethod("CreateFilm")
	out := &RespCreateFilm{} // definition of response struct
	db := obj.GetDB()        // definition of db connection

	l.Infof("Request: %+v", obj)
	defer l.Infof("Response: %+v", out)

	film := &models.Film{
		Name:     obj.Name,
		UserID:   obj.User.ID,
		Duration: obj.Duration,
		Score:    obj.Score,
		Comment:  obj.Comment,
		Watched:  obj.Watched,
	}

	if obj.ReleaseDate != nil {
		rd, err := time.Parse(internal.ReleaseDateLayout, *obj.ReleaseDate)
		if err != nil {
			l.Errorf("date format error: %s", err)
			return nil, errs.New().SetCode(errs.ErrorRequestSyntax).SetMsg("release_date format error: %s", err)
		}
		film.ReleaseDate = &rd
	}

	if errDb := film.Create(db); errDb != nil {
		l.Errorf("error: %s", errDb)
		return out, errDb
	}

	return out, nil
}
