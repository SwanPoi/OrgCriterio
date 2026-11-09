package logger

import (
	"log/slog"
	"os"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

func Setup(env string, levelStr string) Logger {
	level := parseLogLevel(levelStr)
	var log *slog.Logger

	switch env {
	case "local":
		// Текстовый формат с читаемым выводом для локальной разработки
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}))
	default:
		// По умолчанию используем JSON-формат
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}))
	}

	return &SlogLogger{logger: log}
}

func parseLogLevel(levelStr string) slog.Level {
	switch levelStr {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}