package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
)

func FeaturePingContext(s *godog.Suite) {
	s.Step(`^that the scorekeeper service is running$`, thatTheScorekeeperServiceIsRunning)
	s.Step(`^I can ping the scorekeeper service$`, iCanPingTheScorekeeperService)
}

func iCanPingTheScorekeeperService() error {
	url := url("/ping")
	body, err := getDataFromServer(url)
	if err != nil {
		return err
	}
	if string(body) != "pong" {
		return fmt.Errorf(`expected response to be "pong", got: "%s"`, string(body))
	}
	return nil
}
