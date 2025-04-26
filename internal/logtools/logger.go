package logtools

import (
	"log"
	"os"
)

var Logger *log.Logger

func Initialize() {
	cwd, _ := os.Getwd()
	filePath := cwd + "/logs/"
	logFilePath := filePath + "app.log"

	// Delete old app.log if it exists (optional)
	if _, err := os.Stat(logFilePath); err == nil {
		err := os.Remove(logFilePath)
		if err != nil {
			log.Fatalf("Failed to delete existing log file: %s", err)
		}
	}

	// Always ensure the logs/ directory exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create logs directory: %s", err)
		}
	}

	// Now open a fresh app.log
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	// Create the logger
	Logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Llongfile)
	Logger.Println("Logger initialized")
}
