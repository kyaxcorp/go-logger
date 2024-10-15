package app_config

import (
	"sync"

	"github.com/kyaxcorp/go-helper/_struct"
	"github.com/kyaxcorp/go-logger/config"
)

var cfg Config
var mu sync.RWMutex

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
	mu.RLock()
	defer mu.RUnlock()
	return cfg
}

func SetDefaults(cf *Config) error {
	return _struct.SetDefaultValues(cf)
}

func init() {
	err := SetDefaults(&cfg)
	if err != nil {
		panic(err)
	}
}
