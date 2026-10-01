package main

import (
	"fmt"
	"os"
)

const configFile = "udock.yaml"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "up":
		cmdUp()
	case "down":
		cmdDown()
	case "clean":
		cmdClean()
	case "install":
		cmdInstall()
	case "logs":
		cmdLogs(os.Args[2:])
	case "ps":
		cmdPs()
	case "status":
		cmdStatus()
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println(`udock — one config, one command, your whole dev stack.

Usage:
  udock up           Start all services in udock.yaml
  udock down         Stop and remove all services
  udock clean        Force-kill everything (use when stuck)
  udock install      Install udock to $PREFIX/bin (global)
  udock status       Show status of all services
  udock logs <name>  Stream logs from a service
  udock ps           Show raw udocker container info
  udock help         Show this message`)
}
