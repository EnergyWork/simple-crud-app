package rest

import (
	"encoding/json"
	"net/http"
)

type Response interface {
}

func CreateResponse(w http.ResponseWriter, resp Response) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bts)
}
