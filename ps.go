package main

import "fmt"

func cmdPs() {
	if err := listContainers(); err != nil {
		fmt.Printf("Could not list containers: %v\n", err)
	}
}
