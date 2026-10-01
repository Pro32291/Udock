package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func cmdInstall() {
	prefix := os.Getenv("PREFIX")
	if prefix == "" {
		fmt.Println("PREFIX not set. Are you in Termux?")
		os.Exit(1)
	}
	dest := filepath.Join(prefix, "bin", "udock")

	// Build a fresh binary.
	fmt.Println("→ Building...")
	build := exec.Command("go", "build", "-o", "udock", ".")
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		fmt.Printf("Build failed: %v\n", err)
		os.Exit(1)
	}

	// Copy to $PREFIX/bin.
	fmt.Printf("→ Installing to %s\n", dest)
	data, err := os.ReadFile("udock")
	if err != nil {
		fmt.Printf("Read failed: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(dest, data, 0755); err != nil {
		fmt.Printf("Write failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✓ Installed. Run 'udock help' from anywhere.")
}
