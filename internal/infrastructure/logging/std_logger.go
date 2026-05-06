package logging

import (
	"log/slog"
	"os"
)

type StdLogger struct {
	l *slog.Logger
}

func NewStdLogger() *StdLogger {
	return &StdLogger{
		l: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func (l *StdLogger) Info(msg string, fields ...any) {
	l.l.Info(msg, fields...)
}

func (l *StdLogger) Error(msg string, fields ...any) {
	l.l.Error(msg, fields...)
}
