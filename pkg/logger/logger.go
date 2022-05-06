package logger

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
)

func init() {
	WarningLogger = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
}
