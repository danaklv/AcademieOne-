package utils

func ChooseBestPaths(allPaths [][]string, antsCount int) [][]string {
	var bestPaths [][]string
	if antsCount == 1 {
		bestPaths = append([][]string(nil), allPaths[0])
		return bestPaths
	}
	visited := make(map[string]bool)
	FindPaths(allPaths, 0, visited, nil, &bestPaths)
	return bestPaths
}

func FindPaths(allPaths [][]string, index int, visited map[string]bool, currentPaths [][]string, bestPaths *[][]string) {
	if index >= len(allPaths) {
		if len(currentPaths) > len(*bestPaths) {
			*bestPaths = append([][]string(nil), currentPaths...)
		}
		return
	}
	currentPath := allPaths[index]
	hasConflict := false
	pathVisited := make(map[string]bool)
	for _, room := range currentPath[1 : len(currentPath)-1] {
		if visited[room] {
			hasConflict = true
			break
		}
		pathVisited[room] = true
	}

	if !hasConflict {
		for room := range pathVisited {
			visited[room] = true
		}

		currentPaths = append(currentPaths, currentPath)
		FindPaths(allPaths, index+1, visited, currentPaths, bestPaths)
		currentPaths = currentPaths[:len(currentPaths)-1]

		for room := range pathVisited {
			delete(visited, room)
		}
	}

	FindPaths(allPaths, index+1, visited, currentPaths, bestPaths)
}
