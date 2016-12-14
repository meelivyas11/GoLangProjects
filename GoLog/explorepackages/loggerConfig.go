package explorepackages

// install github.com/gogap/logrus_mate using "go get"
import (
	"log"

	"github.com/Sirupsen/logrus"
	"github.com/gogap/logrus_mate"
	"github.com/rifflock/lfshook"
)

var loggerConf logrus_mate.LoggerConfig

func init() {
	loggerConf = logrus_mate.LoggerConfig{
		Level: "debug",
		Formatter: logrus_mate.FormatterConfig{
			Name: "text",
		},
		Out: logrus_mate.WriterConfig{
			Name:    "stderr",
			Options: nil,
		},
	}
}

func CreateLogger(name string) (logger *logrus.Logger) {
	logger, err := logrus_mate.NewLogger(name, loggerConf)
	if err != nil {
		log.Fatalf("Cannot obtain an instance of the logger due to %v", err)
	}

	logger.Hooks.Add(lfshook.NewHook(lfshook.PathMap{
		logrus.DebugLevel: "meeli.log",
		logrus.InfoLevel:  "meeli.log",
		logrus.ErrorLevel: "meeli.log",
	}))
	return logger
}
