// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashmap implements a map backed by a hash table.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package hashmap

import (
	"cmp"
	"fmt"
	"github.com/geange/gods-generic/trees/rbtree"
	"github.com/geange/gods-generic/utils"
)

// Assert Map implementation
// var _ maps.Map = (*Map)(nil)

// Map holds the elements in go's native map
type Map[K, V any] struct {
	m *rbtree.Tree[K, V]
}

// New instantiates a hash map.
func New[K cmp.Ordered, V any]() *Map[K, V] {
	return &Map[K, V]{m: rbtree.New[K, V]()}
}

// NewWith instantiates a hash map with key comparator.
func NewWith[K, V any](comparator utils.CompareFunc[K]) *Map[K, V] {
	return &Map[K, V]{m: rbtree.NewWith[K, V](comparator)}
}

// Put inserts element into the map.
func (m *Map[K, V]) Put(key K, value V) {
	m.m.Put(key, value)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	return m.m.Get(key)
}

// Remove removes the element from the map by key.
func (m *Map[K, V]) Remove(key K) {
	//delete(m.m, key)
	m.m.Remove(key)
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[K, V]) Size() int {
	return m.m.Size()
}

// Keys returns all keys (random order).
func (m *Map[K, V]) Keys() []K {
	return m.m.Keys()
}

// Values returns all values (random order).
func (m *Map[K, V]) Values() []V {
	return m.m.Values()
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	// clear(m.m)
	m.m.Clear()
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "HashMap\n"
	str += fmt.Sprintf("%v", m.m)
	return str
}
