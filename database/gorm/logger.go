package gorm

import "github.com/go-kit/kit/log"

type gormLogger struct {
	logger log.Logger
}

// Create a new gorm log.Logger wrapper
func NewLogger(logger log.Logger) *gormLogger {
	return &gormLogger{logger: logger}
}

func (gl gormLogger) Print(v ...interface{}) {
	if len(v) == 5 {
		gl.logger.Log(
			"path", v[1],
			"duration", v[2],
			"query", v[3],
			"param", v[4],
		)
	}
}
