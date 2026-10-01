package main

import (
	"fmt"
	"os"
	"os/exec"
)

func containerName(service string) string {
	return "udock_" + service
}

func sessionName(service string) string {
	return "udock_" + service
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func startService(name string, svc Service) error {
	if svc.Type == "native" {
		return startNative(name, svc)
	}
	return startContainer(name, svc)
}

func startNative(name string, svc Service) error {
	sess := sessionName(name)
	logFile := containerName(name) + ".log"

	_ = exec.Command("tmux", "kill-session", "-t", sess).Run()

	inner := svc.Command + " 2>&1 | tee " + logFile
	tmuxCmd := exec.Command("tmux", "new-session", "-d", "-s", sess, inner)
	tmuxCmd.Stdout = os.Stdout
	tmuxCmd.Stderr = os.Stderr
	if err := tmuxCmd.Run(); err != nil {
		return fmt.Errorf("tmux start failed: %w", err)
	}

	fmt.Printf("→ Started %s (native) in tmux session %s\n", name, sess)
	fmt.Printf("   logs: tail -f %s\n", logFile)
	return nil
}

func startContainer(name string, svc Service) error {
	cn := containerName(name)
	sess := sessionName(name)
	logFile := cn + ".log"

	_ = exec.Command("udocker", "rm", cn).Run()
	_ = exec.Command("tmux", "kill-session", "-t", sess).Run()

	fmt.Printf("→ Pulling %s\n", svc.Image)
	if err := runCommand("udocker", "pull", svc.Image); err != nil {
		return fmt.Errorf("pull failed: %w", err)
	}

	fmt.Printf("→ Creating %s\n", cn)
	if err := runCommand("udocker", "create", "--name="+cn, svc.Image); err != nil {
		return fmt.Errorf("create failed: %w", err)
	}

	udockerArgs := []string{"run"}
	for _, p := range svc.Ports {
		udockerArgs = append(udockerArgs, "--publish="+p)
	}
	for k, v := range svc.Environment {
		udockerArgs = append(udockerArgs, "--env="+k+"="+v)
	}
	for _, v := range svc.Volumes {
		udockerArgs = append(udockerArgs, "--volume="+v)
	}
	udockerArgs = append(udockerArgs, cn)
	if svc.Command != "" {
		udockerArgs = append(udockerArgs, "sh", "-c", svc.Command)
	}

	inner := "udocker " + joinArgs(udockerArgs) + " 2>&1 | tee " + logFile
	tmuxCmd := exec.Command("tmux", "new-session", "-d", "-s", sess, inner)
	tmuxCmd.Stdout = os.Stdout
	tmuxCmd.Stderr = os.Stderr
	if err := tmuxCmd.Run(); err != nil {
		return fmt.Errorf("tmux start failed: %w", err)
	}

	fmt.Printf("→ Started %s in tmux session %s\n", cn, sess)
	fmt.Printf("   logs: tail -f %s\n", logFile)
	return nil
}

func joinArgs(args []string) string {
	out := ""
	for i, a := range args {
		if i > 0 {
			out += " "
		}
		out += a
	}
	return out
}

func stopService(name string) error {
	sess := sessionName(name)
	_ = exec.Command("tmux", "kill-session", "-t", sess).Run()
	_ = exec.Command("udocker", "rm", containerName(name)).Run()
	return nil
}

func showLogs(name string) error {
	return runCommand("tail", "-f", containerName(name)+".log")
}

func listContainers() error {
	return runCommand("udocker", "ps", "-m")
}
