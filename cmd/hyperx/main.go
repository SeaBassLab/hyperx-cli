package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/hyperx/packages/cli/cmd/internal/create"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usá: hyperx [dev|start|create] [--port=XXXX]")
		os.Exit(1)
	}

	// Flags compartidas
	portFlag := flag.String("port", "", "Puerto en el que correr la app")

	// Parse solo los flags, sin consumir el primer arg (comando)
	flag.CommandLine.Parse(os.Args[2:])

	switch os.Args[1] {
	case "dev":
		runDev(*portFlag)
	case "start":
		runStart(*portFlag)
	case "create":
		create.Run()
	default:
		fmt.Println("Comando no reconocido. Usá: hyperx [dev|build]")
	}
}
func setDefaultEnv(key, value string) {
	if os.Getenv(key) == "" {
		os.Setenv(key, value)
	}
}

func runDev(port string) {
	setDefaultEnv("HYPERX_ENV", "dev")

	if port != "" {
		os.Setenv("PORT", port)
	} else {
		setDefaultEnv("PORT", "4001") // default si no se pasa
	}

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

func runStart(port string) {
	os.Setenv("HYPERX_ENV", "prod")

	if port != "" {
		os.Setenv("PORT", port)
	} else {
		setDefaultEnv("PORT", "4001") // default si no se pasa
	}

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