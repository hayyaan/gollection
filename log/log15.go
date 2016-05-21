package log

import "github.com/inconshreveable/log15"

type Log15 struct {
	logger log15.Logger
}

func NewLog15() Logger {
	return Log15{
		logger: log15.Root(),
	}
}

//Alert: action must be taken immediately
func (l Log15) Alert(msg string, args ...interface{}) {
	l.logger.Crit(msg, args...)
}

//Critical: critical conditions
func (l Log15) Crit(msg string, args ...interface{}) {
	l.logger.Crit(msg, args...)
}

//Error: error conditions
func (l Log15) Error(msg string, args ...interface{}) {
	l.logger.Error(msg, args...)
}

//Warning: warning conditions
func (l Log15) Warn(msg string, args ...interface{}) {
	l.logger.Warn(msg, args...)
}

//Notice: normal but significant condition
func (l Log15) Notice(msg string, args ...interface{}) {
	l.logger.Info(msg, args...)
}

//Informational: informational messages
func (l Log15) Info(msg string, args ...interface{}) {
	l.logger.Info(msg, args...)
}

//Debug: debug-level messages
func (l Log15) Debug(msg string, args ...interface{}) {
	l.logger.Debug(msg, args...)
}
