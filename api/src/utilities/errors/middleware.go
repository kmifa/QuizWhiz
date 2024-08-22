package errors

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/morikuni/failure"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorHandler is a function that handles errors.
func ErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		return
	}

	status := GetHTTPStatus(err)
	w.WriteHeader(status)

	if msg, ok := failure.MessageOf(err); ok {
		e := errorResponse{
			Code:    status,
			Message: msg,
		}
		estr, _ := json.Marshal(e)
		io.WriteString(w, string(estr))
	}
}
