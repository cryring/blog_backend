package log

type GormLogger struct {
	logger *CustomLogger
}

func NewGormLogger(fname string) *GormLogger {
	return &GormLogger{
		logger: NewCustomLogger(fname),
	}
}

func (l *GormLogger) Printf(fmt string, v ...interface{}) {
	l.logger.Debugf(fmt, v...)
}
