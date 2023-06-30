package logger

import (
	"errors"
	"io"

	"github.com/rs/zerolog"
)

type Logger interface {
	Info(msg string)
	Error(msg string, err error)
	Warn(msg string)
}

type ELLogger struct {
	logger *zerolog.Logger
}

var _ Logger = (*ELLogger)(nil)

const serviceName = "evidently-local"

func NewEvidentlyLocalLogger(out io.Writer) (*ELLogger, error) {
	if out == nil {
		return nil, errors.New("out is nil")
	}

	logger := zerolog.New(out).With().Timestamp().Str("service", serviceName).Logger()
	return &ELLogger{logger: &logger}, nil
}

func (l *ELLogger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *ELLogger) Error(msg string, err error) {
	l.logger.Error().Err(err).Msg(msg)
}

func (l *ELLogger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}
