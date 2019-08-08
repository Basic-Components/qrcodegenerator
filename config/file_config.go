package config

import (
	"github.com/Basic-Components/qrcodegenerator/logger"

	"github.com/spf13/viper"
)

// SetFileConfig 从指定的配置文件中读取配置
func SetFileConfig(ConfigViper *viper.Viper, fileName string, filePaths []string) {
	FileConfigViper := viper.New()
	FileConfigViper.SetConfigName(fileName)
	for _, filePath := range filePaths {
		FileConfigViper.AddConfigPath(filePath)
	}
	err := FileConfigViper.ReadInConfig() // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		logger.Logger.Info("config file not found: %s \n", err)
	} else {
		ConfigViper.Set("Address", FileConfigViper.Get("Address"))
	}
}

// InitFileConfig 从默认的配置文件位置读取配置
func InitFileConfig(ConfigViper *viper.Viper) {
	fileName := "config"
	filePaths := []string{"/etc/jwt-signer/", "$HOME/.jwt-signer", "."}
	SetFileConfig(ConfigViper, fileName, filePaths)
}
