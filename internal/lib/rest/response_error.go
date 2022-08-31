package rest

import (
	"encoding/json"
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
)

type Response interface {
	SetError(*errs.Error)
}

func CreateResponseError(w http.ResponseWriter, resp Response, errApi *errs.Error) {
	resp.SetError(errApi)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bts)
}
