package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pborman/uuid"
	"github.com/zencoder/gokay/gkgen"
)

// usage is a string used to provide a user with the application usage
const usage = `usage: gokay <file> [generator-package generator-contructor]
	generator-package        custom package
	generator-contructor     custom generator

examples:
	gokay file.go
	gokay file.go gkcustom NewCustomGKGenerator
`

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, usage)
		return
	}
	log.Println("gokay started. file:", args[0])

	// genPackage := "gkgen"
	// genConstructor := "NewValidator"
	// if len(args) >= 3 {
	// genPackage = args[1]
	// genConstructor = args[2]
	// }

	fileName := args[0]

	fileName, _ = filepath.Abs(fileName)
	fileDir := filepath.Dir(fileName)

	tempName := uuid.NewRandom().String()

	tempDir := fmt.Sprintf("%s/%s", fileDir, tempName)
	if err := os.Mkdir(tempDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating directory %v: %v\n", tempDir, err)
	}
	tempFile := fmt.Sprintf("%s/%s.go", tempDir, tempName)

	outFilePath := fmt.Sprintf("%s_validators.go", strings.TrimSuffix(fileName, filepath.Ext(fileName)))
	// tempOut, err := os.Create(tempFile)
	// if err != nil {
	// 	log.Fatalf("Error while opening %v: %v\n", tempFile, err)
	// }
	// defer tempOut.Close()

	fmt.Println(tempDir)

	// fset := token.NewFileSet() // positions are relative to fset

	// Parse the file given in arguments
	g := gkgen.NewGenerator()
	raw, err := g.GenerateFromFile(fileName)
	if err != nil {
		log.Fatalf("Error while generating validators %v: %v\n", tempFile, err)
	}

	err = ioutil.WriteFile(outFilePath, raw, os.ModePerm)
	if err != nil {
		log.Fatalf("Error while writing to file %v: %v\n", tempFile, err)
	}
	log.Println("gokay finished. file:", args[0])
}
