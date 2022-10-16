package structs

import "fmt"

type Node struct {
	next *Node
	val  int
}

type Linked struct {
	head *Node
	len  int
}

func (l *Linked) Prepend(value int) {
	newNode := Node{val: value}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.len++
	} else {
		l.head = &newNode
		l.len++
	}
}

func (l *Linked) AppendInt(value int) {
	return
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
