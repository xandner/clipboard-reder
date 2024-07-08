package logger

import (
	"log"
	"os"
)

type Logger interface {
	Info(message string)
	Error(message string)
	Debug(message string)
}

type logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

// NewLogger creates a new logger that writes to a file.
func NewLogger(filename string) (Logger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &logger{
		infoLogger:  log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
	}, nil
}

func (l *logger) Info(message string) {
	l.infoLogger.Println(message)
}

func (l *logger) Error(message string) {
	l.errorLogger.Println(message)
}

func (l *logger) Debug(message string) {
	l.debugLogger.Println(message)
}
