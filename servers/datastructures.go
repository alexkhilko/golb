package servers

import "fmt"

type Node struct {
	Prev *Node
	Next *Node
	Val  string
}

func (n *Node) Remove() {
	if n.Prev != nil {
		n.Prev.Next = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	}
}

func (n *Node) String() string {
	return n.Val
}

func newNode(prev, next *Node, val string) *Node {
	return &Node{prev, next, val}
}

type NodeList struct {
	head *Node
	tail *Node
}

func (l *NodeList) Push(val string) {
	if l.head == nil {
		l.head = newNode(nil, nil, val)
		l.tail = l.head
		return
	}
	l.tail.Next = newNode(l.tail, nil, val)
	l.tail = l.tail.Next
}

func (l *NodeList) Pop() *Node {
	if l.head == nil {
		return nil
	}
	tmp := l.head
	l.head = l.head.Next
	if l.head == nil {
		l.tail = nil
	}
	return tmp
}

func (l *NodeList) Remove(n *Node) string {
	if n == l.head {
		l.head = n.Next
	}
	if n == l.tail {
		l.tail = n.Prev
	}
	n.Remove()
	return n.Val
}

func (l *NodeList) Top() *Node {
	return l.head
}

func (l *NodeList) ValuesList() []string {
	if l.head == nil {
		return []string{}
	}
	var (
		cur    = l.head
		result = []string{}
	)
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}

func (l *NodeList) String() string {
	return fmt.Sprintf("NodeList{Head:%s, Tail:%s}", l.head, l.tail)
}

func NewNodeList() *NodeList {
	return &NodeList{}
}
