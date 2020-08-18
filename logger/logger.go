package logger

import (
	"fmt"
	"io"
)

type Logger interface {
	Log(string)
}

type logger struct {
	prefix string
	writer io.Writer
}

func (l *logger) Log(s string) {
	fmt.Fprintf(l.writer, "[%s] %s", l.prefix, s)
}

func NewLogger(prefix string) Logger {
	return &logger{
		prefix: prefix,
		writer: NewWriter(),
	}
}
