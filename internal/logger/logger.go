package logger

import (
	"log"
	"os"
)

var (
    // Logger is the instance of the logger
    Logger *log.Logger
)

// InitLogger initializes the logger
func InitLogger() {
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }

    Logger = log.New(file, "APP ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs informational messages
func Info(message string) {
    Logger.Println("INFO: " + message)
}

// Error logs error messages
func Error(message string) {
    Logger.Println("ERROR: " + message)
}

// Fatal logs fatal messages and exits
func Fatal(message string) {
    Logger.Fatal("FATAL: " + message)
}