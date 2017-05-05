package log

import (
    "github.com/Sirupsen/logrus"
)

var (
    log = logrus.New()
)

func init() {
    log.Formatter = &logrus.TextFormatter{
        ForceColors: true,
        TimestampFormat: "2006-01-02 15:04:05",
        FullTimestamp: true,
    }

    log.Level = logrus.DebugLevel
}

func Info(args ...interface{}) {
    log.Info(args...)
}

func Infof(format string, args ...interface{}) {
    log.Infof(format, args...)
}

func Fatal(args ...interface{}) {
    log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
    log.Fatalf(format, args...)
}

func Error(args ...interface{}) {
    log.Error(args...)
}
