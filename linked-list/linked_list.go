package linkedlist

import "errors"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	next *Node
	prev *Node
	Val  interface{}
}

var ErrEmptyList = errors.New("List is empty")

func NewList(args ...interface{}) *List {
	l := &List{}
	for i, arg := range args {
		n := &Node{
			Val: arg,
		}

		if i == 0 {
			l.head = n
		} else {
			l.tail.next = n
			n.prev = l.tail
		}
		l.tail = n
	}
	return l
}

func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}
	return n.prev
}

func (l *List) PushFront(v interface{}) {
	n := &Node{Val: v}
	if l.head != nil {
		n.next = l.head
		l.head.prev = n
	} else {
		l.tail = n
	}
	l.head = n
}

func (l *List) PushBack(v interface{}) {
	n := &Node{Val: v}
	if l.tail != nil {
		n.prev = l.tail
		l.tail.next = n
	} else {
		l.head = n
	}
	l.tail = n
}

func (l *List) PopFront() (interface{}, error) {
	if l.head == nil {
		l.tail = nil
		return nil, ErrEmptyList
	}

	n := l.head

	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}

	n.next = nil
	return n.Val, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.tail == nil {
		l.head = nil
		return nil, ErrEmptyList
	}

	n := l.tail

	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	}

	n.prev = nil
	return n.Val, nil
}

func (l *List) Reverse() {
	for swapNode := l.head; swapNode != nil; swapNode = swapNode.prev {
		swapNext := swapNode.next
		swapNode.next = swapNode.prev
		swapNode.prev = swapNext
	}

	swap := l.head
	l.head = l.tail
	l.tail = swap
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
