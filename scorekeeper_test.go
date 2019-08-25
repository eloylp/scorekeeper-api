package main

import (
	"bytes"
	"fmt"
	"github.com/DATA-DOG/godog"
	"io/ioutil"
	"net/http"
)

func FeatureScoreKeepingContext(s *godog.Suite) {
	s.Step(`^that the scorekeeper service is running$`, thatTheScorekeeperServiceIsRunning)
	s.Step(`^I can add some points$`, iCanAddSomePoints)
	s.Step(`^I can subs some points$`, iCanSubsSomePoints)
	s.Step(`^I cant multiply points$`, iCantMultiplyPoints)
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
