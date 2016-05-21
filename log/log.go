package log

// Logger is an interface that makes sure that a implementation is RFC5424 compatible
// https://tools.ietf.org/html/rfc5424#page-11
// Emergency shouldn't be used by applications, so it's omitted.
type Logger interface {
	//Alert: action must be taken immediately
	Alert(string, ...interface{})

	//Critical: critical conditions
	Crit(string, ...interface{})

	//Error: error conditions
	Error(string, ...interface{})

	//Warning: warning conditions
	Warn(string, ...interface{})

	//Notice: normal but significant condition
	Notice(string, ...interface{})

	//Informational: informational messages
	Info(string, ...interface{})

	//Debug: debug-level messages
	Debug(string, ...interface{})
}
