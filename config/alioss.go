package config

import "encoding/json"

type AliOSSConfig struct {
	AccessKeyID     string `json:"accessKeyID"`
	AccessKeySecret string `json:"accessKeySecret"`
	Endpoint        string `json:"endpoint"`
	Bucket          string `json:"bucket"`

	Dir         string `json:"dir"`
	PlaceHolder string `json:"placeHolder"`
}

func LoadAliOSSConfig(content []byte) (*AliOSSConfig, error) {
	var cfg AliOSSConfig
	err := json.Unmarshal(content, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
