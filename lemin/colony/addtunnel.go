package colony

import (
	"fmt"
	commonErrors "lemin/lemin/common/errors"
)

// Accepts two room name strings and adds roomA to roomB in connectedList field and vice versa
// Returns error if:
//   - any room is nil
//   - rooms are already connected
func (c *Colony) AddTunnel(roomA, roomB string) error {
	a := c.GetRoom(roomA)
	b := c.GetRoom(roomB)

	switch {
	case a == nil, b == nil:
		return fmt.Errorf("%w from %v to %v", commonErrors.ErrInvalidTunnel, roomA, roomB)
	case Contains(a.ConnectedList, roomB), Contains(b.ConnectedList, roomA):
		return fmt.Errorf("%w from %v to %v", commonErrors.ErrTunnelExists, roomA, roomB)

	default:
		a.ConnectedList = append(a.ConnectedList, b)
		b.ConnectedList = append(b.ConnectedList, a)
	}
	return nil
}

// Returns true if name matches a name field in any instance of Room contained in rooms
func Contains(rooms []*Room, name string) bool {
	for _, r := range rooms {
		if name == r.Name {
			return true
		}
	}
	return false
}
