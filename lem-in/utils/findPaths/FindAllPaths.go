package utils

func FindAllPaths(current string, end string, visited map[string]bool, path []string, allPaths *[][]string, graph map[string][]string) {
	path = append(path, current)

	if current == end {
		*allPaths = append(*allPaths, append([]string{}, path...))
		return
	}
	visited[current] = true
	for _, neighbor := range graph[current] {
		if !visited[neighbor] {
			FindAllPaths(neighbor, end, visited, path, allPaths, graph)
		}
	}

	visited[current] = false
}
