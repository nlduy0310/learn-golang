package stack

import (
	"fmt"
	"strings"
)

// DYNAMIC STACK NODE
type DStackNode struct {
	data int
	next *DStackNode // points to the element under this in a stack
}

func newDStackNode(data int) *DStackNode {
	return &DStackNode{
		data: data,
		next: nil,
	}
}

func (n DStackNode) String() string {
	return fmt.Sprintf("%d -> ", n.data)
}

// DYNAMIC STACK
type DStack struct {
	top *DStackNode
}

// -- BASICS
func (s DStack) IsEmpty() bool {
	return s.top == nil
}

func (s DStack) Size() int {
	res := 0
	cur := s.top
	for ; cur != nil; cur = cur.next {
		res++
	}
	return res
}

func (s DStack) String() string {
	builder := strings.Builder{}
	builder.WriteString("top -> ")
	for tmp := s.top; tmp != nil; tmp = tmp.next {
		builder.WriteString(tmp.String())
	}
	builder.WriteString("bottom")
	return builder.String()
}

// -- PUSH
func (s *DStack) Push(val int) *DStackNode {
	tmp := newDStackNode(val)
	tmp.next = s.top
	s.top = tmp
	return s.top
}

// -- POP
func (s *DStack) Pop() *DStackNode {
	tmp := s.top
	if tmp != nil {
		s.top = tmp.next
		tmp.next = nil
	}
	return tmp
}

// -- PEEK
func (s DStack) Peek() *DStackNode {
	return s.top
}

// -- CLEAR
func (s *DStack) Clear() {
	s.top = nil
}
