package utils

import (
	"fmt"
	"strings"
)

func SimulateAntMovement(antsInPaths map[string][]int) {
	antPositions := make(map[int]string)
	roomOccupancy := make(map[string]int)
	pathsCount := 0
	testAnt := 0

	for path, ants := range antsInPaths {
		pathsCount++
		pathRooms := strings.Split(path, " -> ")
		for _, ant := range ants {
			startRoom := pathRooms[0]
			antPositions[ant] = startRoom
			roomOccupancy[startRoom]++
		}
	}
	for {
		movementInfo := make([]string, 0)
		anyMovement := false
		for path, ants := range antsInPaths {

			pathRooms := strings.Split(path, " -> ")
			if len(pathRooms) == 2 {
				if testAnt < ants[len(ants)-1] {
					endRoom := pathRooms[1]
					testAnt++
					movementInfo = append(movementInfo, fmt.Sprintf("L%d-%s", testAnt, endRoom))
					if pathsCount == 1 {
						anyMovement = true
					}
				}
			} else {
				for _, ant := range ants {
					currentRoom := antPositions[ant]
					roomIndex := -1
					for i, room := range pathRooms {
						if room == currentRoom {
							roomIndex = i
							break
						}
					}
					if roomIndex < len(pathRooms)-1 {
						nextRoom := pathRooms[roomIndex+1]
						if roomOccupancy[nextRoom] == 0 || nextRoom == pathRooms[0] || nextRoom == pathRooms[len(pathRooms)-1] {
							antPositions[ant] = nextRoom
							roomOccupancy[currentRoom]--
							roomOccupancy[nextRoom]++
							movementInfo = append(movementInfo, fmt.Sprintf("L%d-%s", ant, nextRoom))
							anyMovement = true
						}
					}
				}
			}
		}
		if !anyMovement {
			break
		}
		if len(movementInfo) > 0 {
			fmt.Println(strings.Join(movementInfo, " "))
		}
	}

}
