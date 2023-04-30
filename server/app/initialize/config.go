package initialize

import (
	"log"

	"github.com/noobHuKai/app/g"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %s", err)
	}

	err = viper.Unmarshal(&g.Cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	g.JWTSecret = []byte(g.Cfg.Server.Secret)
}
