package applog

import (
	"fmt"
	"github.com/adnanahmady/go-websocket-chat/config"
	"log/slog"
	"os"
	"strings"
)

type Logger interface {
	New(fields ...any) Logger
	Info(format string, args ...any)
	Error(format string, args ...any)
	Debug(format string, args ...any)
	Warn(format string, args ...any)
}

var _ Logger = (*AppLogger)(nil)

type AppLogger struct {
	fields []any
	lgr    *slog.Logger
	cfg    *config.Config
}

func NewAppLogger(cfg *config.Config) *AppLogger {
	return &AppLogger{
		lgr: newSlogLogger(cfg),
		cfg: cfg,
	}
}

func newSlogLogger(cfg *config.Config) *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     getMinShowingLogLevel(cfg),
		AddSource: cfg.Log.ShowSource,
	}))
}

func getMinShowingLogLevel(cfg *config.Config) slog.Level {
	level := slog.LevelInfo
	switch cfg.Log.Level {
	case "debug":
		level = slog.LevelDebug
	case "error":
		level = slog.LevelError
	case "warn", "warning":
		level = slog.LevelWarn
	}
	return level
}

func (l *AppLogger) New(fields ...any) Logger {
	return &AppLogger{
		fields: fields,
		lgr:    newSlogLogger(l.cfg).With(fields...),
		cfg:    l.cfg,
	}
}

func (l *AppLogger) Info(format string, args ...any) {
	msg := processMsg(format, args)
	args = extractFields(format, args)
	l.lgr.Info(msg, args...)
}

func (l *AppLogger) Error(format string, args ...any) {
	msg := processMsg(format, args)
	args = extractFields(format, args)
	l.lgr.Error(msg, args...)
}

func (l *AppLogger) Debug(format string, args ...any) {
	msg := processMsg(format, args)
	args = extractFields(format, args)
	l.lgr.Debug(msg, args...)
}

func (l *AppLogger) Warn(format string, args ...any) {
	msg := processMsg(format, args)
	args = extractFields(format, args)
	l.lgr.Warn(msg, args...)
}

func processMsg(msg string, args []any) string {
	return fmt.Sprintf(msg, args[:findArgSplitPoint(msg)]...)
}

func extractFields(msg string, args []any) []any {
	return args[findArgSplitPoint(msg):]
}

func findArgSplitPoint(msg string) int {
	return strings.Count(msg, "%") - strings.Count(msg, "\\%")
}
