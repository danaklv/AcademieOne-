package utils

import (
	"fmt"
	utils "lm/utils/antsgo"
	utilss "lm/utils/findPaths"
	"sort"
	"strconv"
	"strings"
)

var graph map[string][]string

func CheckFormat(text string) {

	lines := strings.Split(text, "\n")
	var data [][]string
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "##") && line != "##start" && line != "##end" {
			continue
		}
		words := strings.Fields(line)
		data = append(data, words)
	}

	if len(data) == 0 || len(data[0]) == 0 {
		fmt.Println("ERROR: invalid data format")
		return

	}

	AntsCount, err := strconv.Atoi(data[0][0])
	if AntsCount <= 0 || err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	var relations [][]string
	var rooms []string
	hasStart := false
	hasEnd := false
	relationMap := make(map[string]bool)
	var startRoom string
	var endRoom string

	for i, line := range data {
		if line[0] == "##start" {
			if len(data[i+1]) == 3 {
				startRoom = data[i+1][0]
			} else {
				fmt.Println("ERROR: invalid data format")
				return
			}
			if hasStart {
				fmt.Println("ERROR: invalid data format")
				return
			} else {
				hasStart = true
			}
		}
		if line[0] == "##end" {
			if len(data[i+1]) == 3 {
				endRoom = data[i+1][0]
			} else {
				fmt.Println("ERROR: invalid data format")
				return
			}

			if hasEnd {
				fmt.Println("ERROR: invalid data format")
				return
			} else {
				hasEnd = true
			}
		}
		if len(line) == 1 && strings.Contains(line[0], "-") {
			relation := strings.Split(line[0], "-")
			if !Contains(rooms,relation[0]) ||!Contains(rooms,relation[1]) {
				fmt.Println("ERROR: invalid data format")
				return
			}
			relations = append(relations, line)
			if relationMap[relation[1]+"-"+relation[0]] || relation[0] == relation[1] {
				fmt.Println("ERROR: invalid data format")
				return
			}
			relationMap[relation[0]+"-"+relation[1]] = true
		}
		if len(line) == 3 {
			rooms = append(rooms, line[0])
		}
	}
	if hasDuplicates(rooms) {
		fmt.Println("ERROR: invalid data format")
		return
	}

	if len(rooms) == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	if len(relations) == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	if !hasEnd || !hasStart {
		fmt.Println("ERROR: invalid data format")
		return
	}
	

	graph = make(map[string][]string)
	for _, relation := range relations {
		roomsRelations := strings.Split(relation[0], "-")
		graph[roomsRelations[0]] = append(graph[roomsRelations[0]], roomsRelations[1])
		graph[roomsRelations[1]] = append(graph[roomsRelations[1]], roomsRelations[0])
	}
	var currentPath []string
	var allPaths [][]string

	utilss.FindAllPaths(startRoom, endRoom, make(map[string]bool), currentPath, &allPaths, graph)

	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	if len(allPaths) == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}

	bestPaths := utilss.ChooseBestPaths(allPaths, AntsCount)

	antsInPaths := utils.DistributeAnts(AntsCount, bestPaths)
	utils.SimulateAntMovement(antsInPaths)

}

func hasDuplicates(slice []string) bool {
	seen := make(map[string]bool)

	for _, element := range slice {
		if seen[element] {
			return true
		}
		seen[element] = true
	}
	return false
}

func Contains(array []string, element string) bool {
    for _, item := range array {
        if item == element {
            return true
        }
    }
    return false
}
