package create

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const templateRepoURL = "https://github.com/SeaBassLab/hyperx-template.git"

func Run() {
	printLogo()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("📦 Nombre del proyecto: ")
	projectName, _ := reader.ReadString('\n')
	projectName = strings.TrimSpace(projectName)

	if projectName == "" {
		fmt.Println("❌ El nombre del proyecto no puede estar vacío.")
		os.Exit(1)
	}

	// Clonar el template desde GitHub
	fmt.Println("🔄 Clonando template...")
	runCommand("git", []string{"clone", "--depth", "1", templateRepoURL, projectName}, "")

	// Borrar .git para que no herede el repo
	err := os.RemoveAll(filepath.Join(projectName, ".git"))
	check(err, "Error eliminando .git")

	// Reemplazar el nombre del módulo
	err = replaceModuleName(filepath.Join(projectName, "go.mod"), projectName)
	check(err, "Error modificando go.mod")

	// Mensaje final
	fmt.Println("✅ Proyecto creado con éxito!")
	fmt.Printf("👉 Comenzá así:\n\n  cd %s\n  go mod tidy\n  hyperx dev\n\n", projectName)
}

func replaceModuleName(goModPath, projectName string) error {
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 {
		return fmt.Errorf("archivo go.mod vacío")
	}
	lines[0] = "module " + projectName
	return os.WriteFile(goModPath, []byte(strings.Join(lines, "\n")), 0644)
}

func runCommand(name string, args []string, workDir string) {
	cmd := exec.Command(name, args...)
	if workDir != "" {
		cmd.Dir = workDir
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err, fmt.Sprintf("Error ejecutando %s %v", name, args))
}

func check(err error, msg string) {
	if err != nil {
		fmt.Printf("❌ %s: %v\n", msg, err)
		os.Exit(1)
	}
}

func printLogo() {
	fmt.Println(`
╔══════════════════════════════════════╗
║        🚀 HyperX Project Init        ║
╚══════════════════════════════════════╝
`)
}
