package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type Action struct {
	Name        string           `yaml:"name"`
	Description string           `yaml:"description"`
	Author      string           `yaml:"author"`
	Inputs      map[string]Input `yaml:"inputs"`
}

type Input struct {
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default"`
}

type TitleLayout struct {
	Title TitleConfig `hcl:"title,block"`
}

type TitleConfig struct {
	Enabled  bool   `hcl:"enabled"`
	Override string `hcl:"override"`
}

func main() {
	data, err := ioutil.ReadFile("action.yml")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var action Action
	err = yaml.Unmarshal(data, &action)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	hclData, err := ioutil.ReadFile("readme.hcl")
	if err != nil {
		log.Fatalf("Error reading HCL configuration: %v", err)
	}

	layout, err := parseTitleLayout(hclData)
	if err != nil {
		log.Fatalf("Error parsing layout configuration: %v", err)
	}

	readme := generateReadme(action, layout)

	err = ioutil.WriteFile("README.md", []byte(readme), 0644)
	if err != nil {
		log.Fatalf("Error writing README: %v", err)
	}

	fmt.Println("README.md generated successfully!")
}

func parseTitleLayout(hclData []byte) (*TitleLayout, error) {
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL(hclData, "readme.hcl")
	if diags.HasErrors() {
		return nil, diags.Errs()[0]
	}

	var layout TitleLayout
	diags = gohcl.DecodeBody(file.Body, nil, &layout)
	if diags.HasErrors() {
		return nil, diags.Errs()[0]
	}

	return &layout, nil
}

func generateReadme(action Action, layout *TitleLayout) string {
	var builder strings.Builder

	if layout.Title.Enabled {
		titleText := action.Name
		if layout.Title.Override != "" {
			titleText = layout.Title.Override
		}
		builder.WriteString(fmt.Sprintf("# GitHub Action: %s\n\n", titleText))
	}

	builder.WriteString(action.Description + "\n\n")

	// Inputs section
	builder.WriteString("## Inputs\n\n")
	builder.WriteString("| Name | Description | Required | Default Value |\n")
	builder.WriteString("|------|-------------|----------|---------------|\n")

	for name, input := range action.Inputs {
		defaultValue := "`Null`"
		if input.Default != "" {
			defaultValue = fmt.Sprintf("`%s`", input.Default)
		}
		required := "`false`"
		if input.Required {
			required = "`true`"
		}
		builder.WriteString(fmt.Sprintf("| `%s` | %s | %s | %s |\n", name, input.Description, required, defaultValue))
	}

	return builder.String()
}
