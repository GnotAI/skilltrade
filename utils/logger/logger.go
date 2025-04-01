package logger

import (
  "os"
  "path/filepath"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger initializes a zap logger
func NewLogger() *zap.Logger {
	// Define log directory and file
	logDir := "logs"
	logFile := filepath.Join(logDir, "server.log")

	// Ensure the logs directory exists
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, 0755)
		if err != nil {
			panic("Failed to create logs directory: " + err.Error())
		}
	}

	// Open or create the log file
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}
	// Create a core that writes logs to the file
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(file),
		zap.InfoLevel,
	)

	// Create a logger that logs to both console and file
	logger := zap.New(fileCore, zap.AddCaller())

	return logger
}
