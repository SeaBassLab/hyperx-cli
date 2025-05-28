package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hyperx/packages/cli/cmd/internal/builder"
	"github.com/hyperx/packages/cli/cmd/internal/create"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usá: hyperx [dev|build|start]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "dev":
		runDev()
	case "build":
		builder.RunBuild()
	case "start":
		runStart()
	case "create":
		create.Run()
	default:
		fmt.Println("Comando no reconocido. Usá: hyperx [dev|build]")
	}
}

func runDev() {
	os.Setenv("HYPERX_ENV", "dev")

	cmd := exec.Command("air")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error al ejecutar el servidor en modo desarrollo: %s\n", err)
		os.Exit(1)
	}
}

func runStart() {
	os.Setenv("HYPERX_ENV", "prod")

	cmd := exec.Command("go", "run", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error al ejecutar el servidor: %s\n", err)
		os.Exit(1)
	}
}