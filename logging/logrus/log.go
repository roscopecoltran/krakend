package logrus

/*
import (
	"io"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	// "github.com/o1egl/gormrus" // logs-logrus

	"github.com/roscopecoltran/krakend/config"
	"github.com/roscopecoltran/krakend/logging"
)
)

// NewLogger returns a krakend logger wrapping a gologging logger
func NewLogger(level string, out io.Writer, prefix string) (logging.Logger, error) {
	module := config.Config.Service.Name

	log := logrus.New()

	log.Formatter = new(prefixed.TextFormatter)

	// config.Config.Service.Debug.Level = INFO,WARN,ERROR,FATAL
	log.Level = logrus.DebugLevel

	// config.Config.Debug.Log.Formatter = JSON or TXT
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	// config.Config.Debug.Log.Output = Stderr, Stdout or File
	// log.SetOutput(os.Stdout)

	// You could set this to any `io.Writer` such as a file
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	//  log.Out = file
	// } else {
	//  log.Info("Failed to log to file, using default stderr")
	// }

	return Logger{log}, nil
}

// Logger is a wrapper over a github.com/op/go-logging logger
type Logger struct {
	Logger *logrus.Logger
}

// Debug implements the logger interface
func (l Logger) Debug(v ...interface{}) {
	l.Logger.Debug(v...)
}

// Info implements the logger interface
func (l Logger) Info(v ...interface{}) {
	l.Logger.Info(v...)
}

// Warning implements the logger interface
func (l Logger) Warning(v ...interface{}) {
	l.Logger.Warning(v...)
}

// Error implements the logger interface
func (l Logger) Error(v ...interface{}) {
	l.Logger.Error(v...)
}

// Critical implements the logger interface
func (l Logger) Critical(v ...interface{}) {
	l.Logger.Critical(v...)
}

// Fatal implements the logger interface
func (l Logger) Fatal(v ...interface{}) {
	l.Logger.Fatal(v...)
}
*/
