package logging

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"io"
)

type Logger interface {
	Info(msg string)
	Infof(format string, args ...any)

	Warning(msg string)
	Warningf(format string, args ...any)

	Error(msg string)
	Errorf(format string, args ...any)

	IsTTY() bool
}

type fromWriter struct {
	output io.Writer
	isTTY  bool
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

func (l *fromWriter) IsTTY() bool {
	return l.isTTY
}

func New(out io.Writer) Logger {
	fdGetter, ok := out.(interface{ Fd() uintptr })

	var isTTY bool
	if ok {
		isTTY = isatty.IsTerminal(fdGetter.Fd())
	}

	return &fromWriter{out, isTTY}
}

// Noop is a noop logger
// All logs are discarded.
var Noop = &fromWriter{}
