package main

import (
	"fmt"
	"log"
)

func cmdUp() {
	path, err := findConfigFile()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := loadConfig(path)
	if err != nil {
		log.Fatalf("Could not read %s: %v", path, err)
	}
	if len(cfg.Services) == 0 {
		log.Fatal("No services defined in udock.yaml")
	}
	for name, svc := range cfg.Services {
		fmt.Printf("\n[%s]\n", name)
		if err := startService(name, svc); err != nil {
			log.Fatalf("Failed to start %s: %v", name, err)
		}
	}
	fmt.Println("\n✓ All services are up.")
}
