package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

func (l linkedList) Len() int {
	return l.length
}

func (l linkedList) Display() {
	for l.head != nil {
		fmt.Printf("%v -> ", l.head.data)
		l.head = l.head.next
	}
	fmt.Println()
}

func (l *linkedList) PushBack(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		l.tail = n
		l.length++
	}
}

func (l *linkedList) Delete(key int) {
	if l.head.data == key {
		l.head = l.head.next
		l.length--
		return
	}
	var prev *node = nil
	curr := l.head
	for curr != nil && curr.data != key {
		prev = curr
		curr = curr.next
	}
	if curr == nil {
		fmt.Println("key not found")
		return
	}
	prev.next = curr.next
	l.length--
	fmt.Println("Node Deleted")
}

func (l linkedList) Front() (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("Cannot Find Front value in an Empty linked list")
	}
	return l.head.data, nil
}

func (l linkedList) Back() (int, error) {
	if l.tail == nil {
		return 0, fmt.Errorf("Cannot Find Back value in an Empty linked list")
	}
	return l.tail.data, nil
}

func (l linkedList) Reverse() {
	curr := l.head
	l.tail = l.head
	var prev *node
	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}
	l.head = prev
}

func main() {
	list := linkedList{}
	node1 := &node{data: 60}
	node2 := &node{data: 30}
	node3 := &node{data: 90}
	list.PushBack(node1)
	list.PushBack(node2)
	list.PushBack(node3)
	fmt.Println("Length = ", list.Len())
	list.Display()
	//list.Delete(30)
	fmt.Println("Length = ", list.Len())
	list.Display()
	// list.Reverse()
	// list.Display()
}
