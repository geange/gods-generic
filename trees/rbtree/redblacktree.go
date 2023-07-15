// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rbtree implements a red-black tree.
//
// Used by TreeSet and TreeMap.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Red%E2%80%93black_tree
package rbtree

import (
	"github.com/geange/gods-generic/cmp"
	"github.com/geange/gods-generic/utils"
)

// Assert Tree implementation
//var _ trees.Tree = (*Tree)(nil)

type color bool

const (
	black, red color = true, false
)

// Tree holds elements of the red-black tree
type Tree[K, V any] struct {
	root       *Node[K, V]
	size       int
	comparator utils.CompareFunc[K]
}

// New instantiates a red-black tree.
func New[K cmp.Ordered, V any]() *Tree[K, V] {
	return &Tree[K, V]{comparator: cmp.Compare[K]}
}

// NewWith instantiates a red-black tree with the custom comparator.
func NewWith[K, V any](comparator utils.CompareFunc[K]) *Tree[K, V] {
	return &Tree[K, V]{comparator: comparator}
}

func (t *Tree[K, V]) Comparator() utils.CompareFunc[K] {
	return t.comparator
}

func (t *Tree[K, V]) Root() *Node[K, V] {
	return t.root
}

// Put inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Put(key K, value V) {
	var insertedNode *Node[K, V]
	if t.root == nil {
		// Assert key is of comparator's type for initial tree
		t.comparator(key, key)
		t.root = &Node[K, V]{Key: key, Value: value, color: red}
		insertedNode = t.root
	} else {
		node := t.root
		loop := true
		for loop {
			compare := t.comparator(key, node.Key)
			switch {
			case compare == 0:
				node.Key = key
				node.Value = value
				return
			case compare < 0:
				if node.Left == nil {
					node.Left = &Node[K, V]{Key: key, Value: value, color: red}
					insertedNode = node.Left
					loop = false
				} else {
					node = node.Left
				}
			case compare > 0:
				if node.Right == nil {
					node.Right = &Node[K, V]{Key: key, Value: value, color: red}
					insertedNode = node.Right
					loop = false
				} else {
					node = node.Right
				}
			}
		}
		insertedNode.Parent = node
	}
	t.insertCase1(insertedNode)
	t.size++
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Get(key K) (value V, found bool) {
	node := t.lookup(key)
	if node != nil {
		return node.Value, true
	}
	return value, false
}

// GetNode searches the node in the tree by key and returns its node or nil if key is not found in tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) GetNode(key K) *Node[K, V] {
	return t.lookup(key)
}

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Remove(key K) bool {
	var child *Node[K, V]
	node := t.lookup(key)
	if node == nil {
		return false
	}
	if node.Left != nil && node.Right != nil {
		pred := node.Left.maximumNode()
		node.Key = pred.Key
		node.Value = pred.Value
		node = pred
	}
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.color == black {
			node.color = nodeColor(child)
			t.deleteCase1(node)
		}
		t.replaceNode(node, child)
		if node.Parent == nil && child != nil {
			child.color = black
		}
	}
	t.size--
	return true
}

// Empty returns true if tree does not contain any nodes
func (t *Tree[K, V]) Empty() bool {
	return t.size == 0
}

// Size returns number of nodes in the tree.
func (t *Tree[K, V]) Size() int {
	return t.size
}

// Keys returns all keys in-order
func (t *Tree[K, V]) Keys() []K {
	keys := make([]K, t.size)
	it := t.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key()
	}
	return keys
}

// Values returns all values in-order based on the key.
func (t *Tree[K, V]) Values() []V {
	values := make([]V, t.size)
	it := t.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value()
	}
	return values
}

// Left returns the left-most (min) node or nil if tree is empty.
func (t *Tree[K, V]) Left() *Node[K, V] {
	var parent *Node[K, V]
	current := t.root
	for current != nil {
		parent = current
		current = current.Left
	}
	return parent
}

// Right returns the right-most (max) node or nil if tree is empty.
func (t *Tree[K, V]) Right() *Node[K, V] {
	var parent *Node[K, V]
	current := t.root
	for current != nil {
		parent = current
		current = current.Right
	}
	return parent
}

// Floor Finds floor node of the input key, return the floor node or nil if no floor is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree are larger than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Floor(key K) (floor *Node[K, V], found bool) {
	found = false
	node := t.root
	for node != nil {
		compare := t.comparator(key, node.Key)
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			node = node.Left
		case compare > 0:
			floor, found = node, true
			node = node.Right
		}
	}
	if found {
		return floor, true
	}
	return nil, false
}

// Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling is found.
// Second return parameter is true if ceiling was found, otherwise false.
//
// Ceiling node is defined as the smallest node that is larger than or equal to the given node.
// A ceiling node may not be found, either because the tree is empty, or because
// all nodes in the tree are smaller than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Ceiling(key K) (ceiling *Node[K, V], found bool) {
	found = false
	node := t.root
	for node != nil {
		compare := t.comparator(key, node.Key)
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			ceiling, found = node, true
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	if found {
		return ceiling, true
	}
	return nil, false
}

