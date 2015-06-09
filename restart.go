package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

var cmdRestart = &Command{
	Exec:        runRestart,
	UsageLine:   "restart [OPTIONS] SERVER [SERVER...]",
	Description: "Restart a running server",
	Help:        "Restart a running server.",
}

func init() {
	cmdRestart.Flag.BoolVar(&restartHelp, []string{"h", "-help"}, false, "Print usage")
}

// Flags
var restartHelp bool // -h, --help flag

func runRestart(cmd *Command, args []string) {
	if restartHelp {
		cmd.PrintUsage()
	}
	if len(args) < 1 {
		cmd.PrintShortUsage()
	}

	hasError := false
	for _, needle := range args {
		server := cmd.API.GetServerID(needle)
		err := cmd.API.PostServerAction(server, "reboot")
		if err != nil {
			if err.Error() != "server is being stopped or rebooted" {
				log.Errorf("failed to restart server %s: %s", server, err)
				hasError = true
			}
		} else {
			fmt.Println(needle)
		}
		if hasError {
			os.Exit(1)
		}
	}
}
