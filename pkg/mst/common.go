package mst

import (
	"net/http"

	"github.com/tinkler/moonmist/pkg/mlog"
)

func New(code int, v ...interface{}) Status {
	if len(v) == 0 {
		mlog.Error(http.StatusText(code))
	}
	return &HttpStatus{code: code, msg: http.StatusText(code)}
}

func Unauthorized(v ...interface{}) Status {
	return New(http.StatusUnauthorized)
}

func Forbidden(v ...interface{}) Status {
	return New(http.StatusForbidden)
}

func InternalServerError(v ...interface{}) Status {
	return New(http.StatusInternalServerError)
}

func BadGateway(v ...interface{}) Status {
	return New(http.StatusBadGateway)
}

func BadRequest(v ...interface{}) Status {
	return New(http.StatusBadRequest)
}

func OK(v ...interface{}) Status {
	return New(http.StatusOK)
}

func Any(err error, v ...interface{}) Status {
	switch err.(type) {
	case *HttpStatus:
		return err.(Status)
	default:
		return InternalServerError(err)
	}
}
