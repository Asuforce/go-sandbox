package linkedlist

import (
	"fmt"
	"sync"
)

type Node struct {
	content string
	next    *Node
}

type ItemLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

func (ll *ItemLinkedList) Append(t string) {
	ll.lock.Lock()
	node := Node{t, nil}
	if ll.head == nil {
		ll.head = &node
	} else {
		last := ll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	ll.size++
	ll.lock.Unlock()
}

func (ll *ItemLinkedList) Insert(i int, t string) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i > ll.size {
		return fmt.Errorf("Index out of bounds")
	}
	addNode := Node{t, nil}
	if i == 0 {
		addNode.next = ll.head
		ll.head = &addNode
		return nil
	}
	node := ll.head
	for j := 0; j < i-2; j++ {
		node = node.next
	}
	addNode.next = node.next
	node.next = &addNode
	ll.size++
	return nil
}

func (ll *ItemLinkedList) RemoveAt(i int) (*string, error) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i > ll.size {
		return nil, fmt.Errorf("Index out of bounds")
	}
	node := ll.head
	for j := 0; j < i-1; j++ {
		node = node.next
	}
	remove := node.next
	node.next = remove.next
	ll.size--
	return &remove.content, nil
}

func (ll *ItemLinkedList) IndexOf(t string) int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := ll.head
	for j := 0; ; j++ {
		if node.content == t {
			return j
		}
		if node.next == nil {
			return -1
		}
		node = node.next
	}
}

func (ll *ItemLinkedList) IsEmpty() bool {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	if ll.head == nil {
		return true
	}
	return false
}

func (ll *ItemLinkedList) Size() int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	size := 1
	last := ll.head
	for {
		if last == nil || last.next == nil {
			break
		}
		last = last.next
		size++
	}
	return size
}

func (ll *ItemLinkedList) String() {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := ll.head
	for j := 0; ; j++ {
		if node == nil {
			break
		}
		fmt.Print(node.content)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}

func (ll *ItemLinkedList) Head() *Node {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	return ll.head
}
