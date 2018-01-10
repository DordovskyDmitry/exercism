package linkedlist

import (
	"errors"
)

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(values []int) *List {
	l := new(List)
	if len(values) == 0 {
		return l
	}
	current := &Element{data: values[0]}

	l.size = 1
	l.head = current

	for _, v := range values[1:] {
		e := &Element{data: v}
		current.next = e
		current = e
		l.size += 1
	}

	return l
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Push(v int) {
	if list.size == 0 {
		e := &Element{data: v}
		list.head = e
		list.size = 1
		return
	}
	current := list.head
	for {
		if current.next == nil {
			break
		}
		current = current.next
	}
	e := &Element{data: v}
	current.next = e
	list.size += 1
}

func (list *List) Pop() (int, error) {
	if list.size == 0 {
		return 0, errors.New("Empty list")
	}
	if list.size == 1 {
		list.size = 0
		v := list.head.data
		list.head = nil
		return v, nil
	}
	current := list.head
	for {
		if current.next.next == nil {
			break
		}
		current = current.next
	}
	v := current.next.data
	current.next = nil
	list.size -= 1
	return v, nil
}

func (list *List) Array() (res []int) {
	if list.size == 0 {
		return res
	}
	current := list.head
	for {
		res = append(res, current.data)
		if current.next == nil {
			break
		}
		current = current.next
	}
	return
}

func (list *List) Reverse() *List {
	if list.size < 2 {
		return list
	}
	l := new(List)
	for list.size > 0 {
		v, _ := list.Pop()
		l.Push(v)
	}
	return l
}
