package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/mec07/rununtil"
	"github.com/pkg/errors"
)

// This is just to run godog when running go test
func TestMain(m *testing.M) {
	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		FeatureContext(s)
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

func thatTheScorekeeperServiceIsRunning() error {
	return iCanPingTheScorekeeperService()
}

func iCanPingTheScorekeeperService() error {
	url := "http://localhost:8080/ping"
	res, err := http.Get(url)
	if err != nil {
		return errors.Wrapf(err, "pinging %s", url)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "reading response body")
	}
	if string(body) != "pong" {
		return fmt.Errorf(`expected response to be "pong", got: "%s"`, string(body))
	}
	return nil
}

func iCanAddSomePoints() error {
	url := "http://localhost:8080/points"
	json := []byte(`{"user": "Bob", "points": 5, "opType": "ADD"}`)
	b := bytes.NewReader(json)
	res, err := http.Post(url, "application/json", b)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	expected := `{"success":true,"message":"Total points for user Bob are now 5"}`
	if string(body) != expected {
		return fmt.Errorf(`expected response to be "%s"", got: "%s"`, expected, string(body))
	}
	return nil
}

func iCanSubsSomePoints() error {
	url := "http://localhost:8080/points"
	json := []byte(`{"user": "Bob", "points": 5, "opType": "SUBS"}`)
	b := bytes.NewReader(json)
	res, err := http.Post(url, "application/json", b)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	expected := `{"success":true,"message":"Total points for user Bob are now 0"}`
	if string(body) != expected {
		return fmt.Errorf(`expected response to be "%s"", got: "%s"`, expected, string(body))
	}
	return nil
}
func iCantMultiplyPoints() error {
	url := "http://localhost:8080/points"
	json := []byte(`{"user": "Bob", "points": 5, "opType": "MULTIPLY"}`)
	b := bytes.NewReader(json)
	res, err := http.Post(url, "application/json", b)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	expected := `{"success":false,"message":"Not a valid scorer operation"}`
	if string(body) != expected {
		return fmt.Errorf(`expected response to be "%s"", got: "%s"`, expected, string(body))
	}
	return nil
}
func FeatureContext(s *godog.Suite) {
	// Before & After steps
	var killServer context.CancelFunc
	s.BeforeScenario(func(interface{}) {
		killServer = rununtil.Killed(main)
	})

	s.AfterScenario(func(interface{}, error) {
		killServer()
		time.Sleep(time.Second * 1)
	})

	// Given steps
	s.Step(`^that the scorekeeper service is running$`, thatTheScorekeeperServiceIsRunning)
	s.Step(`^I can ping the scorekeeper service$`, iCanPingTheScorekeeperService)

	// When steps

	// Then steps
	s.Step(`^that the scorekeeper service is running$`, thatTheScorekeeperServiceIsRunning)
	s.Step(`^I can add some points$`, iCanAddSomePoints)
	s.Step(`^I can subs some points$`, iCanSubsSomePoints)
	s.Step(`^I cant multiply points$`, iCantMultiplyPoints)
}
