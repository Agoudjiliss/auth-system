package config

import (
	"fmt"
	"strings"
	"github.com/spf13/viper"
  "github.com/agoudjiliss/auth-system/data"
)


var Config *datatype.Configuration

func NewConfig()  (*datatype.Configuration,error){
  viper.AddConfigPath("internal/config")
  viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  viper.AutomaticEnv()
  viper.SetEnvKeyReplacer(strings.NewReplacer(`.`,`_`))

  err :=viper.ReadInConfig() 
  if err != nil {
	  	return nil, fmt.Errorf("error loading config file: %s", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}

	return Config, nil

}
