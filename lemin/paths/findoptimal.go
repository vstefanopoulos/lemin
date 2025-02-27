package paths

// Returns the path group with the most paths and longenst path length beeing equal or smaller that ants number
func FindOptimal(ants int) {
	if len(Groups) == 0 {
		return
	}

	bestIndex := 0
	for i, group := range Groups {
		if len(group) > len(Groups[bestIndex]) &&
			len(group[len(group)-1]) <= ants {
			bestIndex = i
		}
	}
	BestGroup = Groups[bestIndex]
}
