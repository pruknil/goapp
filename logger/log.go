package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"

	//"runtime"
	"fmt"
	"time"
)

//const LogFilePath = "/tmp/"
const FUNCNM = "funcnm"
const STATUS = "status"
const RSUID = "rsuid"
const TM = "tm"

func New(fileName, lvl string) *logrus.Logger {
	LogFilePath := os.Getenv("LOG_PATH")
	if LogFilePath == "" {
		LogFilePath = "/tmp"
	}

	logg := logrus.New()
	level, err := logrus.ParseLevel(lvl)
	if err == nil {
		logg.SetLevel(level)
	}
	lumberjackLogrotate := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s-%s.log", LogFilePath, fileName, os.Getenv("CARDSERVICE_HTTP_PORT")),
		MaxSize:    50, // Max megabytes before log is rotated
		MaxBackups: -1, // Max number of old log files to keep
		MaxAge:     90, // Max number of days to retain log files

		Compress: false,
	}
	switch fileName {
	case "perf":
		logg.SetFormatter(&perfFormatter{})
	case "rest":
		logg.SetFormatter(&restFormatter{})
	default:
		logg.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC3339})
	}

	logMultiWriter := io.MultiWriter(os.Stdout, lumberjackLogrotate) //os.Stdout,
	logg.SetOutput(logMultiWriter)

	return logg
}

type perfFormatter struct {
	logrus.TextFormatter
}

func (f *perfFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return append([]byte(fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s", entry.Time.Format("2006-01-02T15:04:05.000000Z07:00"), entry.Data[FUNCNM], entry.Message, entry.Data[TM], entry.Data[STATUS], "\n"))), nil
}

type restFormatter struct {
	logrus.TextFormatter
}

func (f *restFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return append([]byte(fmt.Sprintf("%s\t%s\t%s\n", entry.Time.Format("2006-01-02T15:04:05.000000Z07:00"), entry.Data[RSUID], entry.Message))), nil

}

type AppLog struct {
	Trace *logrus.Logger
	Perf  *logrus.Logger
	Error *logrus.Logger
	Rest  *logrus.Logger
	Audit *logrus.Logger
}
