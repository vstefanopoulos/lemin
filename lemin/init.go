package lemin

import (
	"fmt"
	"lemin/lemin/colony"
	commonErrors "lemin/lemin/common/errors"
	"lemin/lemin/handle"
	"lemin/lemin/moves"
	"lemin/lemin/paths"
	"os"
	"strings"
)

func Please() (string, [][]string, error) {
	args := os.Args
	if len(args) != 2 {
		return "", nil, commonErrors.ErrInvalidArguments
	}

	fd := &handle.FileData{}
	col := &colony.Colony{}

	fd.FileName = args[1]
	err := fd.SetUpColony(col)
	if err != nil {
		return "", nil, err
	}

	err = col.FindAllPaths()
	if err != nil {
		return "", nil, err
	}

	paths.Group(*col.AllPaths)
	paths.FindOptimal(fd.Ants)

	moves.Make(fd.Ants, paths.BestGroup)
	moves := moves.Sort()

	return fd.InputStr, moves, nil
}

func PrintResult(input string, output [][]string) {
	fmt.Println(input + "\n")

	for i, v := range output {
		fmt.Println(strings.Join(v, " "))
		if i == len(output)-1 {
			fmt.Printf("\nTotal turns: %d\n", len(output))
		}
	}
}
