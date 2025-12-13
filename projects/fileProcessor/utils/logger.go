package utils

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	file  *os.File
	mutex sync.Mutex
}

func NewLogger(path string) (*Logger, error) {
	file, err := os.OpenFile(
		"./"+path,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)

	if err != nil {
		return nil, err
	}

	return &Logger{file: file}, nil
}

func (l *Logger) log(level string, msg string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	timestamp := time.Now().Format(time.RFC3339)

	line := fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, msg)

	fmt.Print(line)

	l.file.WriteString(line)
}

func (l *Logger) Info(msg string) {
	l.log("Info", msg)
}

func (l *Logger) Error(msg string) {
	l.log("Error", msg)
}

func (l *Logger) Close() {
	l.file.Close()
}
