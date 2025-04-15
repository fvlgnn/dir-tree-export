package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fatih/color"
)

// variabili globali
var (
	maxDepth    = flag.Int("depth", 3, "Profondità massima della scansione")
	showHidden  = flag.Bool("show-hidden", false, "Mostra anche i file nascosti")
	outputFile  = flag.String("output", "dir_tree_export_output.txt", "Nome del file di output")
	includeSize = flag.Bool("show-size", false, "Mostra la dimensione dei file")
	onlyDirs    = flag.Bool("only-dirs", false, "Mostra solo directory")
	blacklist   = map[string]bool{
		".git":         true,
		"node_modules": true,
		".vscode":      true,
		".idea":        true,
		".DS_Store":    true,
		"venv":         true,
		".venv":        true,
		"__pycache__":  true,
	}
)

// Funzione principale
func main() {
	flag.Parse()

	// Percorso da analizzare come argomento (default = ".")
	root := "."
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}

	var output strings.Builder
	err := printTree(root, 0, &output, root, "", []bool{})
	if err != nil {
		color.Red("[ERROR]", err)
		os.Exit(1)
	}

	err = os.WriteFile(*outputFile, []byte(output.String()), 0644)
	if err != nil {
		color.Red("[ERROR] Impossibile scrivere il file di output:", err)
		os.Exit(1)
	}

	color.Cyan("[INFO] Esportazione completata in '%s'\n", *outputFile)
}

func printTree(path string, depth int, output *strings.Builder, root string, prefix string, parentLast []bool) error {
	if depth > *maxDepth {
		return nil
	}

	if depth == 0 {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		dirName := filepath.Base(absPath)
		output.WriteString(dirName + "\n")
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	// Filtra gli entry prima del loop
	var filtered []os.DirEntry
	for _, entry := range entries {
		name := entry.Name()
		if blacklist[name] || (!*showHidden && strings.HasPrefix(name, ".")) {
			continue
		}
		if *onlyDirs && !entry.IsDir() {
			continue
		}
		filtered = append(filtered, entry)
	}

	for i, entry := range filtered {
		name := entry.Name()
		isLast := i == len(filtered)-1

		var linePrefix string
		for _, last := range parentLast {
			if last {
				linePrefix += "    "
			} else {
				linePrefix += "│   "
			}
		}

		symbol := "├───"
		if isLast {
			symbol = "└───"
		}

		line := fmt.Sprintf("%s%s%s", linePrefix, symbol, name)
		if *includeSize && !entry.IsDir() {
			info, err := entry.Info()
			if err == nil {
				line += fmt.Sprintf(" (%s)", humanReadableSize(info.Size()))
			}
		}

		output.WriteString(line + "\n")

		if entry.IsDir() {
			nextParentLast := append([]bool{}, parentLast...)
			nextParentLast = append(nextParentLast, isLast)
			printTree(filepath.Join(path, name), depth+1, output, root, prefix+"    ", nextParentLast)
		}
	}

	return nil
}

func humanReadableSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}
