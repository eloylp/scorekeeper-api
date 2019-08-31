package webserver

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "pong")
}
