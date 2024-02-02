// Package paasio provides a solution for PaaS I/O on Exercism's Go Track.
package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
// NOTICE: practically we can use a single readWriteCounter struct here
// and use it in all three options (read, write, read & write),
// that would greatly reduce number or lines and duplicate logic,
// but diverge from the original task
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

// NewReadCounter returns a ReadCounter with Reader set to a given one
func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

// Read counts read operations and bytes while using underlying reader for reads
func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.reader.Read(p)
	if err != nil {
		return 0, err
	}
	rc.rwmutex.Lock()
	defer rc.rwmutex.Unlock()
	rc.nops++
	rc.bytes += int64(n)
	return n, nil
}

// ReadCount returns amount of bytes read and read operations so far
func (rc *readCounter) ReadCount() (int64, int) {
	rc.rwmutex.RLock()
	defer rc.rwmutex.RUnlock()
	return rc.bytes, rc.nops
}

// NewWriteCounter returns a WriteCounter with Writer set to a given one
func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

// Write counts write operations and bytes while using underlying writer for writes
func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.writer.Write(p)
	if err != nil {
		return 0, err
	}
	wc.rwmutex.Lock()
	defer wc.rwmutex.Unlock()
	wc.nops++
	wc.bytes += int64(n)
	return n, nil
}

// WriteCount returns amount of bytes wrote and write operations so far
func (wc *writeCounter) WriteCount() (int64, int) {
	wc.rwmutex.RLock()
	defer wc.rwmutex.RUnlock()
	return wc.bytes, wc.nops
}

// NewReadWriteCounter returns a ReadWriteCounter with ReadWriter set to a give one
func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		rc: NewReadCounter(readwriter),
		wc: NewWriteCounter(readwriter),
	}
}

// Read counts read operations and bytes while using underlying reader for reads
func (rwc *readWriteCounter) Read(p []byte) (int, error) {
	return rwc.rc.Read(p)
}

// Write counts write operations and bytes while using underlying writer for writes
func (rwc *readWriteCounter) Write(p []byte) (int, error) {
	return rwc.wc.Write(p)
}

// ReadCount returns amount of bytes read and read operations so far
func (rwc *readWriteCounter) ReadCount() (int64, int) {
	return rwc.rc.ReadCount()
}

// WriteCount returns amount of bytes wrote and write operations so far
func (rwc *readWriteCounter) WriteCount() (int64, int) {
	return rwc.wc.WriteCount()
}
