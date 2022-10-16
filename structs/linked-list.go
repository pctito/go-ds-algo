package structs

import "fmt"

type node struct {
	next *node
	val  int
}

type Linked struct {
	head *node
	len  int
}

func (l *Linked) PrependInt(value int) {
	newNode := node{val: value}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.len++
	} else {
		l.head = &newNode
		l.len++
	}
}

func (l *Linked) PrintLs() {
	if l.head == nil {
		return
	}
	cnode := l.head
	for cnode != nil {
		fmt.Println(cnode.val)
		cnode = cnode.next
	}
}
