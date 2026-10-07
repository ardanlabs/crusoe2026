// Package testapp
package testapp

import (
	"context"
	"encoding/json"
	"net/http"

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
	return status{
		Status: "OK",
	}
}
