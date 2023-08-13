package logger

import (
	"errors"
	"io"
	"log/slog"
)

type Logger interface {
	Info(msg string)
	Error(msg string, err error)
	Warn(msg string)
}

type ELLogger struct {
	logger *slog.Logger
}

var _ Logger = (*ELLogger)(nil)

const serviceName = "evidently-local"

func NewEvidentlyLocalLogger(out io.Writer) (*ELLogger, error) {
	if out == nil {
		return nil, errors.New("out is nil")
	}

	handler := slog.NewJSONHandler(out, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	logger := slog.New(handler)

	return &ELLogger{logger: logger}, nil
}

func (l *ELLogger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *ELLogger) Error(msg string, err error) {
	if err == nil {
		l.logger.Error(msg)
		return
	}
	l.logger.Error(msg, slog.String("error", err.Error()))
}

func (l *ELLogger) Warn(msg string) {
	l.logger.Warn(msg)
}
