package logrus

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

var (
	log *logrus.Logger
)

func init() {
	log = logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}
}

func Info(args ...interface{}) {
	log.Info(args...)
}
