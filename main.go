package main

import linked_list "argo_exer/linked-list"

func main() {
	l := linked_list.NewLinkedList()
	l.Push("A")
	l.Push("B")

	n := l.First()

	for {
		println(n.Value)
		n = n.Next()
		if n == nil {
			break
		}
	}
}
