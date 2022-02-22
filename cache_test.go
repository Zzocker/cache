package cache

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheNew(t *testing.T) {
	is := assert.New(t)

	l := New(10)
	is.Equal(10, l.capacity)
	is.Equal(0, l.size)
	is.Equal(&l.queue, l.queue.next)
	is.Equal(&l.queue, l.queue.prev)
	is.Empty(l.mp)
}

func TestCacheGetCapacity(t *testing.T) {
	is := assert.New(t)

	l := New(10)

	is.Equal(10, l.GetCapacity())
}

func TestCacheGetSet(t *testing.T) {
	is := assert.New(t)

	l := New(3)
	// invalid key
	is.Equal(ErrInvalidKey, l.Set(nil, nil))

	is.NoError(l.Set("one", tmpData{data: 1}))
	is.NoError(l.Set("two", tmpData{data: 2}))
	is.NoError(l.Set("three", tmpData{data: 3}))
	// one should be kicked out
	is.NoError(l.Set("four", tmpData{data: 4}))
	is.Nil(l.Get("one"))

	is.NoError(l.Set("two", tmpData{data: 99}))
	is.Equal(99, l.Get("two"))

	is.NoError(l.Set("five", tmpData{data: 5}))
	// three should be kicked out
	is.Nil(l.Get("three"))

	// get invalid key
	is.Nil(l.Get(nil))

}

func TestCacheSetCapacity(t *testing.T) {
	is := assert.New(t)

	l := New(10)

	for i := 0; i < 10; i++ {
		is.NoError(l.Set(fmt.Sprintf("%d", i), tmpData{data: i}))
	}
	// 0, 1, 2, 3 should be removed
	l.SetCapacity(6)
	for i := 0; i < 4; i++ {
		is.Nil(l.Get(fmt.Sprintf("%d", i)))
	}
}

func benchmarkSet(b *testing.B, cache *LRUCache, count int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < count; j++ {
			cache.Set(j, tmpData{data: j})
		}
	}
}

func benchmarkGet(b *testing.B, cache *LRUCache, count int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < count; j++ {
			cache.Get(j)
		}
	}
}

func BenchmarkSet10(b *testing.B) {
	n := 10
	b.StopTimer()
	cache := New(n)
	b.StartTimer()
	benchmarkSet(b, cache, n)
}

func BenchmarkSet100(b *testing.B) {
	n := 100
	b.StopTimer()
	cache := New(n)
	b.StartTimer()
	benchmarkSet(b, cache, n)
}

func BenchmarkSet1000(b *testing.B) {
	n := 1000
	b.StopTimer()
	cache := New(n)
	b.StartTimer()
	benchmarkSet(b, cache, n)
}

func BenchmarkSet10000(b *testing.B) {
	n := 10000
	b.StopTimer()
	cache := New(n)
	b.StartTimer()
	benchmarkSet(b, cache, n)
}

func BenchmarkSet100000(b *testing.B) {
	n := 100000
	b.StopTimer()
	cache := New(n)
	b.StartTimer()
	benchmarkSet(b, cache, n)
}

func BenchmarkSet1000000(b *testing.B) {
	n := 1000000
	b.StopTimer()
	cache := New(n)
	b.StartTimer()
	benchmarkSet(b, cache, n)
}

func BenchmarkGet10(b *testing.B) {
	n := 10
	b.StopTimer()
	cache := New(n)
	for i := 0; i < n; i++ {
		cache.Set(i, tmpData{data: i})
	}
	b.StartTimer()
	benchmarkGet(b, cache, n)
}

func BenchmarkGet100(b *testing.B) {
	n := 100
	b.StopTimer()
	cache := New(n)
	for i := 0; i < n; i++ {
		cache.Set(i, tmpData{data: i})
	}
	b.StartTimer()
	benchmarkGet(b, cache, n)
}

func BenchmarkGet1000(b *testing.B) {
	n := 1000
	b.StopTimer()
	cache := New(n)
	for i := 0; i < n; i++ {
		cache.Set(i, tmpData{data: i})
	}
	b.StartTimer()
	benchmarkGet(b, cache, n)
}

func BenchmarkGet10000(b *testing.B) {
	n := 10000
	b.StopTimer()
	cache := New(n)
	for i := 0; i < n; i++ {
		cache.Set(i, tmpData{data: i})
	}
	b.StartTimer()
	benchmarkGet(b, cache, n)
}

func BenchmarkGet100000(b *testing.B) {
	n := 100000
	b.StopTimer()
	cache := New(n)
	for i := 0; i < n; i++ {
		cache.Set(i, tmpData{data: i})
	}
	b.StartTimer()
	benchmarkGet(b, cache, n)
}

func BenchmarkGet1000000(b *testing.B) {
	n := 1000000
	b.StopTimer()
	cache := New(n)
	for i := 0; i < n; i++ {
		cache.Set(i, tmpData{data: i})
	}
	b.StartTimer()
	benchmarkGet(b, cache, n)
}
