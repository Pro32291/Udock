package main

import (
	"fmt"
	"log"
)

func cmdDown() {
	path, err := findConfigFile()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := loadConfig(path)
	if err != nil {
		log.Fatalf("Could not read %s: %v", path, err)
	}
	for name := range cfg.Services {
		fmt.Printf("→ Stopping %s\n", name)
		if err := stopService(name); err != nil {
			fmt.Printf("  warning: %v\n", err)
		}
	}
	fmt.Println("\n✓ All services stopped.")
}
