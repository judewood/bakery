package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/knadh/koanf/parsers/toml/v2"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// osEnvFunction allows us to mock os.Getenv so we can test that New
// panics when the config file does not exist
var osEnvFunction = os.Getenv

// Config wraps around a config provider to decouple from the main code
type Config struct {
	provider *koanf.Koanf
}

// New instantiates and configures the config provider
// We are using koanf because it uses fewer dependencies than viper
func New(folder string) *Config {
	k := koanf.New(".")
	var env string
	if env = osEnvFunction("ENVIRONMENT"); env == "" {
		env = "local"
	}
	fmt.Printf("\nCurrent environment is %v \n", env)
	configFile := filepath.Join(folder, fmt.Sprintf("%v.toml", env))
	err := k.Load(file.Provider(configFile), toml.Parser())
	if err != nil {
		log.Panicf("Unable to load config file: %s\n %v", configFile, err)
	}
	return &Config{provider: k}
}

// GetStringSetting returns value of setting key as a string
func (c Config) GetStringSetting(key string) string {
	return c.provider.String(key)
}
