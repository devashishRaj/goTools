package safeWriter

import "io"

type SafeWriter struct {
	w    io.Writer
	WErr error
}

func NewSafeWriter(writer io.Writer) *SafeWriter {
	return &SafeWriter{w: writer}
}

func (sw *SafeWriter) Write(data []byte) {
	if sw.WErr != nil {
		return
	}
	_, err := sw.w.Write(data)
	if err != nil {
		sw.WErr = err
	}
}
