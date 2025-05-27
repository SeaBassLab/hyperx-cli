package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func RunBuild() {
	fmt.Println("🔧 Iniciando build de HyperX...")

	err := os.RemoveAll("dist")
	checkErr(err)
	err = os.MkdirAll("dist", 0755)
	checkErr(err)

	layout := mustParseTemplate("_layout.html")

	// Registrar partials
	err = filepath.Walk("partials", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}
		content, err := os.ReadFile(path)
		checkErr(err)
		name := strings.TrimSuffix(filepath.Base(path), ".html")
		_, err = layout.New(name).Parse(string(content))
		checkErr(err)
		return nil
	})
	checkErr(err)

	// Iterar páginas
	err = filepath.Walk("pages", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}

		pageContent, err := os.ReadFile(path)
		checkErr(err)

		// Clonar layout + parsear el bloque de contenido
		pageTpl, err := layout.Clone()
		checkErr(err)

		_, err = pageTpl.New("content").Parse(string(pageContent))
		checkErr(err)

		// Resolver path destino
		rel, _ := filepath.Rel("pages", path)
		outPath := filepath.Join("dist", strings.ReplaceAll(rel, "[slug]", "ejemplo"))

		err = os.MkdirAll(filepath.Dir(outPath), 0755)
		checkErr(err)

		outFile, err := os.Create(outPath)
		checkErr(err)
		defer outFile.Close()

		err = pageTpl.Execute(outFile, nil)
		checkErr(err)

		fmt.Printf("✅ %s -> %s\n", path, outPath)
		return nil
	})
	checkErr(err)

	fmt.Println("🎉 Build finalizado en ./dist")
}

func mustParseTemplate(path string) *template.Template {
	tpl, err := template.ParseFiles(path)
	checkErr(err)
	return tpl
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}
}
