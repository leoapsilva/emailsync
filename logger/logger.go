package logger

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	logrus.Logger
}

func Init() {

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.00000",
		DisableSorting:  true,
	})

	log.SetReportCaller(true)
	log.SetLevel(logrus.InfoLevel)
}
