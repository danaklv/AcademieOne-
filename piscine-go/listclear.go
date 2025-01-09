package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListClear(l *List) {
	for l.Head != nil {
		l.Head = l.Head.Next
	}
	l.Tail = nil
}
