// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treebidimap implements a bidirectional map backed by two red-black tree.
//
// This structure guarantees that the map will be in both ascending key and value order.
//
// Other than key and value ordering, the goal with this structure is to avoid duplication of elements, which can be significant if contained elements are large.
//
// A bidirectional map, or hash bag, is an associative entry structure in which the (key,value) pairs form a one-to-one correspondence.
// Thus the binary relation is functional in each direction: value can also act as a key to key.
// A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Bidirectional_map
package treebidimap

import (
	"cmp"
	"fmt"
	"github.com/geange/gods-generic/trees/rbtree"
	"github.com/geange/gods-generic/utils"
	"strings"
)

// Assert Map implementation
//var _ maps.BidiMap = (*Map)(nil)

// Map holds the elements in two red-black trees.
type Map[K, V any] struct {
	forwardMap      *rbtree.Tree[K, *entry[K, V]]
	inverseMap      *rbtree.Tree[V, *entry[K, V]]
	keyComparator   utils.CompareFunc[K]
	valueComparator utils.CompareFunc[V]
}

type entry[K, V any] struct {
	key   K
	value V
}

// New instantiates a bidirectional map.
func New[K, V cmp.Ordered]() *Map[K, V] {
	return &Map[K, V]{
		forwardMap:      rbtree.New[K, *entry[K, V]](),
		inverseMap:      rbtree.New[V, *entry[K, V]](),
		keyComparator:   cmp.Compare[K],
		valueComparator: cmp.Compare[V],
	}
}

// NewWith instantiates a bidirectional map with key+value comparator.
func NewWith[K, V any](keyComparator utils.CompareFunc[K], valueComparator utils.CompareFunc[V]) *Map[K, V] {
	return &Map[K, V]{
		forwardMap:      rbtree.NewWith[K, *entry[K, V]](keyComparator),
		inverseMap:      rbtree.NewWith[V, *entry[K, V]](valueComparator),
		keyComparator:   keyComparator,
		valueComparator: valueComparator,
	}
}

func (m *Map[K, V]) New() *Map[K, V] {
	return &Map[K, V]{
		forwardMap:      rbtree.NewWith[K, *entry[K, V]](m.keyComparator),
		inverseMap:      rbtree.NewWith[V, *entry[K, V]](m.valueComparator),
		keyComparator:   m.keyComparator,
		valueComparator: m.valueComparator,
	}
}

// Put inserts element into the map.
func (m *Map[K, V]) Put(key K, value V) {
	if d, ok := m.forwardMap.Get(key); ok {
		m.inverseMap.Remove(d.value)
	}
	if d, ok := m.inverseMap.Get(value); ok {
		m.forwardMap.Remove(d.key)
	}
	d := &entry[K, V]{key: key, value: value}
	m.forwardMap.Put(key, d)
	m.inverseMap.Put(value, d)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	if d, ok := m.forwardMap.Get(key); ok {
		return d.value, true
	}
	return value, false
}

// GetKey searches the element in the map by value and returns its key or nil if value is not found in map.
// Second return parameter is true if value was found, otherwise false.
func (m *Map[K, V]) GetKey(value V) (key K, found bool) {
	if d, ok := m.inverseMap.Get(value); ok {
		return d.key, true
	}
	return key, false
}

// Remove removes the element from the map by key.
func (m *Map[K, V]) Remove(key K) {
	if d, found := m.forwardMap.Get(key); found {
		m.forwardMap.Remove(key)
		m.inverseMap.Remove(d.value)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[K, V]) Size() int {
	return m.forwardMap.Size()
}

// Keys returns all keys (ordered).
func (m *Map[K, V]) Keys() []K {
	return m.forwardMap.Keys()
}

// Values returns all values (ordered).
func (m *Map[K, V]) Values() []V {
	return m.inverseMap.Keys()
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	m.forwardMap.Clear()
	m.inverseMap.Clear()
}

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (m *Map[K, V]) Iterator() Iterator[K, V] {
	return Iterator[K, V]{iterator: m.forwardMap.Iterator()}
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "TreeBidiMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"
}
