[![GoDoc](https://godoc.org/github.com/geange/gods-generic?status.svg)](https://godoc.org/github.com/geange/gods-generic)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/geange/gods-generic/tree/main.svg?style=shield)](https://dl.circleci.com/status-badge/redirect/gh/geange/gods-generic/tree/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/geange/gods-generic)](https://goreportcard.com/report/github.com/geange/gods-generic)
[![codecov](https://codecov.io/gh/geange/gods-generic/branch/master/graph/badge.svg)](https://codecov.io/gh/geange/gods-generic)
[![Source graph](https://sourcegraph.com/github.com/geange/gods-generic/-/badge.svg)](https://sourcegraph.com/github.com/geange/gods-generic?badge)
[![Release](https://img.shields.io/github/release/geange/gods-generic.svg?style=flat-square)](https://github.com/geange/gods-generic/releases)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=gods&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=gods)
[![PyPI](https://img.shields.io/badge/License-BSD_2--Clause-green.svg)](https://github.com/geange/gods-generic/blob/main/LICENSE)

# GoDS (Golang范型数据结构)

## Other languages

* [English](README_en.md)

> go1.18+

Go中各种数据结构和算法的实现。

该项目基于[gods](https://github.com/emirpasic/gods)项目进行开发。在开发[lucene-go](https://github.com/geange/lucene-go)
的过程中使用了大量的go1.18+范型语法，在使用gods的过程中，由于原项目代码并非范型语法，开发过程中遇到不少问题，
便萌生想法实现一个范型版本的gods。

gods-generic主要使用范型的方式重新实现Sets, Lists, Stacks, Maps, Trees, Queues等数据结构，并移除了avltree（个人不太喜欢这种数据格式），
并移除了数据结构的序列化方式（由于时间问题，尚未重新进行设计）

---

## 数据结构

- [Containers](#containers)
    - [Lists](#lists)
        - [ArrayList](#arraylist)
        - [SinglyLinkedList](#singlylinkedlist)
        - [DoublyLinkedList](#doublylinkedlist)
    - [Sets](#sets)
        - [HashSet](#hashset)
        - [TreeSet](#treeset)
        - [LinkedHashSet](#linkedhashset)
    - [Stacks](#stacks)
        - [LinkedListStack](#linkedliststack)
        - [ArrayStack](#arraystack)
    - [Maps](#maps)
        - [HashMap](#hashmap)
        - [TreeMap](#treemap)
        - [LinkedHashMap](#linkedhashmap)
        - [HashBidiMap](#hashbidimap)
        - [TreeBidiMap](#treebidimap)
    - [Trees](#trees)
        - [RedBlackTree](#rbtree)
        - [BTree](#btree)
        - [BinaryHeap](#binaryheap)
    - [Queues](#queues)
        - [LinkedListQueue](#linkedlistqueue)
        - [ArrayQueue](#arrayqueue)
        - [CircularBuffer](#circularbuffer)
        - [PriorityQueue](#priorityqueue)
- [Functions](#functions)
    - [Comparator](#comparator)
    - [Iterator](#iterator)
        - [IteratorWithIndex](#iteratorwithindex)
        - [IteratorWithKey](#iteratorwithkey)
        - [ReverseIteratorWithIndex](#reverseiteratorwithindex)
        - [ReverseIteratorWithKey](#reverseiteratorwithkey)
    - [Enumerable](#enumerable)
        - [EnumerableWithIndex](#enumerablewithindex)
        - [EnumerableWithKey](#enumerablewithkey)
    - [Serialization](#serialization)
        - [JSONSerializer](#jsonserializer)
        - [JSONDeserializer](#jsondeserializer)
    - [Sort](#sort)
    - [Container](#container)
- [Appendix](#appendix)

## containers

```go
// Container is base interface that all data structures implement.
type Container[T any] interface {
    Empty() bool
    Size() int
    Clear()
    Values() []T
    String() string
}
```

### lists

```go
// List interface that all lists implement
type List[T any] interface {
    Get(index int) (T, bool)
    Remove(index int)
    Add(values ...T)
    Contains(values ...T) bool
    Sort(comparator utils.Comparator)
    Swap(index1, index2 int)
    Insert(index int, values ...T)
    Set(index int, value T)

    containers.Container[T]
}
```

#### arrayList

```go
package main

import (
    "github.com/geange/gods-generic/cmp"
    "github.com/geange/gods-generic/lists/arraylist"
)

// ArrayListExample to demonstrate basic usage of ArrayList
func main() {
    list := arraylist.New[string]()
    list.Add("a")                         // ["a"]
    list.Add("c", "b")                    // ["a","c","b"]
    list.Sort(cmp.Compare[string])        // ["a","b","c"]
    _, _ = list.Get(0)                    // "a",true
    _, _ = list.Get(100)                  // nil,false
    _ = list.Contains("a", "b", "c")      // true
    _ = list.Contains("a", "b", "c", "d") // false
    list.Swap(0, 1)                       // ["b","a",c"]
    list.Remove(2)                        // ["b","a"]
    list.Remove(1)                        // ["b"]
    list.Remove(0)                        // []
    list.Remove(0)                        // [] (ignored)
    _ = list.Empty()                      // true
    _ = list.Size()                       // 0
    list.Add("a")                         // ["a"]
    list.Clear()                          // []
}

```

#### singlyLinkedList

```go
package main

import (
	"github.com/geange/gods-generic/cmp"
	sll "github.com/geange/gods-generic/lists/singlylinkedlist"
)

// SinglyLinkedListExample to demonstrate basic usage of SinglyLinkedList
func main() {
	list := sll.New[string]()
	list.Add("a")                         // ["a"]
	list.Append("b")                      // ["a","b"] (same as Add())
	list.Prepend("c")                     // ["c","a","b"]
	list.Sort(cmp.Compare[string])        // ["a","b","c"]
	_, _ = list.Get(0)                    // "a",true
	_, _ = list.Get(100)                  // nil,false
	_ = list.Contains("a", "b", "c")      // true
	_ = list.Contains("a", "b", "c", "d") // false
	list.Remove(2)                        // ["a","b"]
	list.Remove(1)                        // ["a"]
	list.Remove(0)                        // []
	list.Remove(0)                        // [] (ignored)
	_ = list.Empty()                      // true
	_ = list.Size()                       // 0
	list.Add("a")                         // ["a"]
	list.Clear()                          // []
}
```

#### doublyLinkedList

```go
package main

import (
	"github.com/geange/gods-generic/cmp"
	dll "github.com/geange/gods-generic/lists/doublylinkedlist"
)

// DoublyLinkedListExample to demonstrate basic usage of DoublyLinkedList
func main() {
	list := dll.New[string]()
	list.Add("a")                         // ["a"]
	list.Append("b")                      // ["a","b"] (same as Add())
	list.Prepend("c")                     // ["c","a","b"]
	list.Sort(cmp.Compare[string])        // ["a","b","c"]
	_, _ = list.Get(0)                    // "a",true
	_, _ = list.Get(100)                  // nil,false
	_ = list.Contains("a", "b", "c")      // true
	_ = list.Contains("a", "b", "c", "d") // false
	list.Remove(2)                        // ["a","b"]
	list.Remove(1)                        // ["a"]
	list.Remove(0)                        // []
	list.Remove(0)                        // [] (ignored)
	_ = list.Empty()                      // true
	_ = list.Size()                       // 0
	list.Add("a")                         // ["a"]
	list.Clear()                          // []
}

```

### sets

```go
type Set[T any] interface {
	Add(elements ...T)
	Remove(elements ...T)
	Contains(elements ...T) bool

	containers.Container[T]
}

```

#### hashset

```go
package main

import "github.com/geange/gods-generic/sets/hashset"

// HashSetExample to demonstrate basic usage of HashSet
func main() {
	set := hashset.New[int]() // empty (keys are of type int)
	set.Add(1)                // 1
	set.Add(2, 2, 3, 4, 5)    // 3, 1, 2, 4, 5 (random order, duplicates ignored)
	set.Remove(4)             // 5, 3, 2, 1 (random order)
	set.Remove(2, 3)          // 1, 5 (random order)
	set.Contains(1)           // true
	set.Contains(1, 5)        // true
	set.Contains(1, 6)        // false
	_ = set.Values()          // []int{5,1} (random order)
	set.Clear()               // empty
	set.Empty()               // true
	set.Size()                // 0
}

```

#### treeSet

```go
package main

import "github.com/geange/gods-generic/sets/treeset"

// TreeSetExample to demonstrate basic usage of TreeSet
func main() {
	set := treeset.New[int]() // empty
	set.Add(1)                // 1
	set.Add(2, 2, 3, 4, 5)    // 1, 2, 3, 4, 5 (in order, duplicates ignored)
	set.Remove(4)             // 1, 2, 3, 5 (in order)
	set.Remove(2, 3)          // 1, 5 (in order)
	set.Contains(1)           // true
	set.Contains(1, 5)        // true
	set.Contains(1, 6)        // false
	_ = set.Values()          // []int{1,5} (in order)
	set.Clear()               // empty
	set.Empty()               // true
	set.Size()                // 0
}
```

#### linkedHashset

```go
package main

import "github.com/geange/gods-generic/sets/linkedhashset"

// LinkedHashSetExample to demonstrate basic usage of LinkedHashSet
func main() {
	set := linkedhashset.New[int]() // empty
	set.Add(5)                      // 5
	set.Add(4, 4, 3, 2, 1)          // 5, 4, 3, 2, 1 (in insertion-order, duplicates ignored)
	set.Remove(4)                   // 5, 3, 2, 1 (in insertion-order)
	set.Remove(2, 3)                // 5, 1 (in insertion-order)
	set.Contains(1)                 // true
	set.Contains(1, 5)              // true
	set.Contains(1, 6)              // false
	_ = set.Values()                // []int{5, 1} (in insertion-order)
	set.Clear()                     // empty
	set.Empty()                     // true
	set.Size()                      // 0
}
```

### stacks

```go
type Stack[T any] interface {
	Push(value T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)

	containers.Container[T]
}
```

#### linkedListStack

```go
package main

import lls "github.com/geange/gods-generic/stacks/linkedliststack"

// LinkedListStackExample to demonstrate basic usage of LinkedListStack
func main() {
	stack := lls.New[int]() // empty
	stack.Push(1)           // 1
	stack.Push(2)           // 1, 2
	stack.Values()          // 2, 1 (LIFO order)
	_, _ = stack.Peek()     // 2,true
	_, _ = stack.Pop()      // 2, true
	_, _ = stack.Pop()      // 1, true
	_, _ = stack.Pop()      // nil, false (nothing to pop)
	stack.Push(1)           // 1
	stack.Clear()           // empty
	stack.Empty()           // true
	stack.Size()            // 0
}
```

#### arraystack

```go
package main

import "github.com/geange/gods-generic/stacks/arraystack"

// ArrayStackExample to demonstrate basic usage of ArrayStack
func main() {
	stack := arraystack.New[int]() // empty
	stack.Push(1)                  // 1
	stack.Push(2)                  // 1, 2
	stack.Values()                 // 2, 1 (LIFO order)
	_, _ = stack.Peek()            // 2,true
	_, _ = stack.Pop()             // 2, true
	_, _ = stack.Pop()             // 1, true
	_, _ = stack.Pop()             // nil, false (nothing to pop)
	stack.Push(1)                  // 1
	stack.Clear()                  // empty
	stack.Empty()                  // true
	stack.Size()                   // 0
}

```

### maps

```go
type Map[K, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, found bool)
	Remove(key K)
	Keys() []K

	containers.Container[K]
}

```

#### hashmap

```go
package main

import "github.com/geange/gods-generic/maps/hashmap"

// HashMapExample to demonstrate basic usage of HashMap
func main() {
	m := hashmap.New[int, string]() // empty
	m.Put(1, "x")                   // 1->x
	m.Put(2, "b")                   // 2->b, 1->x  (random order)
	m.Put(1, "a")                   // 2->b, 1->a (random order)
	_, _ = m.Get(2)                 // b, true
	_, _ = m.Get(3)                 // nil, false
	_ = m.Values()                  // []interface {}{"b", "a"} (random order)
	_ = m.Keys()                    // []interface {}{1, 2} (random order)
	m.Remove(1)                     // 2->b
	m.Clear()                       // empty
	m.Empty()                       // true
	m.Size()                        // 0
}
```

#### treemap

```go
package main

import "github.com/geange/gods-generic/maps/treemap"

// TreeMapExample to demonstrate basic usage of TreeMap
func main() {
	m := treemap.New[int, string]() // empty (keys are of type int)
	m.Put(1, "x")                   // 1->x
	m.Put(2, "b")                   // 1->x, 2->b (in order)
	m.Put(1, "a")                   // 1->a, 2->b (in order)
	_, _ = m.Get(2)                 // b, true
	_, _ = m.Get(3)                 // nil, false
	_ = m.Values()                  // []interface {}{"a", "b"} (in order)
	_ = m.Keys()                    // []interface {}{1, 2} (in order)
	m.Remove(1)                     // 2->b
	m.Clear()                       // empty
	m.Empty()                       // true
	m.Size()                        // 0
}

```

#### linkedhashmap

```go
package main

import "github.com/geange/gods-generic/maps/linkedhashmap"

// LinkedHashMapExample to demonstrate basic usage of LinkedHashMapExample
func main() {
	m := linkedhashmap.New[int, string]() // empty (keys are of type int)
	m.Put(2, "b")                         // 2->b
	m.Put(1, "x")                         // 2->b, 1->x (insertion-order)
	m.Put(1, "a")                         // 2->b, 1->a (insertion-order)
	_, _ = m.Get(2)                       // b, true
	_, _ = m.Get(3)                       // nil, false
	_ = m.Values()                        // []interface {}{"b", "a"} (insertion-order)
	_ = m.Keys()                          // []interface {}{2, 1} (insertion-order)
	m.Remove(1)                           // 2->b
	m.Clear()                             // empty
	m.Empty()                             // true
	m.Size()                              // 0
}

```

#### hashbidimap

```go
package main

import "github.com/geange/gods-generic/maps/hashbidimap"

// HashBidiMapExample to demonstrate basic usage of HashMap
func main() {
	m := hashbidimap.New[int, string]() // empty
	m.Put(1, "x")                       // 1->x
	m.Put(3, "b")                       // 1->x, 3->b (random order)
	m.Put(1, "a")                       // 1->a, 3->b (random order)
	m.Put(2, "b")                       // 1->a, 2->b (random order)
	_, _ = m.GetKey("a")                // 1, true
	_, _ = m.Get(2)                     // b, true
	_, _ = m.Get(3)                     // nil, false
	_ = m.Values()                      // []interface {}{"a", "b"} (random order)
	_ = m.Keys()                        // []interface {}{1, 2} (random order)
	m.Remove(1)                         // 2->b
	m.Clear()                           // empty
	m.Empty()                           // true
	m.Size()                            // 0
}
```

#### treebidimap

```go
package main

import (
	"github.com/geange/gods-generic/maps/treebidimap"
)

// TreeBidiMapExample to demonstrate basic usage of TreeBidiMap
func main() {
	m := treebidimap.New[int, string]()
	m.Put(1, "x")        // 1->x
	m.Put(3, "b")        // 1->x, 3->b (ordered)
	m.Put(1, "a")        // 1->a, 3->b (ordered)
	m.Put(2, "b")        // 1->a, 2->b (ordered)
	_, _ = m.GetKey("a") // 1, true
	_, _ = m.Get(2)      // b, true
	_, _ = m.Get(3)      // nil, false
	_ = m.Values()       // []interface {}{"a", "b"} (ordered)
	_ = m.Keys()         // []interface {}{1, 2} (ordered)
	m.Remove(1)          // 2->b
	m.Clear()            // empty
	m.Empty()            // true
	m.Size()             // 0
}
```

### trees

```go
type Tree[T any] interface {
	containers.Container[T]
}
```

#### rbtree

```go
package main

import (
	"fmt"
	"github.com/geange/gods-generic/trees/rbtree"
)

// RedBlackTreeExample to demonstrate basic usage of RedBlackTree
func main() {
	tree := rbtree.New[int, string]() // empty(keys are of type int)

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)

	fmt.Println(tree)
	//
	//  RedBlackTree
	//  │           ┌── 6
	//  │       ┌── 5
	//  │   ┌── 4
	//  │   │   └── 3
	//  └── 2
	//       └── 1

	_ = tree.Values() // []interface {}{"a", "b", "c", "d", "e", "f"} (in order)
	_ = tree.Keys()   // []interface {}{1, 2, 3, 4, 5, 6} (in order)

	tree.Remove(2) // 1->a, 3->c, 4->d, 5->e, 6->f (in order)
	fmt.Println(tree)
	//
	//  RedBlackTree
	//  │       ┌── 6
	//  │   ┌── 5
	//  └── 4
	//      │   ┌── 3
	//      └── 1

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0
}

```

#### btree

```go
package main

import (
	"fmt"
	"github.com/geange/gods-generic/trees/btree"
)

// BTreeExample to demonstrate basic usage of BTree
func main() {
	tree := btree.New[int, string](3) // empty (keys are of type int)

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)
	tree.Put(7, "g") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f, 7->g (in order)

	fmt.Println(tree)
	// BTree
	//         1
	//     2
	//         3
	// 4
	//         5
	//     6
	//         7

	_ = tree.Values() // []interface {}{"a", "b", "c", "d", "e", "f", "g"} (in order)
	_ = tree.Keys()   // []interface {}{1, 2, 3, 4, 5, 6, 7} (in order)

	tree.Remove(2) // 1->a, 3->c, 4->d, 5->e, 6->f (in order)
	fmt.Println(tree)
	// BTree
	//     1
	//     3
	// 4
	//     5
	//     6

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0

	// Other:
	tree.Height()     // gets the height of the tree
	tree.Left()       // gets the left-most (min) node
	tree.LeftKey()    // get the left-most (min) node's key
	tree.LeftValue()  // get the left-most (min) node's value
	tree.Right()      // get the right-most (max) node
	tree.RightKey()   // get the right-most (max) node's key
	tree.RightValue() // get the right-most (max) node's value
}
```

#### binaryheap

```go
package main

import (
	"github.com/geange/gods-generic/trees/binaryheap"
	"github.com/geange/gods-generic/utils"
)

// BinaryHeapExample to demonstrate basic usage of BinaryHeap
func main() {

	// Min-heap
	heap := binaryheap.New[int]() // empty (min-heap)
	heap.Push(2)                  // 2
	heap.Push(3)                  // 2, 3
	heap.Push(1)                  // 1, 3, 2
	heap.Values()                 // 1, 3, 2
	_, _ = heap.Peek()            // 1,true
	_, _ = heap.Pop()             // 1, true
	_, _ = heap.Pop()             // 2, true
	_, _ = heap.Pop()             // 3, true
	_, _ = heap.Pop()             // nil, false (nothing to pop)
	heap.Push(1)                  // 1
	heap.Clear()                  // empty
	heap.Empty()                  // true
	heap.Size()                   // 0

	// Max-heap
	inverseIntComparator := func(a, b int) int {
		return -utils.IntComparator(a, b)
	}
	heap = binaryheap.NewWith(inverseIntComparator) // empty (min-heap)
	heap.Push(2)                                    // 2
	heap.Push(3)                                    // 3, 2
	heap.Push(1)                                    // 3, 2, 1
	heap.Values()                                   // 3, 2, 1
}
```

### queues

```go
type Queue[T any] interface {
	Enqueue(value T)
	Dequeue() (value T, ok bool)
	Peek() (value T, ok bool)

	containers.Container[T]
}
```

#### linkedlistqueue

```go
package main

import llq "github.com/geange/gods-generic/queues/linkedlistqueue"

// LinkedListQueueExample to demonstrate basic usage of LinkedListQueue
func main() {
	queue := llq.New[int]() // empty
	queue.Enqueue(1)        // 1
	queue.Enqueue(2)        // 1, 2
	_ = queue.Values()      // 1, 2 (FIFO order)
	_, _ = queue.Peek()     // 1,true
	_, _ = queue.Dequeue()  // 1, true
	_, _ = queue.Dequeue()  // 2, true
	_, _ = queue.Dequeue()  // nil, false (nothing to deque)
	queue.Enqueue(1)        // 1
	queue.Clear()           // empty
	queue.Empty()           // true
	_ = queue.Size()        // 0
}
```

#### arrayqueue

```go
package main

import aq "github.com/geange/gods-generic/queues/arrayqueue"

// ArrayQueueExample to demonstrate basic usage of ArrayQueue
func main() {
	queue := aq.New[int]() // empty
	queue.Enqueue(1)       // 1
	queue.Enqueue(2)       // 1, 2
	_ = queue.Values()     // 1, 2 (FIFO order)
	_, _ = queue.Peek()    // 1,true
	_, _ = queue.Dequeue() // 1, true
	_, _ = queue.Dequeue() // 2, true
	_, _ = queue.Dequeue() // nil, false (nothing to deque)
	queue.Enqueue(1)       // 1
	queue.Clear()          // empty
	queue.Empty()          // true
	_ = queue.Size()       // 0
}
```

#### circularbuffer

```go
package main

import cb "github.com/geange/gods-generic/queues/circularbuffer"

// CircularBufferExample to demonstrate basic usage of CircularBuffer
func main() {
	queue := cb.New[int](3) // empty (max size is 3)
	queue.Enqueue(1)        // 1
	queue.Enqueue(2)        // 1, 2
	queue.Enqueue(3)        // 1, 2, 3
	_ = queue.Values()      // 1, 2, 3
	queue.Enqueue(3)        // 4, 2, 3
	_, _ = queue.Peek()     // 4,true
	_, _ = queue.Dequeue()  // 4, true
	_, _ = queue.Dequeue()  // 2, true
	_, _ = queue.Dequeue()  // 3, true
	_, _ = queue.Dequeue()  // nil, false (nothing to deque)
	queue.Enqueue(1)        // 1
	queue.Clear()           // empty
	queue.Empty()           // true
	_ = queue.Size()        // 0
}
```

#### priorityqueue

```go
package main

import (
	pq "github.com/geange/gods-generic/queues/priorityqueue"
	"github.com/geange/gods-generic/utils"
)

// Element is an entry in the priority queue
type Element struct {
	name     string
	priority int
}

// comparator function (sort by element's priority value in descending order)
func byPriority(a, b Element) int {
	priorityA := a.priority
	priorityB := b.priority
	return -utils.IntComparator(priorityA, priorityB) // "-" descending order
}

// PriorityQueueExample to demonstrate basic usage of BinaryHeap
func main() {
	a := Element{name: "a", priority: 1}
	b := Element{name: "b", priority: 2}
	c := Element{name: "c", priority: 3}

	queue := pq.NewWith(byPriority) // empty
	queue.Enqueue(a)                // {a 1}
	queue.Enqueue(c)                // {c 3}, {a 1}
	queue.Enqueue(b)                // {c 3}, {b 2}, {a 1}
	_ = queue.Values()              // [{c 3} {b 2} {a 1}]
	_, _ = queue.Peek()             // {c 3} true
	_, _ = queue.Dequeue()          // {c 3} true
	_, _ = queue.Dequeue()          // {b 2} true
	_, _ = queue.Dequeue()          // {a 1} true
	_, _ = queue.Dequeue()          // <nil> false (nothing to dequeue)
	queue.Clear()                   // empty
	_ = queue.Empty()               // true
	_ = queue.Size()                // 0
}
```

### License

gods-generic
是在BSD风格的许可证下分发的，该许可证位于 [LICENSE](https://github.com/geange/gods-generic/blob/main/LICENSE).