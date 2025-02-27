package colony

import (
	"fmt"
	commonErrors "lemin/lemin/common/errors"
	"strconv"
	"strings"
)

// Adds room to colony with room name. Returns an error if:
//   - Input field is less that 3
//   - Room already exists
//   - x or y is not a valid int
func (c *Colony) AddRoom(input string) error {
	parts := strings.Fields(input)

	if len(parts) < 3 {
		return commonErrors.ErrInvalidRoomInput
	}

	name := parts[0]
	if Contains(c.Rooms, name) {
		return fmt.Errorf("%w %v", commonErrors.ErrRoomAlreadyExists, name)
	}

	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("%w X: %v", commonErrors.ErrCoord, err)
	}

	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("%w Y: %v", commonErrors.ErrCoord, err)
	}

	c.Rooms = append(c.Rooms, &Room{Name: name, X: x, Y: y})
	return nil
}
