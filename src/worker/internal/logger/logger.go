package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const pathPrefix = "/app/"

type Logger struct {
	file   *os.File
	mu     sync.Mutex
	prefix string
}

func NewLogger(filePath, prefix string) (*Logger, error) {
	file, err := os.OpenFile(pathPrefix+filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	return &Logger{
		file:   file,
		prefix: prefix,
	}, nil
}

func (l *Logger) Log(data []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	logMessage := fmt.Sprintf("%s [%s] %s\n", time.Now().Format(time.RFC3339), l.prefix, string(data))
	_, err := l.file.WriteString(logMessage)
	if err != nil {
		return fmt.Errorf("failed to write log message: %w", err)
	}

	return nil
}

func (l *Logger) Close() error {
	return l.file.Close()
}
