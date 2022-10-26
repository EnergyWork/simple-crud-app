package films

import (
	"fmt"
	"simple-crud-app/api/usecase"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqCreateFilm struct {
	usecase.CustomHeader
	Name        string     `json:"name" validate:"required"`
	ReleaseDate *time.Time `json:"release_date"`
	Duration    *string    `json:"duration"`
	Score       *uint64    `json:"score"`
	Comment     *string    `json:"comment"`
}

type RespCreateFilm struct {
	rest.Header
}

func (obj *ReqCreateFilm) Validate() *errs.Error {
	eng := en.New()
	uni := ut.New(eng, eng)

	trans, ok := uni.GetTranslator(obj.Language.Parent().String())
	if !ok {
		fmt.Println("translator not found")
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("translator not found")
	}

	validate := validator.New()

	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		fmt.Println(1, err)
		return errs.New().SetCode(errs.ErrorInternal).SetMsg(err.Error())
	}

	err := validate.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Translate(trans))
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
		Name:        obj.Name,
		UserID:      obj.User.ID,
		ReleaseDate: obj.ReleaseDate,
		Duration:    obj.Duration,
		Score:       obj.Score,
		Comment:     obj.Comment,
	}
	if errDb := film.Create(db); errDb != nil {
		l.Errorf("error: %s", errDb)
		return out, errDb
	}

	return out, nil
}
