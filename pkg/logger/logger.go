package logger

import (
	"log"
	"os"
	"strings"
)

// LogLevel represents the logging level
type LogLevel int

const (
	Info LogLevel = iota
	Error
	Warn
	Debug
)

// Logger struct with Log method and set flags
type Logger struct {
	infoLogger  *log.Logger
	errLogger   *log.Logger
	debugLogger *log.Logger
	warnLogger  *log.Logger
	logLevel    LogLevel
}

// NewLogger returns a new logger
func New() *Logger {
	levelStr := strings.ToLower(os.Getenv("LOG_LEVEL"))

	logLevel := Info // Default log level
	if levelStr == "error" {
		logLevel = Error
	} else if levelStr == "warn" {
		logLevel = Warn
	} else if levelStr == "debug" {
		logLevel = Debug
	}

	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errLogger:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLogger:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		logLevel:    logLevel,
	}
}

// Info logs the info message
func (l *Logger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

// Infof logs with format
func (l *Logger) Infof(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

// Error logs the error message
func (l *Logger) Error(v ...interface{}) {
	l.errLogger.Println(v...)
}

// Debug logs the debug message if log level is Debug
func (l *Logger) Debug(v ...interface{}) {
	if l.logLevel >= Debug {
		l.debugLogger.Println(v...)
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.logLevel >= Debug {
		l.debugLogger.Printf(format, v...)
	}
}

// Warn logs the warn message if log level is Warn
func (l *Logger) Warn(v ...interface{}) {
	if l.logLevel >= Warn {
		l.warnLogger.Println(v...)
	}
}
