package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"log/slog"
	"os"
	"slices"
	"strings"
	"text/template"
)

//go:embed all:templates
var templates embed.FS

type data struct {
	Package  string              `json:"package"`
	Endpoint string              `json:"endpoint"`
	Methods  map[string][]method `json:"methods"`
}

type method struct {
	Name          string   `json:"name"`
	Many          bool     `json:"many"`
	Signature     []string `json:"signature"`
	ReturnType    string   `json:"return_type"`
	HasOptions    bool     `json:"has_options"`
	Endpoint      string   `json:"-"`
	URLParts      []string `json:"url_parts"`
	Documentation string   `json:"documentation"`
}

func main() {
	os.Exit(executeTemplate())
}

func executeTemplate() int {
	output := flag.String("o", "-", "Path to write output to")
	flag.Parse()

	inputFileName := flag.Arg(0)
	if inputFileName == "" {
		slog.Error("no file name was given")
		return 1
	}

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		slog.Error("reading input file", "err", err)
		return 1
	}

	defer inputFile.Close()

	var data data
	if err := json.NewDecoder(inputFile).Decode(&data); err != nil {
		slog.Error("decoding file data", "err", err)
		return 1
	}

	var buf bytes.Buffer
	tmplHeading, err := template.New("_heading.tmpl").ParseFS(templates, "templates/_heading.tmpl")
	if err != nil {
		slog.Error("parsing heading template", "err", err)
		return 1
	}

	if err := tmplHeading.ExecuteTemplate(&buf, "heading", data); err != nil {
		slog.Error("executing heading template", "err", err)
		return 1
	}

	funcMap := template.FuncMap{
		"join": strings.Join,
	}

	for method, values := range data.Methods {
		tmplMain, err := template.New("main").Funcs(funcMap).ParseFS(templates, "templates/main.tmpl")
		if err != nil {
			slog.Error("parsing main template", "err", err)
			return 1
		}

		method = strings.ToLower(method)
		methodTemplateName := fmt.Sprintf("templates/_method_%s.tmpl", method)
		tmplMethod, err := tmplMain.ParseFS(templates, methodTemplateName)
		if err != nil {
			slog.Error("parsing method template", "err", err)
			return 1
		}

		for methodData := range slices.Values(values) {
			methodData.Endpoint = data.Endpoint
			if err := tmplMethod.ExecuteTemplate(&buf, "main", methodData); err != nil {
				slog.Error("executing template for method", "method", method, "err", err)
				return 1
			}
		}
	}

	formattedBytes, err := format.Source(buf.Bytes())
	if err != nil {
		slog.Error("formatting output", "err", err)
		slog.Info("dumping data")
		fmt.Println(buf.String())
	}

	outputFile := os.Stdout
	if *output != "-" {
		outputFile, err = os.Create(*output)
		if err != nil {
			slog.Error("creating output file", "err", err)
			return 1
		}

		defer outputFile.Close()
	}

	if n, err := outputFile.Write(formattedBytes); err != nil {
		fmt.Println(n)
		slog.Error("writing template to output", "err", err)
		return 1
	}

	return 0
}
