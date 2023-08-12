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

type Example struct {
	Enabled bool   `hcl:"enabled"`
	Name    string `hcl:"name"`
	Path    string `hcl:"path"`
}

type ExamplesConfig struct {
	Examples []Example `hcl:"example,block"`
}

type FooterConfig struct {
	From string `hcl:"footer_from"`
}

type QuickstartConfig struct {
	Path string `hcl:"path"`
}

type TitleConfig struct {
	Enabled  bool   `hcl:"enabled"`
	Override string `hcl:"override"`
}

type LayoutConfig struct {
	Version    string           `hcl:"version"`
	Title      TitleConfig      `hcl:"title,block"`
	Quickstart QuickstartConfig `hcl:"quickstart,block"`
	Examples   ExamplesConfig   `hcl:"examples,block"`
	Footer     FooterConfig     `hcl:"footer,block"`
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

	layout, err := parseLayoutConfig(hclData)
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

func parseLayoutConfig(hclData []byte) (*LayoutConfig, error) {
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL(hclData, "readme.hcl")
	if diags.HasErrors() {
		return nil, diags.Errs()[0]
	}

	var layout LayoutConfig
	diags = gohcl.DecodeBody(file.Body, nil, &layout)
	if diags.HasErrors() {
		return nil, diags.Errs()[0]
	}

	return &layout, nil
}

func generateReadme(action Action, layout *LayoutConfig) string {
	var builder strings.Builder

	// Title section
	if layout.Title.Enabled {
		titleText := action.Name
		if layout.Title.Override != "" {
			titleText = layout.Title.Override
		}
		builder.WriteString(fmt.Sprintf("# GitHub Action: %s\n\n", titleText))
	}

	// Quickstart section
	if layout.Quickstart.Path != "" {
		builder.WriteString("## Quickstart\n\n")
		quickstartContent, err := ioutil.ReadFile(layout.Quickstart.Path)
		if err != nil {
			log.Printf("Error reading quickstart file: %v\n", err)
		} else {
			quickstartContent = replaceVersionPlaceholder(quickstartContent, layout.Version)
			builder.Write(quickstartContent)
		}
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

	// Examples section
	builder.WriteString("\n## Examples\n\n")
	for _, example := range layout.Examples.Examples {
		if example.Enabled {
			builder.WriteString(fmt.Sprintf("### %s\n\n", example.Name))
			exampleContent, err := ioutil.ReadFile(example.Path)
			if err != nil {
				log.Printf("Error reading example file: %v\n", err)
			} else {
				exampleContent = replaceVersionPlaceholder(exampleContent, layout.Version)
				builder.Write(exampleContent)
			}
		}
	}

	// Footer section
	if layout.Footer.From != "" {
		footerContent, err := ioutil.ReadFile(layout.Footer.From)
		if err != nil {
			log.Printf("Error reading footer file: %v\n", err)
		} else {
			builder.WriteString("\n---\n")
			builder.Write(footerContent)
		}
	}

	return builder.String()
}

func replaceVersionPlaceholder(content []byte, version string) []byte {
	placeholder := "${{version}}"
	return []byte(strings.ReplaceAll(string(content), placeholder, version))
}
