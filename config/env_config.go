package config

import (
	"github.com/spf13/viper"
)

// InitEnvConfig 从环境变量获得的配置内容初始化配置
func InitEnvConfig(ConfigViper *viper.Viper) {
	EnvConfigViper := viper.New()
	EnvConfigViper.SetEnvPrefix("jwtsigner") // will be uppercased automatically
	EnvConfigViper.BindEnv("address")

	if EnvConfigViper.Get("address") != nil {
		ConfigViper.Set("Address", EnvConfigViper.Get("address"))
	}
}
