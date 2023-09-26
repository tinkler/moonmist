package gs

import (
	"context"
	"net/http"

	"github.com/tinkler/moonmist/pkg/jsonz/sjson"
	"github.com/tinkler/moonmist/pkg/mlog"
	"github.com/tinkler/moonmist/pkg/mst"
)

const (
	ContentType = "application/json"
)

type Model[T any, S any] struct {
	Data T `json:",omitempty"`
	Args S `json:",omitempty"`
}

type Res[T any, S any] struct {
	Data T `json:",omitempty"`
	Resp S `json:",omitempty"`
}

type simple[T any] struct {
	v T
	f func(ctx context.Context) error
}

func Simple[T any](v T, f func(ctx context.Context) error) simple[T] {
	return simple[T]{v, f}
}

func (h simple[T]) Handle(w http.ResponseWriter, r *http.Request) {
	m := Model[T, any]{Data: h.v}
	err := sjson.Bind(r, &m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := Res[T, any]{Data: m.Data}
	err = h.f(r.Context())
	if err != nil {
		HandleError(w, err)
		return
	}
	HandleResponse(w, res)
}

type simple2[T any, S any] struct {
	v T
	f func(ctx context.Context, args S) error
}

func Simple2[T any, S any](v T, f func(ctx context.Context, args S) error) simple2[T, S] {
	return simple2[T, S]{v, f}
}

func (h simple2[T, S]) Handle(w http.ResponseWriter, r *http.Request) {
	m := Model[T, S]{Data: h.v}
	err := sjson.Bind(r, &m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := Res[T, any]{Data: m.Data}
	err = h.f(r.Context(), m.Args)
	if err != nil {
		HandleError(w, err)
		return
	}
	HandleResponse(w, res)
}

type deliver[T any, S any] struct {
	v T
	f func(ctx context.Context) (S, error)
}

func Deliver[T any, S any](d T, f func(ctx context.Context) (S, error)) deliver[T, S] {
	return deliver[T, S]{d, f}
}

func (h deliver[T, S]) Handle(w http.ResponseWriter, r *http.Request) {
	m := Model[T, any]{Data: h.v}
	err := sjson.Bind(r, &m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := Res[T, S]{Data: m.Data}
	res.Resp, err = h.f(r.Context())
	if err != nil {
		HandleError(w, err)
		return
	}
	HandleResponse(w, res)
}

type deliver2[T any, S any, R any] struct {
	v T
	f func(ctx context.Context, args S) (R, error)
}

func Deliver2[T any, S any, R any](d T, f func(ctx context.Context, args S) (R, error)) deliver2[T, S, R] {
	return deliver2[T, S, R]{d, f}
}

func (h deliver2[T, S, R]) Handle(w http.ResponseWriter, r *http.Request) {
	m := Model[T, S]{Data: h.v}
	err := sjson.Bind(r, &m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := Res[T, R]{Data: m.Data}
	res.Resp, err = h.f(r.Context(), m.Args)
	if err != nil {
		HandleError(w, err)
		return
	}
	HandleResponse(w, res)
}

func HandleResponse(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", ContentType)
	byt, err := sjson.Marshal(map[string]interface{}{
		"code":    0,
		"message": "",
		"data":    res,
	})
	if err != nil {
		mlog.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(byt)
	if err != nil {
		mlog.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func HandleError(w http.ResponseWriter, err error) {
	switch s := err.(type) {
	case mst.Status:
		if s.Code() == http.StatusOK {
			w.Header().Set("Content-Type", ContentType)
			_, _ = w.Write([]byte(`{"code":1,"message":"` + s.Msg() + `"}`))
			return
		}
		http.Error(w, s.Error(), s.Code())
		return
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
