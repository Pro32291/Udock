package main

import (
	"fmt"
	"os/exec"
)

func cmdClean() {
	fmt.Println("→ Stopping all tmux sessions...")
	_ = exec.Command("tmux", "kill-server").Run()

	fmt.Println("→ Killing stray processes...")
	_ = exec.Command("pkill", "-f", "udocker").Run()
	_ = exec.Command("pkill", "-f", "nginx").Run()
	_ = exec.Command("pkill", "-f", "http.server").Run()

	fmt.Println("→ Removing leftover udocker containers...")
	_ = exec.Command("udocker", "rm", "udock_web").Run()
	_ = exec.Command("udocker", "rm", "udock_files").Run()

	fmt.Println("\n✓ Clean.")
}
