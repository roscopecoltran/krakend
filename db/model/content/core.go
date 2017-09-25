package cms

import (
	// "github.com/roscopecoltran/sniperkit-limo/config"      // app-config
	"github.com/sirupsen/logrus"                           // logs-logrus
	prefixed "github.com/x-cray/logrus-prefixed-formatter" // logs-logrus
)

func init() {
	log.Out = os.Stdout                      // logs
	formatter := new(prefixed.TextFormatter) // logs
	log.Formatter = formatter                // logs
	log.Level = logrus.DebugLevel            // logs
}
