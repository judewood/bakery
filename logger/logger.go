package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type String string

func InitLogger(minLevel string) {
	fmt.Printf("\nLogging at level: %s\n", minLevel)
	opts := &slog.HandlerOptions{
		Level: ToLogLevel(minLevel),
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	slog.SetDefault(logger)
}

func ToLogLevel(s string) slog.Level {
	switch strings.ToLower(strings.Trim(s, " ")) {
	case "debug":
		return slog.LevelDebug
	case "warn":
	case "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	}
	return slog.LevelInfo
}
