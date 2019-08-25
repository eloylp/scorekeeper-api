package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/eloylp/scorekeeper-api/webserver"
	"github.com/mec07/rununtil"
	"net"
	"os"
	"testing"
	"time"
)

// This is just to run godog when running go test
func TestMain(m *testing.M) {
	// Register here all the feature contexts. Mind that callbacks will be applied to all contexts.
	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		httpServerContext(s)
		FeatureScoreKeepingContext(s)
		FeaturePingContext(s)
	}, godog.Options{
		Format:    "progress",
		Paths:     []string{"features"},
		Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func httpServerContext(s *godog.Suite) {
	// Before & After steps
	var serverRunner rununtil.RunnerFunc
	var serverShutdown rununtil.ShutdownFunc
	s.BeforeScenario(func(interface{}) {
		serverRunner = webserver.NewRunner()
		serverShutdown = serverRunner()
	})
	s.AfterScenario(func(interface{}, error) {
		serverShutdown()
	})
}

func thatTheScorekeeperServiceIsRunning() error {
	d := net.Dialer{}
	conn, err := d.Dial("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	if err := conn.Close(); err != nil {
		return err
	}
	return nil
}
