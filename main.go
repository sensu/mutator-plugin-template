package main

import (
	"fmt"
	"log"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
)

type MutatorConfig struct {
	sensu.PluginConfig
	Example string
}

type ConfigOptions struct {
	Example sensu.PluginConfigOption
}

var (
	mutatorConfig = MutatorConfig{
		PluginConfig: sensu.PluginConfig{
			Name:     "{{ .GithubProject }}",
			Short:    "{{ .Description }}",
			Timeout:  10,
			Keyspace: "sensu.io/plugins/{{ .GithubProject }}/config",
		},
	}

	mutatorConfigOptions = ConfigOptions{
		Example: sensu.PluginConfigOption{
			Path:      "example",
			Env:       "MUTATOR_EXAMPLE",
			Argument:  "example",
			Shorthand: "e",
			Default:   "",
			Usage:     "An example configuration option",
			Value:     &mutatorConfig.Example,
		},
	}

	options = []*sensu.PluginConfigOption{
		&mutatorConfigOptions.Example,
	}
)

func main() {
	mutator := sensu.NewGoMutator(&mutatorConfig.PluginConfig, options, checkArgs, executeMutator)
	mutator.Execute()
}

func checkArgs(_ *types.Event) error {
	if len(mutatorConfig.Example) == 0 {
		return fmt.Errorf("--example or MUTATOR_EXAMPLE environment variable is required")
	}
	return nil
}

func executeMutator(event *types.Event) (*types.Event, error) {
	log.Println("executing mutator with --example", mutatorConfig.Example)
	return &types.Event{}, nil
}
