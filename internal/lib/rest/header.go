package rest

import (
	"net/http"

	"golang.org/x/text/language"

	errs "simple-crud-app/internal/lib/errors"
	"simple-crud-app/internal/models"
)

type Header struct {
	Error     *errs.Error  `json:"Error"`
	Token     string       `json:"-"`
	AccessKey string       `json:"-"`
	Login     string       `json:"-"`
	Digest    string       `json:"-"`
	ReqID     string       `json:"-"`
	Language  language.Tag `json:"-"`
}

const (
	TokenHeader     = "X-Token"         // Session token
	AccessKeyHeader = "X-AccessKey"     // Access Key
	LoginHeader     = "X-User"          // User login
	AcceptLanguage  = "Accept-Language" // Request language
)

// SetHeader sets request headers to header struct
func (h *Header) SetHeader(r *http.Request) {
	h.Token = r.Header.Get(TokenHeader)
	h.AccessKey = r.Header.Get(AccessKeyHeader)
	h.Login = r.Header.Get(LoginHeader)
	lang := r.Header.Get(AcceptLanguage)
	tag, err := language.Parse(lang)
	if err != nil {
		h.Language = language.English
	} else {
		h.Language = tag
	}
}

func (h *Header) SetError(err *errs.Error) {
	h.Error = err
}

func (h *Header) SetReqID(reqID string) {
	if h.ReqID == "" {
		h.ReqID = reqID
	}
}

func (h *Header) Validate() *errs.Error {
	return nil
}

// Authorize : default authorization method
func (h *Header) Authorize(db models.DB) *errs.Error {
	return nil
}

func (h *Header) CheckSession(db models.DB) *errs.Error {
	return nil
}
