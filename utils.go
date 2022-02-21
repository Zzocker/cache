package cache

import (
	"reflect"
)

func isKeyTypeValid(key interface{}) bool {
	if key == nil {
		return false
	}
	switch reflect.TypeOf(key).Kind() {
	case reflect.Invalid, reflect.Uintptr, reflect.Complex64, reflect.Complex128, reflect.Array, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return false
	}
	return true
}

// node : of a doubly circular linked list.
type node struct {
	data       DataHandler
	prev, next *node
}

// insert : node at given node
// before calling: at <--> at.next
// after calling: at <--> n <--> old_at.next.
func (n *node) insert(at *node) {
	nxt := at.next
	at.next = n
	n.prev = at
	n.next = nxt
	nxt.prev = n
}

// remove: node from the chain
// before calling: n.prev <--> n <--> n.next
// after calling: old_n.prev <--> old_n.next.
func (n *node) remove() {
	n.prev.next = n.next
	n.next.prev = n.prev
	// clean up pointers
	n.prev = nil
	n.next = nil
}
