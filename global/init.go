package global

import (
	"os"

	"github.com/sirupsen/logrus"
)

/**
以环境变量设置最低log等级和log格式。

V2EXERP_LOGLEVEL 设置log等级，可选 debug, warn, error，若为其他值则取默认值info

V2EXERP_LOGFROMAT 设置log格式，可填写 text，若为其他值则取默认值json
*/

/**
以环境变量 V2EXERP_CONFIG 设置配置文件路径，如果不设置就取工作目录下的 app.ini 文件
*/

var (
	ConfigINIPath = "app.ini"
)

//Init : 初始化
func Init() {
	// 设置最低loglevel
	logLevel := os.Getenv("V2EXERP_LOGLEVEL")
	switch logLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	// 设置 log format
	logFormat := os.Getenv("V2EXERP_LOGFORMAT")
	switch logFormat {
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	default:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	logrus.SetOutput(os.Stdout)

	// 获取配置文件路径
	if path := os.Getenv("V2EXERP_CONFIG"); path != "" {
		ConfigINIPath = path
	}
}
