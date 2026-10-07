// Package web.
package web

import (
	"context"
	"net/http"
	"uuid"
)

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request)

type App struct {
	*http.ServeMux
	mw []MidFunc
}

func NewApp(mw ...MidFunc) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		mw:       mw,
	}
}

// HandleFunc IS MY NEW HANDLER FUNCTION.
func (a *App) HandleFunc(pattern string, handlerFunc HandlerFunc, mw ...MidFunc) {
	handlerFunc = wrapMiddleware(mw, handlerFunc)
	handlerFunc = wrapMiddleware(a.mw, handlerFunc)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := setTraceID(r.Context(), uuid.New().String())

		handlerFunc(ctx, w, r)
	}

	a.ServeMux.HandleFunc(pattern, h)
}
