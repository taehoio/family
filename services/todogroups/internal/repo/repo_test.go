package repo

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	todoGroupsRepo = NewMockTodoGroupsRepo()
	todoGroupPermitsRepo = NewMockTodoGroupPermitsRepo()

	retCode := m.Run()
	os.Exit(retCode)
}
