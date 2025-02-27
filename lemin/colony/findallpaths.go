package colony

import (
	"lemin/lemin/common/errors"
	"sort"
)

// Registers all rooms connected with start and end in ascending order according to path length
func (c *Colony) FindAllPaths() error {
	var dfs func(current *Room, visited map[*Room]bool, path []*Room, paths *[][]*Room)

	dfs = func(current *Room, visited map[*Room]bool, path []*Room, paths *[][]*Room) {
		visited[current] = true
		path = append(path, current)

		if current == c.End {
			tmp := make([]*Room, len(path))
			copy(tmp, path)

			*c.AllPaths = append(*c.AllPaths, tmp)
		}

		for _, neighbour := range current.ConnectedList {
			if !visited[neighbour] {
				dfs(neighbour, visited, path, paths)
			}
		}
		delete(visited, current)
	}

	visited := make(map[*Room]bool)
	dfs(c.Start, visited, []*Room{}, c.AllPaths)
	if len(*c.AllPaths) == 0 {
		return errors.ErrStartEndNotConnected
	}
	sortRoutes(*c.AllPaths)
	return nil
}

func sortRoutes(allRoutes [][]*Room) {
	sort.Slice(allRoutes, func(i int, j int) bool {
		return len(allRoutes[i]) < len(allRoutes[j])
	})
}
