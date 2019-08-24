package webserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eloylp/scorekeeper"
	"io/ioutil"
	"net/http"
)

const (
	OperationAdd  = "ADD"
	OperationSubs = "SUBS"
)

func pointsHandler(scorer *scorekeeper.Scorer) http.HandlerFunc {

	type operation struct {
		User   string `json:"user"`
		Points int    `json:"points"`
		OpType string `json:"opType"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		o := &operation{}
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeBadRequestResponse(w, err)
			return
		}
		if err := json.Unmarshal(data, o); err != nil {
			writeBadRequestResponse(w, err)
		}

		switch o.OpType {
		case OperationAdd:
			scorer.Add(o.User, o.Points)
			break
		case OperationSubs:
			scorer.Subs(o.User, o.Points)
			break
		default:
			writeBadRequestResponse(w, errors.New("Not a valid scorer operation"))
			return
		}
		writeSuccessResponse(w, fmt.Sprintf("Total points for user %s are now %v", o.User, scorer.Get(o.User)))
	}
}
