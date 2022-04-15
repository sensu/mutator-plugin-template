package main

import (
	"fmt"
	"log"

	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-plugin-sdk/sensu"
)

// Config represents the mutator plugin config.
type Config struct {
	sensu.PluginConfig
	Example string
}

var (
	mutatorConfig = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "{{ .GithubProject }}",
			Short:    "{{ .Description }}",
			Keyspace: "sensu.io/plugins/{{ .GithubProject }}/config",
		},
	}

	options = []sensu.ConfigOption{
		&sensu.PluginConfigOption[string]{
			Path:      "example",
			Env:       "MUTATOR_EXAMPLE",
			Argument:  "example",
			Shorthand: "e",
			Default:   "",
			Usage:     "An example string configuration option",
			Value:     &mutatorConfig.Example,
		},
	}
)

func main() {
	mutator := sensu.NewGoMutator(&mutatorConfig.PluginConfig, options, checkArgs, executeMutator)
	mutator.Execute()
}

func checkArgs(event *corev2.Event) error {
	if len(mutatorConfig.Example) == 0 {
		return fmt.Errorf("--example or MUTATOR_EXAMPLE environment variable is required")
	}
	return nil
}

func executeMutator(event *corev2.Event) (*corev2.Event, error) {
	log.Println("executing mutator with --example", mutatorConfig.Example)
	return event, nil
}
