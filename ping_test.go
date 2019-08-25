package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func FeaturePingContext(s *godog.Suite) {
	s.Step(`^that the scorekeeper service is running$`, thatTheScorekeeperServiceIsRunning)
	s.Step(`^I can ping the scorekeeper service$`, iCanPingTheScorekeeperService)
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
