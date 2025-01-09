package utils

import "strings"

func DistributeAnts(antsCount int, paths [][]string) map[string][]int {
	antsPerPath := make(map[string]int)
	antsInPaths := make(map[string][]int)
	for antID := 1; antID <= antsCount; antID++ {
		var chosenPathName string
		var minCost float64 = float64(len(paths[len(paths)-1])) + float64(antsCount)
		for pathIndex := 0; pathIndex < len(paths); pathIndex++ {
			path := paths[pathIndex]
			pathName := strings.Join(path, " -> ")
			roomsInPath := len(path) - 1
			antsInCurrentPath := antsPerPath[pathName]
			cost := float64(roomsInPath + antsInCurrentPath)
			if cost < minCost {
				minCost = cost
				chosenPathName = pathName
			}
		}
		antsPerPath[chosenPathName]++
		antsInPaths[chosenPathName] = append(antsInPaths[chosenPathName], antID)
	}

	return antsInPaths
}
