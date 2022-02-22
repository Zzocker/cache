package cache

import (
	"errors"
	"sync"
)

var ErrInvalidKey = errors.New("invalid key")

// LRUCache: a thread safe LRU cache.
type LRUCache struct {
	capacity int
	size     int
	// mapping of key to linked list node
	mp map[interface{}]*node

	// queue implementation using circular doubly linked list
	queue node

	// mutual exclusive lock for keeping cache thread safe
	mux sync.Mutex
}

// New: Create a new thread safe lru cache with provided capacity.
func New(capacity int) *LRUCache {
	l := &LRUCache{capacity: capacity}
	l.reset()

	return l
}

// GetCapacity: get current capacity of the lru cache.
func (l *LRUCache) GetCapacity() int {
	l.mux.Lock()
	defer l.mux.Unlock()

	return l.capacity
}

// SetCapacity: set a new capacity to the cache.
func (l *LRUCache) SetCapacity(newCap int) {
	var toFree []DataHandler
	l.mux.Lock()
	for l.size > newCap {
		toFree = append(toFree, l.queue.prev.data)
		delete(l.mp, l.queue.prev.key)
		l.queue.prev.remove()
		l.size--
	}
	l.capacity = newCap
	l.mux.Unlock()

	for i := range toFree {
		toFree[i].Free()
	}
}

// Set : put a key with data in lru cache.
func (l *LRUCache) Set(key interface{}, data DataHandler) error {
	if !isKeyTypeValid(key) {
		return ErrInvalidKey
	}
	var toFree DataHandler
	l.mux.Lock()
	if l.mp[key] != nil {
		l.mp[key].data = data
		l.mp[key].remove()
	} else {
		l.mp[key] = &node{
			data: data,
			key:  key,
		}
		l.size++
		if l.size > l.capacity {
			toFree = l.queue.prev.data
			delete(l.mp, l.queue.prev.key)
			l.queue.prev.remove()
			l.size--
		}
	}
	l.mp[key].insert(&l.queue)
	l.mux.Unlock()
	if toFree != nil {
		toFree.Free()
	}

	return nil
}

// Get : return data from cache with provided key.
func (l *LRUCache) Get(key interface{}) interface{} {
	if !isKeyTypeValid(key) {
		return nil
	}
	l.mux.Lock()
	defer l.mux.Unlock()
	if l.mp[key] == nil {
		return nil
	}
	l.mp[key].remove()
	l.mp[key].insert(&l.queue)

	return l.mp[key].data.GetData()
}

func (l *LRUCache) reset() {
	l.size = 0
	l.mp = map[interface{}]*node{}
	// empty queue
	l.queue.next = &l.queue
	l.queue.prev = &l.queue
}
