package cmd

import (
	"os"
	"os/exec"
)

func Dev() {
	cmd := exec.Command("go", "run", "app/playground/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
