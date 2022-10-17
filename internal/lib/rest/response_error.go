package rest

import (
	"encoding/json"
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
)

type ResponseError struct {
	Error *errs.Error
}

func CreateResponseError(w http.ResponseWriter, errApi *errs.Error) {
	e := &ResponseError{Error: errApi}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, err := json.Marshal(e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bts)
}
