package config

import (
	"github.com/spf13/viper"
)

// InitDefaultConfig 配置配置项的默认值
func InitDefaultConfig(ConfigViper *viper.Viper) {
	ConfigViper.SetDefault("Address", "0.0.0.0:5000")
}
