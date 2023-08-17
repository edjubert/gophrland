package logging

import (
	"fmt"
	"io"
)

type Logger interface {
	Info(msg string)
	Infof(format string, args ...any)

	Warning(msg string)
	Warningf(format string, args ...any)

	Error(msg string)
	Errorf(format string, args ...any)
}

type fromWriter struct {
	output io.Writer
}

func (l *fromWriter) log(msg string) {
	if l.output != nil {
		_, _ = fmt.Fprintln(l.output, msg)
	}
}

func (l *fromWriter) Info(msg string)    { l.log(msg) }
func (l *fromWriter) Warning(msg string) { l.log(msg) }
func (l *fromWriter) Error(msg string)   { l.log(msg) }

func (l *fromWriter) logf(format string, args ...any) {
	if l.output != nil {
		_, _ = fmt.Fprintf(l.output, format+"\n", args...)
	}
}

func (l *fromWriter) Infof(format string, args ...any)    { l.logf(format, args...) }
func (l *fromWriter) Warningf(format string, args ...any) { l.logf(format, args...) }
func (l *fromWriter) Errorf(format string, args ...any)   { l.logf(format, args...) }

func New(out io.Writer) Logger {
	return &fromWriter{out}
}

// Noop is a noop logger
// All logs are discarded.
var Noop = &fromWriter{}
