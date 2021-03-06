package helper

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	LogTypeAccess = "access_log"
	LogTypeError  = "error_log"
	LogTypeInfo   = "info_log"
	LogTypeFatal  = "fatal_log"
)

func Log(logType string, str string) {
	switch logType {
	case LogTypeAccess:
		{
			logF := logFile(LogTypeAccess)
			logF.Info(str)
		}
	case LogTypeInfo:
		{
			logF := logFile(LogTypeInfo)
			logC := logConsole()

			logF.Info(str)
			logC.Info(str)
		}
	case LogTypeError:
		{
			logF := logFile(LogTypeError)
			logC := logConsole()

			logF.Error(str)
			logC.Error(str)

			sentryGo := sentry.CurrentHub().Clone()
			sentryGo.CaptureMessage(str)
		}
	case LogTypeFatal:
		{
			logC := logConsole()

			sentryGo := sentry.CurrentHub().Clone()
			sentryGo.CaptureMessage(str)

			logC.Fatal(str)
		}
	default:
		return
	}

}

func logFile(logFileName string) *logrus.Logger {

	currentTime := time.Now()

	logFileName = logFileName + "_" + currentTime.Format("2006-01-02") + ".log"

	logPath := os.Getenv("LOG_PATH")

	fileName := path.Join(logPath, logFileName)

	source, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Failed to write/create file")
		os.Exit(1)
	}

	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(source)

	return logger
}

func logConsole() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	return logger
}
