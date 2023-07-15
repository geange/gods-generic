// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linkedhashset is a set that preserves insertion-order.
//
// It is backed by a hash table to store values and doubly-linked list to store ordering.
//
// Note that insertion-order is not affected if an element is re-inserted into the set.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package linkedhashset

import (
	"fmt"
	"strings"

	"github.com/geange/gods-generic/cmp"
	"github.com/geange/gods-generic/lists/doublylinkedlist"
	"github.com/geange/gods-generic/maps/treemap"
	"github.com/geange/gods-generic/utils"
)

// Assert Set implementation
//var _ sets.Set = (*Set)(nil)

// Set holds elements in go's native map
type Set[T any] struct {
	table    *treemap.Map[T, struct{}]
	ordering *doublylinkedlist.List[T]
}

var itemExists = struct{}{}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New[T cmp.Ordered](values ...T) *Set[T] {
	set := &Set[T]{
		table:    treemap.New[T, struct{}](),
		ordering: doublylinkedlist.New[T](),
	}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// NewWith instantiates a new empty set with comparator and adds the passed values, if any, to the set
func NewWith[T any](comparator utils.CompareFunc[T], values ...T) *Set[T] {
	set := &Set[T]{
		table:    treemap.NewWith[T, struct{}](comparator),
		ordering: doublylinkedlist.NewWith[T](comparator),
	}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

func (set *Set[T]) New(values ...T) *Set[T] {
	return &Set[T]{
		table:    treemap.NewWith[T, struct{}](set.table.Comparator()),
		ordering: doublylinkedlist.NewWith[T](set.table.Comparator()),
	}
}

// Add adds the items (one or more) to the set.
// Note that insertion-order is not affected if an element is re-inserted into the set.
func (set *Set[T]) Add(items ...T) {
	for _, item := range items {
		if contains := set.Contains(item); !contains {
			set.table.Put(item, itemExists)
			set.ordering.Append(item)
		}
	}
}

// Remove removes the items (one or more) from the set.
// Slow operation, worst-case O(n^2).
func (set *Set[T]) Remove(items ...T) {
	for _, item := range items {
		if contains := set.Contains(item); contains {
			set.table.Remove(item)
			index := set.ordering.IndexOf(item)
			set.ordering.Remove(index)
		}
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, contains := set.table.Get(item); !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set[T]) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set[T]) Size() int {
	return set.ordering.Size()
}

// Clear clears all values in the set.
func (set *Set[T]) Clear() {
	set.table.Clear()
	set.ordering.Clear()
}

// Values returns all items in the set.
func (set *Set[T]) Values() []T {
	values := make([]T, set.Size())
	it := set.Iterator()
	for it.Next() {
		values[it.Index()] = it.Value()
	}
	return values
}

// String returns a string representation of container
func (set *Set[T]) String() string {
	str := "LinkedHashSet\n"
	items := []string{}
	it := set.Iterator()
	for it.Next() {
		items = append(items, fmt.Sprintf("%v", it.Value()))
	}
	str += strings.Join(items, ", ")
	return str
}

// Intersection returns the intersection between two sets.
// The new set consists of all elements that are both in "set" and "another".
// Ref: https://en.wikipedia.org/wiki/Intersection_(set_theory)
func (set *Set[T]) Intersection(another *Set[T]) *Set[T] {
	result := set.New()

	// Iterate over smaller set (optimization)
	if set.Size() <= another.Size() {
		for _, item := range set.Values() {
			if contains := another.Contains(item); contains {
				result.Add(item)
			}
		}
	} else {
		for _, item := range another.Values() {
			if contains := set.Contains(item); contains {
				result.Add(item)
			}
		}
	}

	return result
}

// Union returns the union of two sets.
// The new set consists of all elements that are in "set" or "another" (possibly both).
// Ref: https://en.wikipedia.org/wiki/Union_(set_theory)
func (set *Set[T]) Union(another *Set[T]) *Set[T] {
	result := set.New()

	items1 := set.Values()
	for i := range items1 {
		result.Add(items1[i])
	}

	items2 := another.Values()
	for i := range items2 {
		result.Add(items2[i])
	}

	return result
}

// Difference returns the difference between two sets.
// The new set consists of all elements that are in "set" but not in "another".
// Ref: https://proofwiki.org/wiki/Definition:Set_Difference
func (set *Set[T]) Difference(another *Set[T]) *Set[T] {
	result := set.New()

	items := set.Values()
	for i := range items {
		item := items[i]
		if contains := another.Contains(item); !contains {
			result.Add(items[i])
		}
	}

	return result
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (set *Set[T]) Iterator() Iterator[T] {
	return Iterator[T]{iterator: set.ordering.Iterator()}
}
