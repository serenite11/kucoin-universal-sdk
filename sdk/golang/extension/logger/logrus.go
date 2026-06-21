package logger

import (
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/sirupsen/logrus"
)

type LogrusAdapter struct {
	logger *logrus.Logger
}

func NewLogrusAdapter() *LogrusAdapter {
	return &LogrusAdapter{
		logger: logrus.New(),
	}
}

func NewLogrusAdapterWith(l *logrus.Logger) *LogrusAdapter {
	return &LogrusAdapter{
		logger: l,
	}
}

func (l *LogrusAdapter) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}

func (l *LogrusAdapter) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *LogrusAdapter) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LogrusAdapter) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *LogrusAdapter) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LogrusAdapter) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *LogrusAdapter) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *LogrusAdapter) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}

func (l *LogrusAdapter) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *LogrusAdapter) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *LogrusAdapter) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *LogrusAdapter) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *LogrusAdapter) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *LogrusAdapter) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l *LogrusAdapter) SetLevel(level logger.Level) {
	switch level {
	case logger.TRACE:
		l.logger.SetLevel(logrus.TraceLevel)
	case logger.DEBUG:
		l.logger.SetLevel(logrus.DebugLevel)
	case logger.INFO:
		l.logger.SetLevel(logrus.InfoLevel)
	case logger.WARN:
		l.logger.SetLevel(logrus.WarnLevel)
	case logger.ERROR:
		l.logger.SetLevel(logrus.ErrorLevel)
	case logger.FATAL:
		l.logger.SetLevel(logrus.FatalLevel)
	case logger.PANIC:
		l.logger.SetLevel(logrus.PanicLevel)
	}
}

func (l *LogrusAdapter) GetLevel() logger.Level {
	switch l.logger.GetLevel() {
	case logrus.TraceLevel:
		return logger.TRACE
	case logrus.DebugLevel:
		return logger.DEBUG
	case logrus.InfoLevel:
		return logger.INFO
	case logrus.WarnLevel:
		return logger.WARN
	case logrus.ErrorLevel:
		return logger.ERROR
	case logrus.FatalLevel:
		return logger.FATAL
	case logrus.PanicLevel:
		return logger.PANIC
	default:
		return logger.INFO
	}
}
