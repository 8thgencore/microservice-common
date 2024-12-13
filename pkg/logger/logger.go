package logger

import (
	"log/slog"
	"os"
	"runtime"
)

var logger *slog.Logger

// Init initializes a new Logger object.
func Init(env string) {
	var handler slog.Handler

	var opts = &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				pc, f, l, _ := runtime.Caller(7)
				a.Value = slog.GroupValue(
					slog.Attr{
						Key:   "file",
						Value: slog.StringValue(f),
					},
					slog.Attr{
						Key:   "line",
						Value: slog.IntValue(l),
					},
					slog.Attr{
						Key:   "function",
						Value: slog.StringValue(runtime.FuncForPC(pc).Name()),
					},
				)
			}
			return a
		},
	}

	switch env {
	case "prod":
		opts.Level = slog.LevelInfo
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		opts.Level = slog.LevelDebug
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	logger = slog.New(handler)
}

// Debug logs a message at the Debug level.
func Debug(msg string, args ...any) {
	logger.Debug(msg, args...)
}

// Info logs a message at the Info level.
func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

// Warn logs a message at the Warn level.
func Warn(msg string, args ...any) {
	logger.Warn(msg, args...)
}

// Error logs a message at the Error level.
func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}

// Fatal logs a message at the Error level and exits the program.
func Fatal(msg string, args ...any) {
	logger.Error(msg, args...)
	os.Exit(1)
}

// With adds additional attributes to the logger.
func With(args ...any) *slog.Logger {
	return logger.With(args...)
}
