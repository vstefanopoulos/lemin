package paths

import "lemin/lemin/colony"

// Iterates over group of paths and returns true if any has common room with path
func HasConflict(paths [][]*colony.Room, path []*colony.Room) bool {
	for _, existing := range paths {
		if !IsCompatible(existing, path) {
			return true
		}
	}
	return false
}

// False if paths have common rooms. Excludes start and end rooms
func IsCompatible(path1, path2 []*colony.Room) bool {
	visitedRooms := make(map[*colony.Room]bool)
	if len(path2) == 2 && len(path1) == 2 {
		return false
	}
	for _, room := range path1[1 : len(path1)-1] {
		visitedRooms[room] = true
	}
	for _, room := range path2[1 : len(path2)-1] {
		if visitedRooms[room] {
			return false
		}
	}
	return true
}
