// Package mux provides something.
package mux

import (
	"github.com/ardanlabs/service/app/domain/testapp"
	"github.com/ardanlabs/service/foundation/web"
)

func WebAPI(mw ...web.MidFunc) *web.App {
	app := web.NewApp(mw...)

	app.HandleFunc("/test", testapp.Test)

	return app
}
