package linked_list

type Node struct {
	Value string
	next  *Node
	prev  *Node
}
type LinkedList struct {
	head *Node
	tail *Node
}

func NewNode(value string) *Node {
	return &Node{
		Value: value,
		next:  nil,
		prev:  nil,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *LinkedList) First() *Node {
	return l.head
}

func (l *LinkedList) Last() *Node {
	return l.tail
}

func (l *LinkedList) Push(value string) {
	node := NewNode(value)
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}
