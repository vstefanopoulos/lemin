package colony

// If Rooms field name in col matches name arg returns the ptr of Room. If not returns empty string
func (c *Colony) GetRoom(name string) *Room {
	for _, r := range c.Rooms {
		if name == r.Name {
			return r
		}
	}
	return nil
}
