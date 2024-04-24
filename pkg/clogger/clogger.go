package clogger

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

type CloggerOpt struct {
	Level       string
	Environment string
}

func Setup(opt CloggerOpt) {
	level, err := logrus.ParseLevel(opt.Level)
	if err != nil {
		panic(err)
	}

	var log = &logrus.Logger{
		Out:   os.Stderr,
		Hooks: make(logrus.LevelHooks),
		Level: level,
	}

	if opt.Environment == "local" {
		log.Formatter = &logrus.TextFormatter{
			FullTimestamp: true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				return "", fmt.Sprintf("(%s:%d)", f.File, f.Line)
			},
		}
	} else {
		log.Formatter = &logrus.JSONFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				return "", fmt.Sprintf("(%s:%d)", f.File, f.Line)
			},
		}
	}

	logger = log
}

func Logger() *logrus.Logger {
	return logger
}
