package logconfigeration

import (
	"os"
	log "github.com/sirupsen/logrus"
	joonix "github.com/joonix/log"
)

// Setup for log configratioin
func Setup() {
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(joonix.NewFormatter())

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  log.SetLevel(log.DebugLevel)
}

// Info log
func Info(args ...interface{}) {
	log.Info(args)
}

// Warn log
func Warn(args ...interface{}) {
	log.Warn(args)
}

// Debug log
func Debug(args ...interface{}) {
	log.Debug(args)
}

// Error log
func Error(args ...interface{}) {
	log.Error(args)
}