package logger

import (
	"log/slog"
	"os"
)

var globalLogger *slog.Logger

// Init initializes a new Logger object.
func Init(env string) {
	var handler slog.Handler

	switch env {
	case "prod":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true,
		})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})
	}

	globalLogger = slog.New(handler)
}

// Debug logs a message at the Debug level.
func Debug(msg string, args ...any) {
	globalLogger.Debug(msg, args...)
}

// Info logs a message at the Info level.
func Info(msg string, args ...any) {
	globalLogger.Info(msg, args...)
}

// Warn logs a message at the Warn level.
func Warn(msg string, args ...any) {
	globalLogger.Warn(msg, args...)
}

// Error logs a message at the Error level.
func Error(msg string, args ...any) {
	globalLogger.Error(msg, args...)
}

// Fatal logs a message at the Error level and exits the program.
func Fatal(msg string, args ...any) {
	globalLogger.Error(msg, args...)
	os.Exit(1)
}

// With adds additional attributes to the logger.
func With(args ...any) *slog.Logger {
	return globalLogger.With(args...)
}
