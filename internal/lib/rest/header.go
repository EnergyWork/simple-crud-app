package rest

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/models"
)

type Header struct {
	Error     *errs.Error `json:"Error"`
	Signature string      `json:"-"`
	Subject   string      `json:"-"` // user name
	Digest    string      `json:"-"`
}

const (
	SignatureHeader = "X-Signature"
	SubjectHeader   = "X-Subject"
)

// SetHeadet sets request headers to header struct
func (h *Header) SetHeader(r *http.Request) {
	h.Signature = r.Header.Get(SignatureHeader)
	h.Subject = r.Header.Get(SubjectHeader)
}

// SetDigest sets sha256 hash of body
func (h *Header) SetDigest(bts []byte) {
	hash := sha256.Sum256(bts)
	h.Digest = hex.EncodeToString(hash[:])
}

func (h *Header) SetError(err *errs.Error) {
	h.Error = err
}

// Authorize : deafult authorization method
func (h *Header) Authorize(db models.DB) *errs.Error {
	return nil
}
