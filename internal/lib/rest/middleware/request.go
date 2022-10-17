package middleware

import (
	"fmt"
	"net/http"
	"simple-crud-app/internal/lib/logger"
)

type LoggerRequest struct {
	handler http.Handler
}

func NewLoggerRequest(handlerToWarp http.Handler) *LoggerRequest {
	return &LoggerRequest{handler: handlerToWarp}
}

func (obj *LoggerRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mwLogger := logger.NewLogger().SetMethod("Server")
	message := fmt.Sprintf("%s | URL:%s", r.Method, r.URL.Path)
	mwLogger.Info(message)
	obj.handler.ServeHTTP(w, r)
}
