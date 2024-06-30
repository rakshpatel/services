package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger implementation
var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.Out = os.Stdout
	Log.SetLevel(logrus.InfoLevel)
	Log.SetFormatter(&logrus.JSONFormatter{})
}