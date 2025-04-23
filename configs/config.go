package configs

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var EnvConfig Config

type Config struct {
	Redis Redis `yaml:"redis"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	DB   int    `yaml:"db"`
}

func init() {
	initViper()
}

func initViper() {
	vp := viper.New()

	vp.AddConfigPath("configs/")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	vp.SetEnvKeyReplacer(replacer)

	err := vp.ReadInConfig()
	if err != nil {
		log.Fatalf("read config file failed: %s", err)
	}
	err = vp.Unmarshal(&EnvConfig)
	if err != nil {
		log.Fatalf("unmarshal config file failed: %s", err)
	}

	//fmt.Printf("viper all settings: %v", vp.AllSettings())
}
