package logging

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
}
