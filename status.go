package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func cmdStatus() {
	path, err := findConfigFile()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := loadConfig(path)
	if err != nil {
		log.Fatalf("Could not read %s: %v", path, err)
	}

	fmt.Println("SERVICE        TYPE       STATUS")
	fmt.Println("─────────────────────────────────────")

	for name, svc := range cfg.Services {
		serviceType := "container"
		if svc.Type == "native" {
			serviceType = "native"
		}
		status := checkStatus(name, svc)
		fmt.Printf("%-14s %-10s %s\n", name, serviceType, status)
	}
}

func checkStatus(name string, svc Service) string {
	sess := sessionName(name)

	tmuxOut, _ := exec.Command("tmux", "ls").Output()
	if !strings.Contains(string(tmuxOut), sess+":") {
		return "✗ not running"
	}

	if svc.Type == "native" {
		return "● running"
	}

	psOut, _ := exec.Command("udocker", "ps", "-m").Output()
	if strings.Contains(string(psOut), containerName(name)) {
		return "● running"
	}
	return "○ tmux up, container gone"
}
