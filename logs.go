package main

import (
	"fmt"
	"os"
)

func cmdLogs(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: udock logs <service-name>")
		os.Exit(1)
	}
	if err := showLogs(args[0]); err != nil {
		fmt.Printf("Could not get logs: %v\n", err)
	}
}
