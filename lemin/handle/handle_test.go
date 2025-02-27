package handle

import (
	"lemin/lemin/colony"
	commonErrors "lemin/lemin/common/errors"
	"os"
	"testing"
)

func TestSetUpColony(t *testing.T) {
	origDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(origDir)

	if err := os.Chdir("../../"); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// ---------------------------------------
	// Test 1 (correct)
	// ---------------------------------------

	fd := &FileData{
		FileName: "example01.txt",
	}
	col := &colony.Colony{}
	err = fd.SetUpColony(col)

	if err != nil {
		t.Errorf("Expected nil. Got %v", err)
	}

	// ---------------------------------------
	// Test 2 (Bad 00)
	// ---------------------------------------

	fd = &FileData{
		FileName: "badexample00.txt",
	}
	col = &colony.Colony{}
	err = fd.SetUpColony(col)

	if err != commonErrors.ErrInvalidAntsInput {
		t.Errorf("Expected: %v. Got %v", commonErrors.ErrInvalidAntsInput, err)
	}

	// ---------------------------------------
	// Test 3 (Bad 01)
	// ---------------------------------------
	fd = &FileData{
		FileName: "badexample01.txt",
	}
	col = &colony.Colony{}
	err = fd.SetUpColony(col)

	expectedMsg := "tunnel already exists from 8 to 7"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message: %q. Got: %q", expectedMsg, err.Error())
	}

	// ---------------------------------------
	// Test 4 (Bad 02)
	// ---------------------------------------
	fd = &FileData{
		FileName: "badexample02.txt",
	}
	col = &colony.Colony{}
	err = fd.SetUpColony(col)

	if err != commonErrors.ErrMissingEndRoom {
		t.Errorf("Expected error message: %q. Got: %q", commonErrors.ErrMissingEndRoom, err.Error())
	}
}
