package logconf

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Init Setup for log conf configuration
func Init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp:true})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log conf the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}