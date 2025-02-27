package handle

import (
	"lemin/lemin/colony"
	commonErrors "lemin/lemin/common/errors"
	"os"
	"strconv"
	"strings"
)

// Fills colony struct fields with file info. Returns an error if:
//   - Os.Args not 2
//   - Ants field is not a possitive int
//   - Start or End room is not present
//   - Start or End room is not connected
func (fd *FileData) SetUpColony(c *colony.Colony) error {

	fileName := "examples/" + fd.FileName
	file, fileErr := os.ReadFile(fileName)
	if fileErr != nil {
		return fileErr
	}

	fd.InputStr = string(file)
	lines := strings.Split(fd.InputStr, "\n")

	var antsErr error
	fd.Ants, antsErr = strconv.Atoi(lines[0])
	if antsErr != nil || fd.Ants < 1 {
		return commonErrors.ErrInvalidAntsInput
	}

	for i := 1; i < len(lines); i++ {
		switch {
		case lines[i] == "":
			continue
		case (lines[i][0] == '#' && len(lines[i]) == 1), (lines[i][0] == '#' && lines[i][1] != '#'):
			continue

		case lines[i] == "##start" && i < len(lines)-1:
			fd.Start = strings.Split(lines[i+1], " ")[0]

		case lines[i] == "##end" && i < len(lines)-1:
			fd.End = strings.Split(lines[i+1], " ")[0]

		case containsNumber(lines[i]) && !strings.Contains(lines[i], "-"):
			roomErr := c.AddRoom(lines[i])
			if roomErr != nil {
				return roomErr
			}
		case strings.Contains(lines[i], "-"):
			parts := strings.Split(lines[i], "-")
			tunnelErr := c.AddTunnel(parts[0], parts[1])
			if tunnelErr != nil {
				return tunnelErr
			}
		default:
			break
		}
	}

	err := setStartEnd(fd, c)
	if err != nil {
		return err
	}

	c.AllPaths = &[][]*colony.Room{}
	return nil
}
