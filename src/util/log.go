package util

import (
	"log"
	"os"
)

// Common logger
var (
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	DebugLog *log.Logger
	WarnLog  *log.Logger
)

func init() {
	ErrorLog = log.New(os.Stderr, "[ERROR]", log.Lshortfile|log.Ldate|log.Ltime)
	InfoLog = log.New(os.Stderr, "[INFO]", log.Lshortfile|log.Ldate|log.Ltime)
	DebugLog = log.New(os.Stderr, "[DEBUG]", log.Lshortfile|log.Ldate|log.Ltime)
	WarnLog = log.New(os.Stderr, "[WARN]", log.Lshortfile|log.Ldate|log.Ltime)
}
