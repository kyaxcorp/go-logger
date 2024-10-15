package app_config

import "github.com/kyaxcorp/go-logger/config"

var cfg Config

type Config struct {
	// AppLogLevel -> it's for the app itself... it's the main (root) LOGGER!
	AppLogLevel int `yaml:"level" mapstructure:"level" default:"4"`

	LogsPath string `yaml:"logs_path" mapstructure:"logs_path" default:"logs"`
	// This is the default channel
	DefaultChannel string `yaml:"default_channel" mapstructure:"default_channel" default:"default"`
	// A default channel will always be and will be created automatically
	Channels map[string]config.Config
}

func GetConfig() Config {
	return cfg
}
