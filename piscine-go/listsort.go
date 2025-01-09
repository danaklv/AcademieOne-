package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	head := l
	if head == nil || head.Next == nil {
		return head
	}
	sorted := &NodeI{Data: -1}
	current := head
	for current != nil {
		next := current.Next
		prev := sorted
		for prev.Next != nil && prev.Next.Data < current.Data {
			prev = prev.Next
		}

		current.Next = prev.Next
		prev.Next = current

		current = next
	}
	return sorted.Next
}
