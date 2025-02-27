package moves

import "lemin/lemin/colony"

// Contains (excluding start including end):
//   - number of rooms for each path
//   - room names in order from start to end
type pathData struct {
	rating int
	path   []*colony.Room
}

// Creates a slice of type pathData for each path
func createPD(paths [][]*colony.Room) (pathList []pathData) {
	for _, p := range paths {
		p = p[1:]
		pathList = append(pathList, pathData{len(p), p})
	}
	return
}
