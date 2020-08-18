package logger

import "io"

type writer struct {
	output []byte
}

func (w *writer) Write(p []byte) (n int, err error) {
	w.output = append(w.output, p...)
	return len(p), nil
}

func NewWriter() io.Writer {
	return &writer{}
}
