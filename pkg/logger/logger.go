package logger

import (
	"log/slog"
	"os"
)

var log *slog.Logger

func InitLogger(level slog.Level) {
	log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
}

func GetLogger() *slog.Logger {
	if log == nil {
		InitLogger(slog.LevelInfo)
	}
	return log
}
