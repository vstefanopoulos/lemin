package paths

import "lemin/lemin/colony"

var (
	Groups    map[int][][]*colony.Room
	BestGroup [][]*colony.Room
)

// Creates a map key with each room of allPaths as first index in value slice. Then adds non coflicting rooms in each value
func Group(allPaths [][]*colony.Room) {
	Groups = make(map[int][][]*colony.Room)
	for i, path := range allPaths {
		Groups[i] = append(Groups[i], path)
	}

	for _, path := range allPaths {
		for i, v := range Groups {
			if !HasConflict(v, path) {
				Groups[i] = append(Groups[i], path)
			}
		}
	}
}
