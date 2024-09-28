package config

import (
	"fmt"
	"strings"
	"github.com/spf13/viper"
  "github.com/agoudjiliss/auth-system/data"
)


var Config *datatype.Configuration

func NewConfig() (*datatype.Configuration, error) {
    viper.AddConfigPath("internal/config")  // Path to look for the config file
    viper.SetConfigName("config")           // Config file name (without extension)
    viper.SetConfigType("yaml")             // Config file type
    viper.AutomaticEnv()                    // Automatic environment variable binding
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Replacer for env variables

    err := viper.ReadInConfig() // Read config file
    if err != nil {
        return nil, fmt.Errorf("error loading config file: %s", err)
    }

    Config = &datatype.Configuration{} // Initialize the Config variable

    err = viper.Unmarshal(Config) // Unmarshal config into struct
    if err != nil {
        return nil, fmt.Errorf("error reading config file: %s", err)
    }

    return Config, nil
}
