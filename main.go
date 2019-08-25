package main

import (
	"github.com/eloylp/scorekeeper-api/webserver"
	"github.com/mec07/rununtil"
)

func main() {
	rununtil.KillSignal(webserver.NewRunner())
}
