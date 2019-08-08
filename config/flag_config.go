package config

import (
	"path"
	"strings"

	"github.com/Basic-Components/qrcodegenerator/logger"

	logrus "github.com/sirupsen/logrus"
	"github.com/small-tk/pathlib"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// InitFlagConfig 初始化命令行传入的参数到配置,返回值为false表示要执行创建秘钥否则为启动服务
func InitFlagConfig(ConfigViper *viper.Viper) bool {
	loglevel := pflag.StringP("loglevel", "l", "WARN", "创建rsa公私钥对")
	confPath := pflag.StringP("config", "c", "", "配置文件位置")
	address := pflag.StringP("address", "a", "", "要启动的服务器地址")
	pflag.Parse()
	switch {
	case strings.ToUpper(*loglevel) == "TRACE":
		logger.Logger.SetLevel(logrus.TraceLevel)
	case strings.ToUpper(*loglevel) == "DEBUG":
		logger.Logger.SetLevel(logrus.DebugLevel)
	case strings.ToUpper(*loglevel) == "INFO":
		logger.Logger.SetLevel(logrus.InfoLevel)
	case strings.ToUpper(*loglevel) == "WARN":
		logger.Logger.SetLevel(logrus.WarnLevel)
	case strings.ToUpper(*loglevel) == "Error":
		logger.Logger.SetLevel(logrus.ErrorLevel)
	}
	if *confPath != "" {
		p, err := pathlib.New(*confPath).Absolute()
		if err != nil {
			logger.Logger.Error("指定的配置文件获取绝对位置失败")
		} else {
			if p.Exists() && p.IsFile() {
				filenameWithSuffix := path.Base(*confPath)
				fileSuffix := path.Ext(filenameWithSuffix)
				fileName := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
				dir, err := p.Parent()
				if err != nil {
					logger.Logger.Error("指定的配置文件获取父文件夹位置失败")
				} else {
					filePaths := []string{dir.Path}
					SetFileConfig(ConfigViper, fileName, filePaths)
				}

			}
		}
	}
	if *address != "" {
		ConfigViper.Set("Address", *address)
	}
	return true

}
