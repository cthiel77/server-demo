package log

import (
	//_ use native embed
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/cthiel77/server-demo/internal/cfg"
	"github.com/cthiel77/server-demo/internal/meta/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// Fields wrapper to provide log fields
type Fields logrus.Fields

// LoggerHook the global log hook
type LoggerHook struct {
	AppName *string
}

var (
	// LogLevelDefault the default log level
	LogLevelDefault = "info"
	// LogLevel the app log level
	LogLevel = LogLevelDefault
	//go_embed usage.txt
	logLevelUsage string
	// TimestampFormat time format
	TimestampFormat = "2006-01-02T15:04:05.000"
	// DisableTimestamp disables timestamps
	DisableTimestamp = false
	// DisableColors disables colors
	DisableColors = false
	// FullTimestamp shows the full timestamp
	FullTimestamp = true
)

/// the default logger

// Levels levels
func (l *LoggerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire fire
func (l *LoggerHook) Fire(e *logrus.Entry) error {
	if l.AppName != nil {
		e.Data["app"] = *l.AppName
	}
	return nil
}

// ErrorHook the error logger
type ErrorHook struct {
}

// Levels set levels to be used for logging
func (h *ErrorHook) Levels() []logrus.Level {
	// fire only on ErrorLevel (.Error(), .Errorf(), etc.) and WarnLevel
	return []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel}
}

// Fire extract location and status info from error message and reset message
func (h *ErrorHook) Fire(e *logrus.Entry) error {
	// e.Data is a map with all fields attached to entry
	myExp := regexp.MustCompile(`Domain:\s+(?P<domain>\d+)\s+Code:\s+(?P<code>\d+)\s+Message:\s+`)

	match := myExp.FindStringSubmatch(e.Message)
	if len(match) == 0 {
		return nil
	}

	for i, name := range myExp.SubexpNames() {
		if i == 0 {
			continue
		}
		v := match[i]
		e.Data[name] = v
	}
	e.Message = strings.Replace(e.Message, match[0], "", -1)
	return nil
}

// NewLogger returns a configured logger
func NewLogger(app model.AppMeta) *logrus.Logger {
	appMode := cfg.GetAppModeKey()

	//AppLogger the app logger
	logger := logrus.New()

	logger.AddHook(&LoggerHook{
		AppName: &app.Name,
	})

	logger.AddHook(&ErrorHook{})

	/** levels
	------------
		"panic" // PanicLevel level, highest level of severity. Logs and then calls panic with the message passed to Debug, Info, ...
		"fatal" // FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the logging level is set to Panic.
		"error" // ErrorLevel level. Logs. Used for errors that should definitely be noted. Commonly used for hooks to send errors to an error tracking service.
	    "warn", "warning" // WarnLevel level. Non-critical entries that deserve eyes.
		"info"  // InfoLevel level. General operational entries about what's going on inside the application.
		"debug" // DebugLevel level. Usually only enabled when debugging. Very verbose logging.
		"trace" // TraceLevel level. Designates finer-grained informational events than the Debug.
	*/
	lvl, parseErr := logrus.ParseLevel(GetLogLevel())
	if parseErr != nil {
		logrus.Warnf("%+v", parseErr)
		lvl, _ = logrus.ParseLevel(LogLevelDefault)
	}
	logger.SetLevel(lvl)

	if appMode == "dev" {
		Formatter := new(logrus.TextFormatter)
		Formatter.TimestampFormat = TimestampFormat
		Formatter.DisableTimestamp = DisableTimestamp
		Formatter.DisableColors = DisableColors
		Formatter.FullTimestamp = FullTimestamp
		logger.SetFormatter(Formatter)

		logger.SetReportCaller(false)
	} else {
		// everything else is prod
		Formatter := new(logrus.JSONFormatter)
		Formatter.TimestampFormat = TimestampFormat
		Formatter.DisableTimestamp = DisableTimestamp
		logger.SetFormatter(Formatter)
	}

	logger.SetOutput(os.Stdout)

	return logger
}

// InitFlags initializes application flags
func InitFlags() {
	pflag.StringVarP(&LogLevel, "log-level", "l", LogLevelDefault, logLevelUsage)
}

// GetLogLevel returns the application log level
func GetLogLevel() string {
	if cfg.Mode == "prod" {
		return "error"
	}
	return LogLevel
}

// PrintConfig prints the config
func PrintConfig() {
	fmt.Printf("AppMode: %v\n", cfg.Mode)
	fmt.Printf("logLevel: %v\n", LogLevel)
	fmt.Printf("TimestampFormat: %v\n", TimestampFormat)
	fmt.Printf("DisableTimestamp: %v\n", DisableTimestamp)
	fmt.Printf("DisableColors: %v\n", DisableColors)
	fmt.Printf("FullTimestamp: %v\n", FullTimestamp)
}