// Clear removes all nodes from the tree.
func (t *Tree[K, V]) Clear() {
	t.root = nil
	t.size = 0
}

// String returns a string representation of container
func (t *Tree[K, V]) String() string {
	str := "RedBlackTree\n"
	if !t.Empty() {
		output(t.root, "", true, &str)
	}
	return str
}

func output[K, V any](node *Node[K, V], prefix string, isTail bool, str *string) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Left, newPrefix, true, str)
	}
}

func (t *Tree[K, V]) lookup(key K) *Node[K, V] {
	node := t.root
	for node != nil {
		compare := t.comparator(key, node.Key)
		switch {
		case compare == 0:
			return node
		case compare < 0:
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	return nil
}

func (t *Tree[K, V]) rotateLeft(node *Node[K, V]) {
	right := node.Right
	t.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (t *Tree[K, V]) rotateRight(node *Node[K, V]) {
	left := node.Left
	t.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

func (t *Tree[K, V]) replaceNode(old *Node[K, V], new *Node[K, V]) {
	if old.Parent == nil {
		t.root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

func (t *Tree[K, V]) insertCase1(node *Node[K, V]) {
	if node.Parent == nil {
		node.color = black
	} else {
		t.insertCase2(node)
	}
}

func (t *Tree[K, V]) insertCase2(node *Node[K, V]) {
	if nodeColor(node.Parent) == black {
		return
	}
	t.insertCase3(node)
}

func (t *Tree[K, V]) insertCase3(node *Node[K, V]) {
	uncle := node.uncle()
	if nodeColor(uncle) == red {
		node.Parent.color = black
		uncle.color = black
		node.grandparent().color = red
		t.insertCase1(node.grandparent())
	} else {
		t.insertCase4(node)
	}
}

func (t *Tree[K, V]) insertCase4(node *Node[K, V]) {
	grandparent := node.grandparent()
	if node == node.Parent.Right && node.Parent == grandparent.Left {
		t.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		t.rotateRight(node.Parent)
		node = node.Right
	}
	t.insertCase5(node)
}

func (t *Tree[K, V]) insertCase5(node *Node[K, V]) {
	node.Parent.color = black
	grandparent := node.grandparent()
	grandparent.color = red
	if node == node.Parent.Left && node.Parent == grandparent.Left {
		t.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		t.rotateLeft(grandparent)
	}
}

func (t *Tree[K, V]) deleteCase1(node *Node[K, V]) {
	if node.Parent == nil {
		return
	}
	t.deleteCase2(node)
}

func (t *Tree[K, V]) deleteCase2(node *Node[K, V]) {
	sibling := node.sibling()
	if nodeColor(sibling) == red {
		node.Parent.color = red
		sibling.color = black
		if node == node.Parent.Left {
			t.rotateLeft(node.Parent)
		} else {
			t.rotateRight(node.Parent)
		}
	}
	t.deleteCase3(node)
}

func (t *Tree[K, V]) deleteCase3(node *Node[K, V]) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == black &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		t.deleteCase1(node.Parent)
	} else {
		t.deleteCase4(node)
	}
}

func (t *Tree[K, V]) deleteCase4(node *Node[K, V]) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == red &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		node.Parent.color = black
	} else {
		t.deleteCase5(node)
	}
}

func (t *Tree[K, V]) deleteCase5(node *Node[K, V]) {
	sibling := node.sibling()
	if node == node.Parent.Left &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == red &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		sibling.Left.color = black
		t.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Right) == red &&
		nodeColor(sibling.Left) == black {
		sibling.color = red
		sibling.Right.color = black
		t.rotateLeft(sibling)
	}
	t.deleteCase6(node)
}

func (t *Tree[K, V]) deleteCase6(node *Node[K, V]) {
	sibling := node.sibling()
	sibling.color = nodeColor(node.Parent)
	node.Parent.color = black
	if node == node.Parent.Left && nodeColor(sibling.Right) == red {
		sibling.Right.color = black
		t.rotateLeft(node.Parent)
	} else if nodeColor(sibling.Left) == red {
		sibling.Left.color = black
		t.rotateRight(node.Parent)
	}
}

func nodeColor[K, V any](node *Node[K, V]) color {
	if node == nil {
		return black
	}
	return node.color
}

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (t *Tree[K, V]) Iterator() Iterator[K, V] {
	return Iterator[K, V]{tree: t, node: nil, position: begin}
}

// IteratorAt returns a stateful iterator whose elements are key/value pairs that is initialised at a particular node.
func (t *Tree[K, V]) IteratorAt(node *Node[K, V]) Iterator[K, V] {
	return Iterator[K, V]{tree: t, node: node, position: between}
}
