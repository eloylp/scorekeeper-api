package webserver

import (
	"context"
	"encoding/json"
	"github.com/eloylp/scorekeeper"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mec07/rununtil"
	"github.com/rs/zerolog/log"
)

const (
	PingEndpoint   = "/ping"
	PointsEndpoint = "/points"
)

type result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// NewRunner returns a function that runs the http webserver
func NewRunner() rununtil.RunnerFunc {
	return func() rununtil.ShutdownFunc {

		s := scorekeeper.NewScorer()
		r := chi.NewRouter()
		r.Get(PingEndpoint, pingHandler)
		r.Post(PointsEndpoint, pointsHandler(&s))
		httpServer := http.Server{Addr: ":8080", Handler: r}
		go runHTTPServer(&httpServer)

		return func() {
			if err := httpServer.Shutdown(context.Background()); err != nil {
				log.Error().Err(err).Msg("error shutting down http server")
			}
		}
	}
}

func runHTTPServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("ListenAndServe")
	}
}

func writeBadRequestResponse(w http.ResponseWriter, err error) {
	result := result{
		Success: false,
		Message: err.Error(),
	}
	resultData, _ := json.Marshal(result)
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(resultData)
}

func writeSuccessResponse(w http.ResponseWriter, msg string) {
	result := result{
		Success: true,
		Message: msg,
	}
	resultData, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resultData)
}
