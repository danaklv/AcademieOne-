package main

import "github.com/01-edu/z01"

type Door struct {
	state int
}

const (
	OPEN  = 0
	CLOSE = 1
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	PrintStr("\n")
	ptrDoor.state = OPEN
	return true
}

func IsDoorOpen(Door Door) bool {
	PrintStr("is the Door opened ?")
	PrintStr("\n")
	return Door.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	PrintStr("\n")
	return ptrDoor.state == CLOSE
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	PrintStr("\n")
	ptrDoor.state = CLOSE
	return true
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
