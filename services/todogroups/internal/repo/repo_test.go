package repo

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	todoGroupsRepo = NewMockGroupsRepo()
	todoGroupPermitsRepo = NewMockPermitsRepo()

	retCode := m.Run()
	os.Exit(retCode)
}
