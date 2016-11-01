package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mkideal/cli"
	"github.com/zencoder/gokay/gkgen"
)

// usage is a string used to provide a user with the application usage
const usage = `usage: gokay <file> [--template]
examples:
	gokay file.go
	gokay file.go --template=some.template --template=someOther.template
	gokay file.go --template-dir=some/dir
`

type rootT struct {
	cli.Helper
	CustomTemplates []string `cli:"template" usage:"custom template files"`
	TemplateDirs    []string `cli:"template-dir" usage:"custom template folders"`
}

func main() {
	cli.Run(new(rootT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		args := ctx.Args()
		if len(args) < 1 {
			fmt.Printf("%#v", argv)
			return fmt.Errorf(usage)
		}
		g := gkgen.NewGenerator()

		if len(argv.CustomTemplates) > 0 {
			ctx.String("Adding templates=%s\n", ctx.Color().Grey(argv.CustomTemplates))

			// Add them one by one, because it's really just a warning if we can't load one of their templates.
			for _, templ := range argv.CustomTemplates {
				err := g.AddTemplateFiles(argv.CustomTemplates...)
				if err != nil {
					ctx.String("Unable to add template '%s': %s\n", ctx.Color().Cyan(templ), ctx.Color().Yellow(err))
				}
			}
		}

		for _, dir := range argv.TemplateDirs {
			ctx.String("Scanning templates in %s\n", ctx.Color().Grey(dir))

			// Walk the directory and add each file one by one.
			filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() {
					ctx.String("Adding template %s\n", ctx.Color().Grey(path))
					err = g.AddTemplateFiles(path)
					if err != nil {
						ctx.String("Unable to add template '%s': %s\n", ctx.Color().Cyan(path), ctx.Color().Yellow(err))
					}
				}
				return err
			})
		}

		ctx.String("gokay started. file: %s\n", ctx.Color().Cyan(args[0]))

		fileName := args[0]
		fileName, _ = filepath.Abs(fileName)
		outFilePath := fmt.Sprintf("%s_validators.go", strings.TrimSuffix(fileName, filepath.Ext(fileName)))

		// Parse the file given in arguments
		raw, err := g.GenerateFromFile(fileName)
		if err != nil {
			return fmt.Errorf("Error while generating validators\nInputFile=%s\nError=%s\n", ctx.Color().Cyan(fileName), ctx.Color().RedBg(err))
		}

		err = ioutil.WriteFile(outFilePath, raw, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Error while writing to file %s: %s\n", ctx.Color().Cyan(outFilePath), ctx.Color().Red(err))
		}
		log.Println("gokay finished. file:", args[0])
		return nil
	})
	return
}
