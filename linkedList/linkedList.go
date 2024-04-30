package linkedlist

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Data int   `json:"data"`
	Next *Node `json:"Next"`
}

func (node *Node) NodeToJson() []byte {
	json, _ := json.Marshal(node)
	return json
}

type LinkedList struct {
	Head   *Node `json:"head"`
	Length int   `json:"length"`
}

func NewLinkedList(value int) LinkedList {
	return LinkedList{
		Head: &Node{
			Data: value,
			Next: nil,
		},
		Length: 1}
}

func (l *LinkedList) AddToEnd(value int) *Node {
	newNodo := &Node{Data: value, Next: nil}
	if l.Head == nil {
		l.Head = newNodo
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNodo
	}

	l.Length++
	return newNodo
}

func (l *LinkedList) AddToBeginning(value int) *Node {
	newNodo := &Node{Data: value, Next: nil}
	if l.Head == nil {
		l.Head = newNodo
	} else {
		aux := l.Head
		l.Head = newNodo
		l.Head.Next = aux
	}
	l.Length++
	return newNodo
}

func (l *LinkedList) TheListToSLice() *[]int {
	s := make([]int, 0)
	if l.Head == nil {
		return &s
	}
	current := l.Head
	for current != nil {
		s = append(s, current.Data)
		current = current.Next
	}
	return &s
}

func (l *LinkedList) ShowLinkedList() {
	fmt.Println()
	if l.Head != nil {
		current := l.Head
		for current != nil {
			fmt.Print(current.Data, " -> ")
			current = current.Next
		}
	}
}
