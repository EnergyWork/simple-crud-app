package rest

import (
	"net/http"

	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/models"
)

type Header struct {
	Error     *errs.Error `json:"Error"`
	Token     string      `json:"-"`
	AccessKey string      `json:"-"`
	Login     string      `json:"-"`
	Digest    string      `json:"-"`
}

const (
	TokenHeader     = "X-Token"     // Session token
	AccessKeyHeader = "X-AccessKey" // Access Key
	LoginHeader     = "X-User"      // User login
)

// SetHeadet sets request headers to header struct
func (h *Header) SetHeader(r *http.Request) {
	h.Token = r.Header.Get(TokenHeader)
	h.AccessKey = r.Header.Get(AccessKeyHeader)
	h.Login = r.Header.Get(LoginHeader)
}

func (h *Header) SetError(err *errs.Error) {
	h.Error = err
}

func (h *Header) Validate() *errs.Error {
	return nil
}

// Authorize : deafult authorization method
func (h *Header) Authorize(db models.DB) *errs.Error {
	return nil
}

func (h *Header) CheckSession(db models.DB) *errs.Error {
	return nil
}
