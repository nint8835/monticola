package config

import (
	"fmt"
)

type AgentConfig struct {
	SharedConfig

	// ListenAddress is the address to listen on
	ListenAddress string `split_words:"true" default:":8081"`
}

func LoadAgentConfig() (*AgentConfig, error) {
	config, err := loadConfig[AgentConfig]()
	if err != nil {
		return nil, err
	}

	err = initLoggingConfig(&config.SharedConfig)
	if err != nil {
		return nil, fmt.Errorf("error initializing logging: %w", err)
	}

	return config, nil
}
