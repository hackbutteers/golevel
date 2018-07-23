
package db

import (
	. "github.com/onsi/gomega"

	"github.com/hackbutteers/groupstore/iterator"
	"github.com/hackbutteers/groupstore/opt"
	"github.com/hackbutteers/groupstore/testutil"
	"github.com/hackbutteers/groupstore/util"
)

type testingDB struct {
	*DB
	ro   *opt.ReadOptions
	wo   *opt.WriteOptions
	stor *testutil.Storage
}

func (t *testingDB) TestPut(key []byte, value []byte) error {
	return t.Put(key, value, t.wo)
}

func (t *testingDB) TestDelete(key []byte) error {
	return t.Delete(key, t.wo)
}

func (t *testingDB) TestGet(key []byte) (value []byte, err error) {
	return t.Get(key, t.ro)
}

func (t *testingDB) TestHas(key []byte) (ret bool, err error) {
	return t.Has(key, t.ro)
}

func (t *testingDB) TestNewIterator(slice *util.Range) iterator.Iterator {
	return t.NewIterator(slice, t.ro)
}

func (t *testingDB) TestClose() {
	err := t.Close()
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	err = t.stor.Close()
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
}

func newTestingDB(o *opt.Options, ro *opt.ReadOptions, wo *opt.WriteOptions) *testingDB {
	stor := testutil.NewStorage()
	db, err := Open(stor, o)
	// FIXME: This may be called from outside It, which may cause panic.
	Expect(err).NotTo(HaveOccurred())
	return &testingDB{
		DB:   db,
		ro:   ro,
		wo:   wo,
		stor: stor,
	}
}

type testingTransaction struct {
	*Transaction
	ro *opt.ReadOptions
	wo *opt.WriteOptions
}

func (t *testingTransaction) TestPut(key []byte, value []byte) error {
	return t.Put(key, value, t.wo)
}

func (t *testingTransaction) TestDelete(key []byte) error {
	return t.Delete(key, t.wo)
}

func (t *testingTransaction) TestGet(key []byte) (value []byte, err error) {
	return t.Get(key, t.ro)
}

func (t *testingTransaction) TestHas(key []byte) (ret bool, err error) {
	return t.Has(key, t.ro)
}

func (t *testingTransaction) TestNewIterator(slice *util.Range) iterator.Iterator {
	return t.NewIterator(slice, t.ro)
}

func (t *testingTransaction) TestClose() {}