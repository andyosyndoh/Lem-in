package utils

import (
	"lem-in/models"
)

// FindPaths finds all possible paths from start to end using BFS
func FindPaths(colony *models.AntColony) []models.Path {
	var allPaths []models.Path
	var queue [][]string
	queue = append(queue, []string{colony.Start})

	for i := range colony.Rooms {
		colony.Rooms[i].IsVisited = false
	}

	for len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]
		currentRoom := currentPath[len(currentPath)-1]
		if currentRoom == colony.End {
			allPaths = append(allPaths, models.Path{Rooms: currentPath})
			continue
		}
		for _, nextRoom := range colony.Links[currentRoom] {
			if !containsRoom(currentPath, nextRoom) {
				// Create a new path with the next room added
				newPath := make([]string, len(currentPath))
				copy(newPath, currentPath)
				newPath = append(newPath, nextRoom)
				queue = append(queue, newPath)
			}
		}
	}

	return OptimizePaths(allPaths)
}

// Helper function to check if a room is in a path
func containsRoom(path []string, room string) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}

func OptimizePaths(paths []models.Path) []models.Path {
	optimized := make([]models.Path, 0)
	optimized = append(optimized, paths[0])
	for i := 1; i < len(paths); i++ {
		if Check(paths[i].Rooms, optimized) {
			optimized = append(optimized, paths[i])
		}
	}
	return optimized
}

func Check(path []string, optimized []models.Path) bool {
	for _, optpath := range optimized{
		for k := 1; k < len(path); k++ {
			if k < len(optpath.Rooms) {
				if path[k] == optpath.Rooms[k] {
					return false
				}
			}
		}
	}
	return true
}
