package memdb

import (
	"testing"

	"github.com/hackbutteers/groupstore/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
