package main

import (
	"bytes"
	"flag"
	"go/format"
	"html/template"
	"log/slog"
	"os"
	"strings"

	_ "embed"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed type.tmpl
var templateData string

type data struct {
	Name  string
	Title string
}

func main() {
	os.Exit(executeTemplate())
}

func executeTemplate() int {
	rawTypes := flag.String("t", "", "List of types to generate a file for")
	output := flag.String("o", "-", "Destination output to write template to")
	flag.Parse()

	if *rawTypes == "" {
		slog.Error("a list of types must be provided")
		return 1
	}

	titleCaser := cases.Title(language.AmericanEnglish)
	var types []data

	for typeName := range strings.SplitSeq(*rawTypes, ",") {
		types = append(types, data{
			Name:  typeName,
			Title: titleCaser.String(typeName),
		})
	}

	tmpl, err := template.New("types").Parse(templateData)
	if err != nil {
		slog.Error("parsing template", "error", err)
		return 1
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "types", types); err != nil {
		slog.Error("executing template", "error", err)
		return 1
	}

	formattedBytes, err := format.Source(buf.Bytes())
	if err != nil {
		slog.Error("formatting template", "error", err)
		return 1
	}

	var file *os.File
	switch *output {
	case "-":
		file = os.Stdout
	default:
		file, err = os.Create(*output)
		if err != nil {
			slog.Error("creating output file", "error", err)
			return 1
		}

		defer file.Close()
	}

	if _, err := file.Write(formattedBytes); err != nil {
		slog.Error("writing bytes to file", "error", err)
		return 1
	}

	return 0
}
