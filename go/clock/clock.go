package clock

import "fmt"

const testVersion = 4

type Clock struct {
	h, m int
}

func New(hour, minute int) Clock {
	minutes := (hour*60 + minute) % (24 * 60)
	if minutes < 0 {
		minutes += (24 * 60)
	}
	h := minutes / 60
	m := minutes % 60
	return Clock{h, m}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.h, c.m+minutes)
}
