package handle

import (
	"lemin/lemin/colony"
	commonErrors "lemin/lemin/common/errors"
	"strings"
)

func containsNumber(str string) bool {
	return strings.ContainsAny(str, "0123456789")
}

// Files Start and End room in Colony struct fields. Returns an error if:
//   - Start or End room is not present
//   - Start or End room is not connected
func setStartEnd(fd *FileData, c *colony.Colony) error {
	c.Start = c.GetRoom(fd.Start)
	if c.Start == nil {
		return commonErrors.ErrMissingStartRoom
	}

	if len(c.Start.ConnectedList) == 0 {
		return commonErrors.ErrStartRoomDisconnected
	}

	c.End = c.GetRoom(fd.End)
	if c.End == nil {
		return commonErrors.ErrMissingEndRoom
	}

	if len(c.End.ConnectedList) == 0 {
		return commonErrors.ErrEndRoomDisconnected
	}
	return nil
}
