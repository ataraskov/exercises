// Package paasio provides a solution for PaaS I/O on Exercism's Go Track.
package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	reader  io.Reader
	rwmutex sync.RWMutex
	nops    int
	bytes   int64
}
type writeCounter struct {
	writer  io.Writer
	rwmutex sync.RWMutex
	nops    int
	bytes   int64
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.
type readWriteCounter struct {
	rc ReadCounter
	wc WriteCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		rc: NewReadCounter(readwriter),
		wc: NewWriteCounter(readwriter),
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.rwmutex.Lock()
	defer rc.rwmutex.Unlock()
	n, err := rc.reader.Read(p)
	if err != nil {
		return 0, err
	}
	rc.nops++
	rc.bytes += int64(n)
	return n, nil
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.rwmutex.RLock()
	defer rc.rwmutex.RUnlock()
	return rc.bytes, rc.nops
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.rwmutex.Lock()
	defer wc.rwmutex.Unlock()
	n, err := wc.writer.Write(p)
	if err != nil {
		return 0, err
	}
	wc.nops++
	wc.bytes += int64(n)
	return n, nil
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.rwmutex.RLock()
	defer wc.rwmutex.RUnlock()
	return wc.bytes, wc.nops
}

func (rwc *readWriteCounter) Read(p []byte) (int, error) {
	return rwc.rc.Read(p)
}

func (rwc *readWriteCounter) Write(p []byte) (int, error) {
	return rwc.wc.Write(p)
}

func (rwc *readWriteCounter) ReadCount() (int64, int) {
	return rwc.rc.ReadCount()
}

func (rwc *readWriteCounter) WriteCount() (int64, int) {
	return rwc.wc.WriteCount()
}
