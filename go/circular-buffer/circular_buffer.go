// Package circular provides a solution for the Circular Buffer on Exercism's Go Track.
package circular

import "fmt"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Buffer type
type Buffer struct {
	buf        []byte
	full       bool
	readIndex  int
	writeIndex int
}

// NewBuffer creates a buffer of given size
func NewBuffer(size int) *Buffer {
	return &Buffer{
		buf:        make([]byte, size),
		readIndex:  0,
		writeIndex: 0,
	}
}

// ReadByte returns next byte from buffer
// Errors if buffer is empty (nothing to read)
func (b *Buffer) ReadByte() (byte, error) {
	if b.readIndex == b.writeIndex && !b.full {
		return 0, fmt.Errorf("buffer empty")
	}

	b.full = false
	value := b.buf[b.readIndex]
	b.readIndex = (b.readIndex + 1) % len(b.buf)
	return value, nil
}

// WriteBytes writes next byte to the buffer
// Errors if buffer is full already
func (b *Buffer) WriteByte(c byte) error {
	fmt.Println("writeIndex", b.writeIndex, "readIndex", b.readIndex)
	if b.full {
		return fmt.Errorf("buffer full")
	}
	b.buf[b.writeIndex] = c
	b.writeIndex = (b.writeIndex + 1) % len(b.buf)
	if b.writeIndex == b.readIndex {
		b.full = true
	}
	return nil
}

// Overwrite writes next byte to the buffer even if it's full already
func (b *Buffer) Overwrite(c byte) {
	if b.full {
		b.full = false
		b.readIndex = (b.readIndex + 1) % len(b.buf)
	}
	b.WriteByte(c)
}

// Reset resets buffer to initial state
func (b *Buffer) Reset() {
	b.full = false
	b.readIndex = 0
	b.writeIndex = 0
}
