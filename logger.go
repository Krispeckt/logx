package logx

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
)

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006.01.02 15:04")

	var level string
	switch entry.Level {
	case logrus.InfoLevel:
		level = "\033[32mINFO\033[0m"
	case logrus.WarnLevel:
		level = "\033[33mWARNING\033[0m"
	case logrus.ErrorLevel:
		level = "\033[31mERROR\033[0m"
	case logrus.FatalLevel:
		level = "\033[31mFATAL\033[0m"
	case logrus.PanicLevel:
		level = "\033[31mPANIC\033[0m"
	case logrus.DebugLevel:
		level = "\033[36mDEBUG\033[0m"
	default:
		level = strings.ToUpper(entry.Level.String())
	}

	var fields []string
	for key, value := range entry.Data {
		fields = append(fields, fmt.Sprintf("%s=%v", key, value))
	}
	sort.Strings(fields)

	fieldsMessage := ""
	if len(fields) > 0 {
		fieldsMessage = " | " + fmt.Sprintf("%s", fields)
	}

	logMessage := fmt.Sprintf("[%s] - %s - %s:%d -> %s%s\n", timestamp, level, entry.Caller.Function, entry.Caller.Line, entry.Message, fieldsMessage)
	return []byte(logMessage), nil
}

func New() *logrus.Logger {
	logger := logrus.New()

	logger.SetFormatter(&CustomFormatter{})
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.DebugLevel)

	return logger
}
