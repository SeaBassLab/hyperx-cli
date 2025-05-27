package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hyperx/packages/cli/cmd/internal/builder"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usá: hyperx [dev|build]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "dev":
		runDev()
	case "build":
		builder.RunBuild()
	default:
		fmt.Println("Comando no reconocido. Usá: hyperx [dev|build]")
	}
}

func runDev() {
	cmd := exec.Command("air")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}
