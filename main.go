package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: hyperx [dev|build|start]")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "dev":
		fmt.Println("Starting HyperX in dev mode...")
		runDev()
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
	}
}

func runDev() {
	cmd := exec.Command("go", "run", "./apps/playground/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
