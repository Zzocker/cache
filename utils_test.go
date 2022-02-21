package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsKeyTypeValid(t *testing.T) {
	is := assert.New(t)
	// key is nil
	is.False(isKeyTypeValid(nil))

	// key is array
	is.False(isKeyTypeValid([]int{1, 2, 8}))

	// key is map
	is.False(isKeyTypeValid(map[string]string{"key": "value"}))

	// key is int
	is.True(isKeyTypeValid(1))

	// key is float
	is.True(isKeyTypeValid(12.3))

	// key is string
	is.True(isKeyTypeValid("this is string key"))

	// key is pointer
	k := "this is pointer key"
	is.False(isKeyTypeValid(&k))
}

type tmpData struct {
	data int
}

func (t tmpData) Free() {}
func (t tmpData) GetData() interface{} {
	return t.data
}

func TestNodeInsert(t *testing.T) {
	is := assert.New(t)

	one := &node{
		data: tmpData{1},
		prev: nil,
	}
	three := &node{
		data: tmpData{3},
		prev: one,
		next: nil,
	}
	one.next = three
	// one <--> three

	two := &node{
		data: tmpData{2},
	}
	two.insert(one)
	// it should be one <--> two <--> three
	is.Equal(2, one.next.data.GetData())
	is.Equal(2, three.prev.data.GetData())
}

func TestNodeRemove(t *testing.T) {
	is := assert.New(t)

	one := &node{
		data: tmpData{1},
	}
	two := &node{
		data: tmpData{2},
	}
	three := &node{
		data: tmpData{3},
	}

	one.next = two
	two.prev = one
	two.next = three
	three.prev = two
	// one <--> two <--> three

	two.remove()
	// it should be one <--> three
	is.Equal(3, one.next.data.GetData())
	is.Equal(1, three.prev.data.GetData())
}
