package cache

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkLRUCache(b *testing.B) {
	c := NewCache(NewLRU(10000))

	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		for pb.Next() {
			key := uint64(r.Intn(1000000))
			c.Get(0, key, func() (int, Value) {
				return 1, key
			}).Release()
		}
	})
}

