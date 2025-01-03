package linkedlist

import (
	"fmt"
	"strings"
)

// NODE
type SNode struct {
	data int
	next *SNode
}

func newSNode(data int) *SNode {
	return &SNode{
		data: data,
		next: nil,
	}
}

func (node SNode) String() string {
	return fmt.Sprintf("%d -> ", node.data)
}

// LIST
type SLinkedList struct {
	head *SNode
}

func (list SLinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list SLinkedList) Length() int {
	len := 0
	tmp := list.head
	for ; tmp != nil; tmp = tmp.next {
		len++
	}
	return len
}

func (list SLinkedList) String() string {
	builder := strings.Builder{}
	builder.WriteString("head -> ")

	for tmp := list.head; tmp != nil; tmp = tmp.next {
		builder.WriteString(tmp.String())
	}

	builder.WriteString("nil")
	return builder.String()
}

// FIND
func (list SLinkedList) First() *SNode {
	return list.head
}

func (list SLinkedList) GetAt(index int) *SNode {
	i := -1
	tmp := list.head
	for {
		if tmp != nil {
			i++
		} else {
			break
		}

		if i == index {
			return tmp
		}
		tmp = tmp.next
	}
	return nil
}

func (list SLinkedList) GetWhere(val int) *SNode {
	tmp := list.head

	for ; tmp != nil; tmp = tmp.next {
		if tmp.data == val {
			return tmp
		}
	}

	return nil
}

func (list SLinkedList) Last() *SNode {
	tmp := list.head
	for ; tmp != nil && tmp.next != nil; tmp = tmp.next {
		continue
	}
	return tmp
}

// ADD
func (list *SLinkedList) Init(data int) bool {
	if !list.IsEmpty() {
		return false
	}

	tmp := newSNode(data)
	list.head = tmp
	return true
}

func (list *SLinkedList) AddHead(data int) *SNode {
	if list.IsEmpty() {
		list.Init(data)
	} else {
		tmp := newSNode(data)
		tmp.next = list.head
		list.head = tmp
	}

	return list.head
}

func (list *SLinkedList) AddTail(data int) *SNode {
	if list.IsEmpty() {
		list.Init(data)
		return list.head
	}

	last := list.Last()
	tmp := newSNode(data)
	last.next = tmp
	return tmp
}

func (list *SLinkedList) AddAt(index, data int) *SNode {
	target := list.GetAt(index)

	if target == nil {
		return nil
	}

	tmp := newSNode(data)
	tmp.next = target.next
	target.next = tmp
	return tmp
}

// REMOVE
func (list *SLinkedList) RemoveHead() *SNode {
	if list.IsEmpty() {
		return nil
	}

	tmp := list.head
	list.head = tmp.next
	tmp.next = nil
	return tmp
}

func (list *SLinkedList) RemoveTail() *SNode {
	var prev, cur *SNode = nil, list.head

	for cur != nil && cur.next != nil {
		prev = cur
		cur = cur.next
	}

	if cur == nil { // there are no last element
		return nil
	} else { // there is a last element
		if prev == nil { // last is also head
			list.head = nil
		} else {
			prev.next = nil
		}
		return cur
	}
}

func (list *SLinkedList) RemoveAt(index int) *SNode {
	var prev, cur *SNode = nil, list.head

	i := -1
	for {
		if cur != nil {
			i++
		} else {
			break
		}

		if i == index { // found the ith element
			if prev == nil {
				list.head = cur.next
			} else {
				prev.next = cur.next
			}
			cur.next = nil
			return cur
		}

		prev = cur
		cur = cur.next
	}

	return nil
}

func (list *SLinkedList) RemoveWhere(val int) *SNode {
	var prev, cur *SNode = nil, list.head
	for cur != nil {
		if cur.data == val {
			if prev == nil { // cur is head
				list.head = cur.next
			} else {
				prev.next = cur.next
			}
			cur.next = nil

			return cur
		}

		prev = cur
		cur = cur.next
	}
	return nil
}

func (list *SLinkedList) Clear() {
	list.head = nil
}

// REVERSE
func (list *SLinkedList) Reverse() {
	var prev, cur, next *SNode = nil, list.head, nil
	for cur != nil {
		next = cur.next

		cur.next = prev
		prev, cur = cur, next
	}

	list.head = prev
}
