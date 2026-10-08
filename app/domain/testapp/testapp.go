// Package testapp
package testapp

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/web"
)

type status struct {
	Status string
}

func (s status) Encode() ([]byte, string, error) {
	d, err := json.Marshal(s)
	return d, "json", err
}

func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) web.Encoder {
	if n := rand.Intn(100); n%2 == 0 {
		return errs.Errorf(errs.NotFound, "This is an test error")
	}

	return status{
		Status: "OK",
	}
}
