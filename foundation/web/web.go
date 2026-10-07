// Package web.
package web

import (
	"context"
	"net/http"
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

		// WE CAN PUT ANY CODE HERE

		handlerFunc(r.Context(), w, r)

		// WE CAN PUT ANY CODE HERE
	}

	a.ServeMux.HandleFunc(pattern, h)
}
