package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] != "dev" {
		fmt.Println("Usá: hyperx dev")
		os.Exit(1)
	}

	cmd := exec.Command("air")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}
