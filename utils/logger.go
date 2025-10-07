package utils

import (
	"log"
	"os"
	"sync"
)

var (
	onceLogger sync.Once
	instance   Logger
)

// Interface
type Logger interface {
	Info(msg string)
	Error(msg string)
}

// Implementations
type fileLogger struct {
	log *log.Logger
}

func (f *fileLogger) Info(msg string) {
	f.log.SetPrefix("INFO: ")
	f.log.Println(msg)
}

func (f *fileLogger) Error(msg string) {
	f.log.SetPrefix("Error: ")
	f.log.Println(msg)
}

// Constructor that will give the Singleton File Logger
func GetLogger() Logger {
	onceLogger.Do(func() {
		file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error in opening/creating the log file")
		}

		logger := log.New(file, "", log.Ldate|log.Ltime)

		instance = &fileLogger{
			log: logger,
		}
	})
	return instance
}
