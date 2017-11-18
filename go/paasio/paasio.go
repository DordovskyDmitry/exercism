package paasio

import (
	"io"
	"sync"
)

type ReadCounterImpl struct {
	reader    io.Reader
	byteCount int64
	opCount   int
	mux       sync.Mutex
}

func (rc *ReadCounterImpl) ReadCount() (n int64, nops int) {
	return rc.byteCount, rc.opCount
}

func (rc *ReadCounterImpl) Read(b []byte) (int, error) {
	res, err := rc.reader.Read(b)
	if err != nil {
		return 0, err
	}
	rc.mux.Lock()
	rc.opCount += 1
	rc.byteCount += int64(res)
	rc.mux.Unlock()
	return res, nil
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &ReadCounterImpl{reader: r}
}

type WriteCounterImpl struct {
	writer    io.Writer
	byteCount int64
	opCount   int
	mux       sync.Mutex
}

func (wc *WriteCounterImpl) WriteCount() (n int64, nops int) {
	return wc.byteCount, wc.opCount
}

func (wc *WriteCounterImpl) Write(b []byte) (int, error) {
	res, err := wc.writer.Write(b)
	if err != nil {
		return 0, err
	}
	wc.mux.Lock()
	wc.opCount += 1
	wc.byteCount += int64(res)
	wc.mux.Unlock()
	return res, nil

}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &WriteCounterImpl{writer: w}
}

type ReadWriteCounterImpl struct {
	readCounter  ReadCounter
	writeCounter WriteCounter
}

func (rwc *ReadWriteCounterImpl) Read(b []byte) (int, error) {
	return rwc.readCounter.Read(b)
}

func (rwc *ReadWriteCounterImpl) ReadCount() (n int64, nops int) {
	return rwc.readCounter.ReadCount()
}

func (rwc *ReadWriteCounterImpl) Write(b []byte) (int, error) {
	return rwc.writeCounter.Write(b)
}

func (rwc *ReadWriteCounterImpl) WriteCount() (n int64, nops int) {
	return rwc.writeCounter.WriteCount()
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	rc := NewReadCounter(rw)
	wc := NewWriteCounter(rw)
	return &ReadWriteCounterImpl{rc, wc}
}
