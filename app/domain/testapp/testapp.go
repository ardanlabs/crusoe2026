// Package testapp
package testapp

import (
	"context"
	"encoding/json"
	"net/http"
)

func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)
}
