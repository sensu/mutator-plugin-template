package main

import (
	"fmt"
	"log"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/api/core/v2"
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

	options = []*sensu.PluginConfigOption{
		&sensu.PluginConfigOption{
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

func checkArgs(_ *v2.Event) error {
	if len(mutatorConfig.Example) == 0 {
		return fmt.Errorf("--example or MUTATOR_EXAMPLE environment variable is required")
	}
	return nil
}

func executeMutator(event *v2.Event) (*v2.Event, error) {
	log.Println("executing mutator with --example", mutatorConfig.Example)
	return &v2.Event{}, nil
}
