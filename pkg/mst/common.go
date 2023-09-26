package mst

import (
	"net/http"

	"github.com/tinkler/moonmist/pkg/mlog"
)

func New(code int, v ...interface{}) Status {
	if len(v) == 0 {
		if code != http.StatusOK {
			mlog.Error(http.StatusText(code))
		}
		return &HttpStatus{code: code, msg: http.StatusText(code)}
	} else {
		msg := mlog.Format(v[0], v[1:]...)
		mlog.Error(msg)
		return &HttpStatus{code: code, msg: msg}
	}
}

func Unauthorized(v ...interface{}) Status {
	return New(http.StatusUnauthorized)
}

func Forbidden(v ...interface{}) Status {
	return New(http.StatusForbidden)
}

func InternalServerError(v ...interface{}) Status {
	return New(http.StatusInternalServerError, v...)
}

func BadGateway(v ...interface{}) Status {
	return New(http.StatusBadGateway)
}

func BadRequest(v ...interface{}) Status {
	return New(http.StatusBadRequest)
}

func OK(v ...interface{}) Status {
	return New(http.StatusOK, v...)
}

func Any(err error, v ...interface{}) Status {
	switch err.(type) {
	case *HttpStatus:
		return err.(Status)
	default:
		return InternalServerError(err)
	}
}
