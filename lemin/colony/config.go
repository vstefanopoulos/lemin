package colony

type Colony struct {
	Rooms    []*Room
	Start    *Room
	End      *Room
	AllPaths *[][]*Room
}

// Holds Room name, coordinates and connections
type Room struct {
	Name          string
	X             int
	Y             int
	ConnectedList []*Room
}
