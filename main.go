package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-yaml/yaml"
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

	readme := generateReadme(action)

	err = ioutil.WriteFile("README.md", []byte(readme), 0644)
	if err != nil {
		log.Fatalf("Error writing README: %v", err)
	}

	fmt.Println("README.md generated successfully!")
}

func generateReadme(action Action) string {
	var builder strings.Builder

	// Title
	builder.WriteString(fmt.Sprintf("# GitHub Action: %s\n\n", action.Name))

	// Description
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
