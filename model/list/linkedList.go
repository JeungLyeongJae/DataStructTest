package list

import (
	"fmt"
	"log"
)

// LinkedListNode 单链表 单节点
type LinkedListNode struct {
	Value any
	Next  *LinkedListNode
}

// LinkedList 单链表
type LinkedList struct {
	head *LinkedListNode
}

func (l *LinkedList) Insert(value any) {
	currentNode := l.head

	if currentNode == nil {
		currentNode = &LinkedListNode{value, nil}
		l.head = currentNode
	} else {
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = &LinkedListNode{value, nil}
	}
}

func (l *LinkedList) Delete() {
	l.head = nil
}

func (l *LinkedList) Contains(value any) bool {
	currentNode := l.head
	if currentNode == nil {
		return false
	} else {
		if currentNode.Value == value {
			return true
		} else {
			for currentNode.Next != nil {
				currentNode = currentNode.Next
				if currentNode.Value == value {
					return true
				}
			}
		}
	}
	return false
}

func (l *LinkedList) Remove(value any) {
	currentNode := l.head
	if currentNode == nil {
		return
	} else {
		if currentNode.Value == value {
			l.head = currentNode.Next
		} else {
			for currentNode.Next != nil {
				if currentNode.Next.Value == value {
					currentNode.Next = currentNode.Next.Next
					break
				} else {
					currentNode = currentNode.Next
				}
			}
		}
	}
}

func (l *LinkedList) RemoveAt(index int) {
	currentCount := 0
	currentNode := l.head
	if currentNode == nil {
		return
	} else {
		if currentCount == index {
			l.head = currentNode.Next
		} else {
			for currentNode.Next != nil {
				currentCount += 1
				if currentCount == index {
					currentNode.Next = currentNode.Next.Next
					break
				} else {
					currentNode = currentNode.Next
				}
			}
		}
	}
}

func (l *LinkedList) Get(index int) any {
	currentCount := 0
	currentNode := l.head
	if currentNode == nil {
		return nil
	} else {
		if currentCount == index {
			return currentNode.Value
		} else {
			for currentNode.Next != nil {
				currentCount += 1
				currentNode = currentNode.Next
				if currentCount == index {
					return currentNode.Value
				}
			}
		}
	}
	return nil
}

func (l *LinkedList) SetAt(index int, value any) {
	currentCount := 0
	currentNode := l.head
	if currentNode == nil {
		if index == 0 {
			currentNode = &LinkedListNode{value, nil}
			l.head = currentNode
			return
		} else {
			log.Panic("list is empty!")
			return
		}
	}
	for currentNode.Next != nil {
		currentCount += 1
		currentNode = currentNode.Next
		if currentCount == index {
			currentNode.Value = value
			return
		}
	}
	currentCount += 1
	if currentNode.Next == nil && currentCount == index {
		currentNode.Next = &LinkedListNode{value, nil}
		return
	}
	if currentCount < index {
		log.Panic("index is out of range!")
		return
	}
}

func (l *LinkedList) IndexOf(value any) (int, error) {
	currentCount := 0
	currentNode := l.head
	if currentNode == nil {
		return -1, fmt.Errorf("list is empty")
	}
	for currentNode.Next != nil {
		currentCount += 1
		currentNode = currentNode.Next
		if currentNode.Value == value {
			return currentCount, nil
		}
	}
	return -1, fmt.Errorf("list does not contain value: %v", value)
}

func (l *LinkedList) Print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Print(currentNode.Value, " -> ")
		currentNode = currentNode.Next
	}
	fmt.Print("nil\n")
}
