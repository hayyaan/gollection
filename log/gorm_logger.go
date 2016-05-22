package log

type GormLogger struct {
	logger Logger
}

func NewGormLogger(l Logger) GormLogger {
	return GormLogger{
		logger: l,
	}
}

func (gl GormLogger) Print(v ...interface{}) {
	gl.logger.Debug("gorm", "path", v[1], "duration", v[2], "query", v[3], "param", v[4])
}
