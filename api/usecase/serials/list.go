package serials

import (
	"simple-crud-app/api/usecase"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/lib/logger"
	"simple-crud-app/internal/lib/rest"
	"simple-crud-app/internal/models"
)

type ReqSerialList struct {
	usecase.CustomHeader
	Offset uint64
	Limit  uint64
}

type RespSerialList struct {
	rest.Header
	Serials []models.SerialFull
	Total   uint64
}

func (obj *ReqSerialList) Validate() *errs.Error {
	// todo
	return nil
}

func (obj *ReqSerialList) Execute() (rest.Response, *errs.Error) {
	l := logger.NewLogger().SetID(obj.ReqID).SetMethod("SerialsList")
	out := &RespSerialList{} // definition of response struct
	db := obj.GetDB()        // definition of db connection

	l.Infof("Request: %+v", obj)
	defer l.Infof("Response: %+v", out)

	serialList := models.SerialList{
		UserID: obj.User.ID,
		Offset: obj.Offset,
		Limit:  obj.Limit,
	}
	tx, _ := db.Begin()
	list, total, errApi := serialList.GetList(tx)
	if errApi != nil {
		_ = tx.Rollback()
		l.Error(errApi)
		return nil, errApi
	}
	_ = tx.Commit()

	out.Serials = list
	out.Total = total

	return out, nil
}
