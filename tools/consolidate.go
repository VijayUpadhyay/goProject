package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Directories to skip
var skipDirs = map[string]bool{
	"tools":     true,
	".git":      true,
	".github":   true,
	"_projects": true, // Keep projects as-is
}

func main() {
	rootDir := ".."
	if len(os.Args) > 1 {
		rootDir = os.Args[1]
	}

	fmt.Println("Starting consolidation of Go project...")
	fmt.Println("Root directory:", rootDir)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip if not a directory
		if !info.IsDir() {
			return nil
		}

		// Skip root and special directories
		relPath, _ := filepath.Rel(rootDir, path)
		if relPath == "." || relPath == ".." {
			return nil
		}

		dirName := filepath.Base(path)
		if skipDirs[dirName] || strings.HasPrefix(dirName, ".") {
			return filepath.SkipDir
		}

		// Check if directory has Go files
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}

		var goFiles []string
		for _, f := range files {
			if !f.IsDir() && strings.HasSuffix(f.Name(), ".go") && f.Name() != "main.go" {
				goFiles = append(goFiles, f.Name())
			}
		}

		// If directory has multiple Go files, consolidate them
		if len(goFiles) > 1 {
			fmt.Printf("\nConsolidating %s (%d files)...\n", relPath, len(goFiles))
			err := consolidateDirectory(path, goFiles)
			if err != nil {
				fmt.Printf("  Error: %v\n", err)
			} else {
				fmt.Printf("  ✓ Created main.go\n")
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✓ Consolidation complete!")
}

func consolidateDirectory(dirPath string, goFiles []string) error {
	var mainContent strings.Builder
	mainContent.WriteString("package main\n\n")

	// Collect all imports
	imports := make(map[string]bool)
	var examples []string

	// Parse each file
	for i, filename := range goFiles {
		filePath := filepath.Join(dirPath, filename)
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("reading %s: %w", filename, err)
		}

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, filePath, content, parser.ParseComments)
		if err != nil {
			// If parsing fails, skip this file
			fmt.Printf("  Warning: Could not parse %s, skipping\n", filename)
			continue
		}

		// Collect imports
		for _, imp := range f.Imports {
			imports[imp.Path.Value] = true
		}

		// Extract non-main functions and types
		exampleName := strings.TrimSuffix(filename, ".go")
		exampleName = strings.ReplaceAll(exampleName, "-", "_")
		exampleName = strings.ReplaceAll(exampleName, " ", "_")

		var exampleCode strings.Builder
		exampleCode.WriteString(fmt.Sprintf("\n// ============================================================================\n"))
		exampleCode.WriteString(fmt.Sprintf("// Example %d: %s\n", i+1, filename))
		exampleCode.WriteString(fmt.Sprintf("// ============================================================================\n\n"))

		// Extract declarations (types, functions, etc.)
		for _, decl := range f.Decls {
			switch d := decl.(type) {
			case *ast.FuncDecl:
				if d.Name.Name == "main" {
					// Rename main to example function
					funcName := fmt.Sprintf("example%d_%s", i+1, exampleName)
					exampleCode.WriteString(fmt.Sprintf("func %s() {\n", funcName))

					// Extract main body
					if d.Body != nil {
						start := fset.Position(d.Body.Lbrace).Offset
						end := fset.Position(d.Body.Rbrace).Offset
						body := string(content[start+1 : end])
						exampleCode.WriteString(body)
					}
					exampleCode.WriteString("}\n\n")
					examples = append(examples, funcName)
				} else {
					// Keep other functions as-is
					start := fset.Position(d.Pos()).Offset
					end := fset.Position(d.End()).Offset
					exampleCode.WriteString(string(content[start:end]))
					exampleCode.WriteString("\n\n")
				}
			case *ast.GenDecl:
				// Keep type, const, var declarations
				start := fset.Position(d.Pos()).Offset
				end := fset.Position(d.End()).Offset
				exampleCode.WriteString(string(content[start:end]))
				exampleCode.WriteString("\n\n")
			}
		}

		mainContent.WriteString(exampleCode.String())
	}

	// Write imports
	if len(imports) > 0 {
		var importBlock strings.Builder
		importBlock.WriteString("import (\n")
		for imp := range imports {
			importBlock.WriteString(fmt.Sprintf("\t%s\n", imp))
		}
		importBlock.WriteString(")\n\n")

		// Insert imports after package declaration
		parts := strings.SplitN(mainContent.String(), "\n\n", 2)
		mainContent.Reset()
		mainContent.WriteString(parts[0])
		mainContent.WriteString("\n\n")
		mainContent.WriteString(importBlock.String())
		if len(parts) > 1 {
			mainContent.WriteString(parts[1])
		}
	}

	// Create main function that calls all examples
	mainContent.WriteString("// ============================================================================\n")
	mainContent.WriteString("// Main Function - Run All Examples\n")
	mainContent.WriteString("// ============================================================================\n\n")
	mainContent.WriteString("func main() {\n")
	mainContent.WriteString(fmt.Sprintf("\tfmt.Println(\"Running %d examples from %s\")\n", len(examples), filepath.Base(dirPath)))
	mainContent.WriteString("\tfmt.Println(\"=\", strings.Repeat(\"=\", 50))\n\n")

	for _, exampleFunc := range examples {
		mainContent.WriteString(fmt.Sprintf("\tfmt.Println(\"\\n--- %s ---\")\n", exampleFunc))
		mainContent.WriteString(fmt.Sprintf("\t%s()\n\n", exampleFunc))
	}

	mainContent.WriteString("\tfmt.Println(\"\\n\" + strings.Repeat(\"=\", 50))\n")
	mainContent.WriteString("\tfmt.Println(\"All examples completed!\")\n")
	mainContent.WriteString("}\n")

	// Write consolidated main.go
	mainPath := filepath.Join(dirPath, "main.go")
	err := ioutil.WriteFile(mainPath, []byte(mainContent.String()), 0644)
	if err != nil {
		return fmt.Errorf("writing main.go: %w", err)
	}

	// Delete old files
	for _, filename := range goFiles {
		oldPath := filepath.Join(dirPath, filename)
		if err := os.Remove(oldPath); err != nil {
			fmt.Printf("  Warning: Could not delete %s: %v\n", filename, err)
		}
	}

	return nil
}

// Made with Bob
