// Package web.
package web

import (
	"context"
	"net/http"
	"uuid"
)

// Encoder defines behavior that can encode a data model and provide
// the content type for that encoding.
type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) Encoder

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

		data := handlerFunc(ctx, w, r)

		bytes, ct, err := data.Encode()
		if err != nil {
			// DO SOMETHING
			return
		}

		w.Header().Add("context-type", ct)
		w.Write(bytes)
	}

	a.ServeMux.HandleFunc(pattern, h)
}
