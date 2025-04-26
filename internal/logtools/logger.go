// Package logtools provides a simple logging utility that initializes
// a logger writing to a local file with date, time, and file location.
package logtools

import (
	"log"
	"os"
)

var Logger *log.Logger

// Initialize sets up the logging system.
// It creates a "logs" directory if it does not exist,
// deletes any existing "app.log" file, and creates a fresh one.
// All logs are written with date, time, and full file path information.
func Initialize() {
	cwd, _ := os.Getwd()
	filePath := cwd + "/logs/"
	logFilePath := filePath + "app.log"

	// Delete old app.log file if it exists
	if _, err := os.Stat(logFilePath); err == nil {
		err := os.Remove(logFilePath)
		if err != nil {
			log.Fatalf("Failed to delete existing log file: %s", err)
		}
	}

	// Ensure the logs/ directory exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create logs directory: %s", err)
		}
	}

	// Open a fresh app.log file for writing
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	// Create a new logger writing to the log file
	Logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Llongfile)
	Logger.Println("Logger initialized")
}
