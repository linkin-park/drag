package logger

import (
	"log"
	"os"
)

// Common logger
var (
	errorLog *log.Logger
	infoLog  *log.Logger
	debugLog *log.Logger
	warnLog  *log.Logger
)

const defaultLogLevel = iota + 1

// LogLevel to be used for logging
type LogLevel int

// loglevel in order
const (
	all = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	fatalLevel
	off
)

// color codes for terminal
var (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

func init() {
	errorLog = log.New(os.Stderr, red+"[ERROR]"+reset, log.Lshortfile|log.Ldate|log.Ltime)
	infoLog = log.New(os.Stdout, blue+"[INFO]"+reset, log.Lshortfile|log.Ldate|log.Ltime)
	debugLog = log.New(os.Stdout, cyan+"[DEBUG]"+reset, log.Lshortfile|log.Ldate|log.Ltime)
	warnLog = log.New(os.Stdout, yellow+"[WARN]"+reset, log.Lshortfile|log.Ldate|log.Ltime)
}

// Println Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func Println(ll LogLevel, x ...interface{}) {

	if ll >= defaultLogLevel {
		switch {
		case ll == ErrorLevel:
			errorLog.Println(x...)
		case ll == DebugLevel:
			debugLog.Println(x...)
		case ll == InfoLevel:
			infoLog.Println(x...)
		case ll == WarnLevel:
			warnLog.Println(x...)
		}
	}
}

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(ll LogLevel, format string, x ...interface{}) {

	if ll >= defaultLogLevel {
		switch {
		case ll == ErrorLevel:
			errorLog.Printf(format, x...)
		case ll == DebugLevel:
			debugLog.Printf(format, x...)
		case ll == InfoLevel:
			infoLog.Printf(format, x...)
		case ll == WarnLevel:
			warnLog.Printf(format, x...)
		}
	}
}
