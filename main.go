package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-yaml/yaml"
)

type Input struct {
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default"`
}

type Action struct {
	Name   string           `yaml:"name"`
	Desc   string           `yaml:"description"`
	Author string           `yaml:"author"`
	Inputs map[string]Input `yaml:"inputs"`
}

func main() {
	data, err := ioutil.ReadFile("action.yml")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var action Action
	err = yaml.Unmarshal(data, &action)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return
	}

	readme := generateReadme(action)
	fmt.Println(readme)
}

func generateReadme(action Action) string {
	readme := fmt.Sprintf("# GitHub action: %s\n\n%s\n\n", action.Name, action.Desc)

	readme += "## Inputs\n\n"
	readme += "| Input name | Required | Default Value |\n"
	readme += "|------------|----------|---------------|\n"

	for inputName, input := range action.Inputs {
		required := "`false`"
		if input.Required {
			required = "`true`"
		}

		defaultValue := input.Default
		if defaultValue == "" {
			defaultValue = "`N/A`"
		} else {
			defaultValue = fmt.Sprintf("`%s`", defaultValue)
		}

		readme += fmt.Sprintf("| `%s` | %s | %s |\n", strings.Title(inputName), required, defaultValue)
	}

	readme += "\n## Author\n\n"
	readme += action.Author

	return readme
}
