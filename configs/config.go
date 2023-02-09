package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type config struct {
	BotToken string
}

var Config config

func init() {

	err := searchConfig()
	if err != nil {
		fmt.Println("Config Viper error: ", err)
		return
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Printf("unable to decode into struct, %v\n", err)
	}
}

func searchConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	return viper.ReadInConfig()
}

func GetToken() string {
	return Config.BotToken
}
