package main

import (
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/eloylp/scorekeeper-api/webserver"
	"net/url"
)

func FeatureScoreKeepingContext(s *godog.Suite) {
	s.Step(`^that the scorekeeper service is running$`, thatTheScorekeeperServiceIsRunning)
	s.Step(`^I add (\d+) points to user "([^"]*)"$`, iAddPointsToUser)
	s.Step(`^I subs (\d+) points to user "([^"]*)"$`, iSubsPointsToUser)
	s.Step(`^"([^"]*)" has now (\d+) points$`, hasNowPoints)
}

func iAddPointsToUser(points int, user string) error {
	eUrl := endPointUrl(webserver.PointsEndpoint)
	sData := []byte(fmt.Sprintf(`{"user": "%s", "points": %v, "opType": "%s"}`, user, points, webserver.OperationAdd))
	_, err := dataToServer(eUrl, sData)
	if err != nil {
		return err
	}
	return nil
}

func iSubsPointsToUser(points int, user string) error {
	eUrl := endPointUrl(webserver.PointsEndpoint)
	sData := []byte(fmt.Sprintf(`{"user": "%s", "points": %v, "opType": "%s"}`, user, points, webserver.OperationSubs))
	_, err := dataToServer(eUrl, sData)
	if err != nil {
		return err
	}
	return nil
}

func hasNowPoints(user string, points int) error {

	eUrl := endPointUrl(webserver.PointsEndpoint)
	pUrl, err := url.Parse(eUrl)
	if err != nil {
		return err
	}
	q := pUrl.Query()
	q.Add("user", user)
	q.Add("points", string(points))
	pUrl.RawQuery = q.Encode()

	data, err := dataFromServer(pUrl.String())
	if err != nil {
		return err
	}

	type response struct {
		Success bool `json:"success"`
		Points  int  `json:"points"`
	}

	r := response{}
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}

	if !r.Success {
		return fmt.Errorf(`expected response to be "true", got: "%v"`, r.Success)
	}

	if r.Points != points {
		return fmt.Errorf(`expected response to be "%v", got: "%v"`, points, r.Points)
	}
	return nil
}
