package config

import (
	"fmt"
	"os"

	"github.com/knadh/koanf/parsers/toml/v2"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Config wraps around a config provider to decouple from the main code
type Config struct {
	provider *koanf.Koanf
}

// New instantiates and configures the config provider
// We are using koanf because it uses fewer dependencies than viper
func New() *Config {
	k := koanf.New(".")
	var env string
	if env = os.Getenv("ENVIRONMENT"); env == "" {
		env = "local"
	}
	fmt.Printf("\nCurrent environment is %v", env)
	configFile := fmt.Sprintf("./environments/%v.toml", env)
	k.Load(file.Provider(configFile), toml.Parser())
	return &Config{provider: k}
}

// GetStringSetting returns value of setting key as a string
func (c Config) GetStringSetting(key string) string {
	return c.provider.String(key)
}
