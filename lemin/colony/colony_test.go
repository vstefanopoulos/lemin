package colony

import (
	commonErrors "lemin/lemin/common/errors"
	"testing"
)

func TestAddRoom(t *testing.T) {
	// -------------------------
	// Test 1 (valid)
	// -------------------------

	c := &Colony{}
	err := c.AddRoom("Test 0 5")
	if err != nil {
		t.Errorf("Expected nil. Got %q", err)
	}

	// -------------------------
	// Test 2 (invalid)
	// -------------------------

	c = &Colony{}
	err = c.AddRoom("Test 0")
	if err != commonErrors.ErrInvalidRoomInput {
		t.Errorf("Expected %q. Got %v", commonErrors.ErrInvalidRoomInput, err)
	}

	// -------------------------
	// Test 3 (invalid)
	// -------------------------

	c = &Colony{}
	err = c.AddRoom("Test err 5")
	expected := "coordinates error X: strconv.Atoi: parsing \"err\": invalid syntax"
	if err.Error() != expected {
		t.Errorf("Expected %q. Got %v", expected, err)
	}
}
