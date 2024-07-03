package servers
import "fmt"

type Node struct {
	prev *Node
	next *Node
	val string
}

func (n *Node) String() string {
	return n.val
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
	l.tail.next = newNode(l.tail, nil, val)
	l.tail = l.tail.next
}

func (l *NodeList) Pop() *Node {
	if l.head == nil {
		return nil
	}
	tmp := l.head
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	return tmp
}

func (l *NodeList) Remove(val string) string {
	if l.head == nil {
		return ""
	}
	var (
		prev *Node
		cur = l.head
	)
	for cur != nil {
		if cur.val == val {
			if prev != nil {
				prev.next = cur.next
			}
			if cur.next != nil {
				cur.next.prev = prev
			}
			return val
		}
		prev = cur
		cur = cur.next
	}
	return ""
}

func (l *NodeList) String() string {
	return fmt.Sprintf("NodeList{Head:%s, Tail:%s}", l.head, l.tail)
}
