package logger

import (
	"log/slog"
	"os"
)

var globalLogger *slog.Logger

// Init создает новый объект Logger.
func Init(env string) {
	var handler slog.Handler

	switch env {
	case "prod":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	globalLogger = slog.New(handler)
}

// Debug записывает сообщение уровня Debug.
func Debug(msg string, args ...any) {
	globalLogger.Debug(msg, args...)
}

// Info записывает сообщение уровня Info.
func Info(msg string, args ...any) {
	globalLogger.Info(msg, args...)
}

// Warn записывает сообщение уровня Warn.
func Warn(msg string, args ...any) {
	globalLogger.Warn(msg, args...)
}

// Error записывает сообщение уровня Error.
func Error(msg string, args ...any) {
	globalLogger.Error(msg, args...)
}

// With добавляет дополнительные атрибуты к логгеру.
func With(args ...any) *slog.Logger {
	return globalLogger.With(args...)
}
