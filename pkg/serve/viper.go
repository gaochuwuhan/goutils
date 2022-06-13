package serve

import (
	"github.com/spf13/viper"
	"log"
)

func Viper() *viper.Viper {
	v:=viper.New()
	v.SetConfigName("conf") //TODO 后续改为从环境变量中获取
	v.SetConfigType("toml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)

	}
	return v
}
