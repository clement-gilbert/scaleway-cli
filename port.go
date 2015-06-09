package main

import log "github.com/Sirupsen/logrus"

var cmdPort = &Command{
	Exec:        runPort,
	UsageLine:   "port [OPTIONS] SERVER [PRIVATE_PORT[/PROTO]]",
	Description: "Lookup the public-facing port that is NAT-ed to PRIVATE_PORT",
	Help:        "List port mappings for the SERVER, or lookup the public-facing port that is NAT-ed to the PRIVATE_PORT",
}

func init() {
	cmdPort.Flag.BoolVar(&portHelp, []string{"h", "-help"}, false, "Print usage")
}

// FLags
var portHelp bool // -h, --help flag

func runPort(cmd *Command, args []string) {
	if portHelp {
		cmd.PrintUsage()
	}
	if len(args) < 1 {
		cmd.PrintShortUsage()
	}

	serverID := cmd.API.GetServerID(args[0])
	server, err := cmd.API.GetServer(serverID)
	if err != nil {
		log.Fatalf("Failed to get server information for %s: %v", serverID, err)
	}

	command := []string{"netstat -lutn 2>/dev/null | grep LISTEN"}
	err = serverExec(server, command, true)
	if err != nil {
		log.Fatalf("Command execution failed: %v", err)
	}
}
