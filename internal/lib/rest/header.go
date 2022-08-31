package rest

import (
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/models"
)

type Header struct {
	Error   *errs.Error `json:"Error"`
	AuthKey string      `json:"-"`
}

func (h *Header) SetKey(key string) {
	h.AuthKey = key
}

func (h *Header) SetError(err *errs.Error) {
	h.Error = err
}

func (h *Header) Authorize(db models.DB) *errs.Error {
	// models.Auth......
	return nil
}
