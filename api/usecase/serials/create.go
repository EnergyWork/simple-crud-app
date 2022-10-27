package serials

import (
	"simple-crud-app/api/usecase"
	"simple-crud-app/internal/domain"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqCreateSerial struct {
	usecase.CustomHeader
	Serial  domain.Serial    `json:"serial"`
	Seasons *[]domain.Season `json:"seasons"`
}

type RespCreateSerial struct {
	rest.Header
}

func (obj *ReqCreateSerial) Validate() *errs.Error {
	return nil
}

func (obj *ReqCreateSerial) Execute() (rest.Response, *errs.Error) {
	l := logger.NewLogger().SetID(obj.ReqID).SetMethod("CreateSerial")
	out := &RespCreateSerial{} // definition of response struct
	db := obj.GetDB()          // definition of db connection

	l.Infof("Request: %+v", obj)
	defer l.Infof("Response: %+v", out)

	// FIXME :CRITICAL: доработать бизнес логику

	tx, _ := db.Begin()
	defer func() {
		_ = tx.Commit()
	}()

	// business logic
	serial := &models.Serial{
		Name:        obj.Serial.Name,
		UserID:      obj.User.ID,
		ReleaseDate: obj.Serial.ReleaseDate,
		Score:       obj.Serial.Score,
		Comment:     obj.Serial.Comment,
	}
	if err := serial.Create(tx); err != nil {
		_ = tx.Rollback()
		l.Errorf("error: %s", err)
		return nil, err
	}

	for _, season := range *obj.Seasons {
		seasonTmp := models.Season{
			SerialID: serial.ID,
			Number:   season.Number,
			Series:   season.Series,
		}
		if err := seasonTmp.Create(tx); err != nil {
			_ = tx.Rollback()
			l.Errorf("error: %s", err)
			return nil, err
		}
	}

	return out, nil
}
