package circular

import "errors"

const testVersion = 4

type Buffer struct {
	size         int
	readPointer  int
	writePointer int
	buf          []byte
	used         []bool
}

func NewBuffer(size int) *Buffer {
	buf := make([]byte, size)
	used := make([]bool, size)
	return &Buffer{size, 0, 0, buf, used}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.used[b.readPointer] {
		b.used[b.readPointer] = false
		defer IncreaseReadPointer(b)
		return b.buf[b.readPointer], nil
	}
	return 0, errors.New("Error while reading")
}

func (b *Buffer) WriteByte(c byte) error {
	if b.used[b.writePointer] {
		return errors.New("Error while writing")
	}
	b.used[b.writePointer] = true
	b.buf[b.writePointer] = c
	IncreaseWritePointer(b)
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.used[b.writePointer] {
		b.used[b.readPointer] = true
		b.buf[b.readPointer] = c
		IncreaseReadPointer(b)
	} else {
		b.WriteByte(c)
	}
}

func (b *Buffer) Reset() {
	for i := 0; i < b.size; i++ {
		b.used[i] = false
	}
	b.readPointer = 0
	b.writePointer = 0
}

func IncreaseWritePointer(b *Buffer) {
	b.writePointer = ((b.writePointer + 1) % b.size)
}

func IncreaseReadPointer(b *Buffer) {
	b.readPointer = ((b.readPointer + 1) % b.size)
}
