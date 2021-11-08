package logconfigeration

import (
	"os"
	log "github.com/sirupsen/logrus"
	joonix "github.com/joonix/log"
)

// Setup for log configratioin
func init() {
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(joonix.NewFormatter())

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  log.SetLevel(log.DebugLevel)
}