package piscine

func ListReverse(l *List) {
	var prev, current, next *NodeL
	current = l.Head
	l.Tail = l.Head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
