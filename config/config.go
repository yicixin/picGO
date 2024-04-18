package config

import "encoding/json"

type (
	ConfigType string

	Config struct {
		Type ConfigType `json:"type"`
	}
)

const (
	ConfigTypeAliOSS ConfigType = "ali"
)

func Load(content []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
