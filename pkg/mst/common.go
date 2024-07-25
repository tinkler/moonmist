package mst

import (
	"errors"
	"net/http"

	"github.com/tinkler/moonmist/pkg/mlog"
	"gorm.io/gorm"
)

func New(code int, v ...interface{}) Status {
	if len(v) == 0 {
		if code != http.StatusOK {
			mlog.Error(http.StatusText(code))
		}
		return &HttpStatus{code: code, msg: http.StatusText(code)}
	} else {
		msg := mlog.Format(v[0], v[1:]...)
		if code != http.StatusOK {
			mlog.Error(msg)
		}
		return &HttpStatus{code: code, msg: msg}
	}
}

func Unauthorized(v ...interface{}) Status {
	return New(http.StatusUnauthorized, v...)
}

func Forbidden(v ...interface{}) Status {
	return New(http.StatusForbidden, v...)
}

func InternalServerError(v ...interface{}) Status {
	return New(http.StatusInternalServerError, v...)
}

func BadGateway(v ...interface{}) Status {
	return New(http.StatusBadGateway, v...)
}

func BadRequest(v ...interface{}) Status {
	return New(http.StatusBadRequest, v...)
}

func OK(v ...interface{}) Status {
	return New(http.StatusOK, v...)
}

// TODO: v as err
func Any(err error, v ...interface{}) Status {
	if err == nil {
		return nil
	}
	switch err.(type) {
	case *HttpStatus:
		return err.(Status)
	default:
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return OK("no data")
		}
		return InternalServerError(err)
	}
}
