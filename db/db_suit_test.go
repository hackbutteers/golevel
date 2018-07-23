package db

import (
	"testing"

	"github.com/hackbutteers/groupstore/testutil"
)

func TestDB(t *testing.T) {
	testutil.RunSuite(t, "db Suite")
}
