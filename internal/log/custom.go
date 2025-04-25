package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CustomLogger struct {
	logger *zap.SugaredLogger
}

func NewCustomLogger(fname string) *CustomLogger {
	writeSyncer := getLogWriter(fname)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	return &CustomLogger{
		logger: zap.New(core, zap.AddCaller()).Sugar(),
	}
}

func (l *CustomLogger) Fatal(v ...interface{}) {
	l.logger.Fatal(v...)
}

func (l *CustomLogger) Error(v ...interface{}) {
	l.logger.Error(v...)
}

func (l *CustomLogger) Info(v ...interface{}) {
	l.logger.Info(v...)
}

func (l *CustomLogger) Warn(v ...interface{}) {
	l.logger.Warn(v...)
}

func (l *CustomLogger) Debug(v ...interface{}) {
	l.logger.Debug(v...)
}

func (l *CustomLogger) Fatalf(fmt string, v ...interface{}) {
	l.logger.Fatalf(fmt, v...)
}

func (l *CustomLogger) Errorf(fmt string, v ...interface{}) {
	l.logger.Errorf(fmt, v...)
}

func (l *CustomLogger) Infof(fmt string, v ...interface{}) {
	l.logger.Infof(fmt, v...)
}

func (l *CustomLogger) Warnf(fmt string, v ...interface{}) {
	l.logger.Warnf(fmt, v...)
}

func (l *CustomLogger) Debugf(fmt string, v ...interface{}) {
	l.logger.Debugf(fmt, v...)
}
