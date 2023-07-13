// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package hashset

import (
	"cmp"
	"fmt"
	"strings"

	"github.com/geange/gods-generic/maps/treemap"
	"github.com/geange/gods-generic/utils"
)

// Assert Set implementation
// var _ sets.Set = (*Set)(nil)

// Set holds elements in go's native map
type Set[T any] struct {
	items *treemap.Map[T, struct{}]
}

var itemExists = struct{}{}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New[T cmp.Ordered](values ...T) *Set[T] {
	set := &Set[T]{items: treemap.New[T, struct{}]()}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// NewWith instantiates a new empty set with comparator and adds the passed values, if any, to the set
func NewWith[T cmp.Ordered](comparator utils.CompareFunc[T], values ...T) *Set[T] {
	set := &Set[T]{items: treemap.NewWith[T, struct{}](comparator)}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

func (set *Set[T]) New() *Set[T] {
	return &Set[T]{
		items: treemap.NewWith[T, struct{}](set.items.Comparator()),
	}
}

// Add adds the items (one or more) to the set.
func (set *Set[T]) Add(items ...T) {
	for _, item := range items {
		set.items.Put(item, struct{}{})
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set[T]) Remove(items ...T) {
	for _, item := range items {
		set.items.Remove(item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		_, contains := set.items.Get(item)
		if !contains {
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
	return set.items.Size()
}

// Clear clears all values in the set.
func (set *Set[T]) Clear() {
	set.items.Clear()
}

// Values returns all items in the set.
func (set *Set[T]) Values() []T {
	return set.items.Keys()
}

// String returns a string representation of container
func (set *Set[T]) String() string {
	str := "HashSet\n"
	items := []string{}
	for k := range set.items.Values() {
		items = append(items, fmt.Sprintf("%v", k))
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
