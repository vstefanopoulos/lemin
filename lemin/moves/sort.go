package moves

// Sort moves to expected format for printing
func Sort() [][]string {
	lastPath := len(MovesPerAnt) - 1
	output := make([][]string, len(MovesPerAnt[lastPath]))
	for _, ant := range MovesPerAnt {
		for i, step := range ant {
			if step != "" {
				output[i] = append(output[i], step)
			}
		}
	}
	return output
}
