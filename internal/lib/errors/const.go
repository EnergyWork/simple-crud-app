package errors

const (
	ErrorNone = 0

	ErrorRequestSyntax = 400
	ErrorUnauthorized  = 401
	ErrorForbidden     = 403

	ErrorNotFound     = 404
	ErrorUserNotFound = 40401

	ErrorMethodNotAllowed = 405
	ErrorTimeout          = 408
	ErrorAlreadyExists    = 409

	ErrorSessionExpired = 440

	ErrorInternal       = 500
	ErrorNotImplemented = 501
)
