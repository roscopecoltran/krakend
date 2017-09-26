package zap

/*
import (
	"io"

	"github.com/uber-go/zap"

	"github.com/roscopecoltran/krakend/config"
	"github.com/roscopecoltran/krakend/logging"
)

// Short-hand functions for logging.
var (
	Base64    = zap.Base64
	Bool      = zap.Bool
	Float64   = zap.Float64
	Int       = zap.Int
	Int64     = zap.Int64
	Marshaler = zap.Marshaler
	Nest      = zap.Nest
	Skip      = zap.Skip
	Stack     = zap.Stack
	String    = zap.String
	Stringer  = zap.Stringer
	Time      = zap.Time
	Uint      = zap.Uint
	Uint64    = zap.Uint64
	Uintptr   = zap.Uintptr
	Error     = zap.Error
)

// New returns new zap.Logger
func NewLogger() (level string, out io.Writer, prefix string) (logging.Logger, error) {
	_, filename, _, _ := runtime.Caller(1)
	name := filepath.Dir(truncFilename(filename))

	var enabler zap.AtomicLevel
	if e, ok := enablers[name]; ok {
		enabler = e
	} else {
		enabler = zap.DynamicLevel()
		enablers[name] = enabler
	}

	// 	if err != nil {
	//	 	return nil, err
	// 	}

	setLogLevelFromEnv(name, enabler)
	return Logger{
		zap.New(
			zap.NewTextEncoder(zap.TextNoTime()),
			enabler, addHook(),
		),
	}, nil
}

// Logger is a wrapper over a github.com/op/go-logging logger
type Logger struct {
	Logger *zap.Logger
	// Logger *logrus.Logger
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

// Interface ...
func Interface(key string, val interface{}) zap.Field {
	if val, ok := val.(fmt.Stringer); ok {
		return zap.Stringer(key, val)
	}
	return zap.Object(key, val)
}

// Duration formats time.Duration in human-readable time
func Duration(key string, val time.Duration) zap.Field {
	return zap.Stringer(key, val)
}

// Int32 ...
func Int32(key string, val int32) zap.Field {
	return zap.Int(key, int(val))
}

// Object ...
func Object(key string, val interface{}) zap.Field {
	return zap.Stringer(key, dump(val))
}
*/
