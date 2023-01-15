package tree

type node struct {
	value     int
	rightNode *node
	leftNode  *node
}

type tree struct {
	root *node
}

func (t *tree) _insert(currNode *node, num int) *node {
	if currNode == nil {
		return &node{num, nil, nil}
	}

	if currNode.value < num {
		currNode.leftNode = t._insert(currNode.leftNode, num)
	}

	if currNode.value > num {
		currNode.rightNode = t._insert(currNode.rightNode, num)
	}

	return currNode
}

func (t *tree) insert(num int) {
	t._insert(t.root, num)
}

func (t *tree) _max(currNode *node) int {
	if currNode.rightNode == nil {
		return currNode.value
	}
	t_max := t._max(currNode.rightNode)
	return t_max
}

func (t *tree) max() int {
	return t._max(t.root)
}

func main() {
	root := &node{10, nil, nil}
	t := &tree{root}
	t.insert(1)
}
