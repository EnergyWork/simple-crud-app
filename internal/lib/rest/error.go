package rest

import (
	"encoding/json"
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
)

type RplError struct {
	Error errs.Error `json:"Error"`
}

func (e *RplError) Set(err *errs.Error) {
	e.Error = *err
}

func CreateRplError(w http.ResponseWriter, errApi *errs.Error) {
	rplErr := &RplError{}
	rplErr.Set(errApi)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, err := json.Marshal(rplErr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bts)
}
