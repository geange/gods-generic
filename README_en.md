[![GoDoc](https://godoc.org/github.com/geange/gods-generic?status.svg)](https://godoc.org/github.com/geange/gods-generic)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/geange/gods-generic/tree/main.svg?style=shield)](https://dl.circleci.com/status-badge/redirect/gh/geange/gods-generic/tree/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/geange/gods-generic)](https://goreportcard.com/report/github.com/geange/gods-generic)
[![codecov](https://codecov.io/gh/geange/gods-generic/branch/master/graph/badge.svg)](https://codecov.io/gh/geange/gods-generic)
[![Sourcegraph](https://sourcegraph.com/github.com/geange/gods-generic/-/badge.svg)](https://sourcegraph.com/github.com/geange/gods-generic?badge)
[![Release](https://img.shields.io/github/release/geange/gods-generic.svg?style=flat-square)](https://github.com/geange/gods-generic/releases)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=gods&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=gods)
[![PyPI](https://img.shields.io/badge/License-BSD_2--Clause-green.svg)](https://github.com/geange/gods-generic/blob/main/LICENSE)

# GoDS (Go Generic Data Structures)

> go1.18+

Implementation of various data structures and algorithms in Go. This project is developed based on
the [gods](https://github.com/emirpasic/gods) project. In the process of
developing [lucene-go](https://github.com/geange/lucene-go), a large amount of
Go1.18+paradigm syntax was used. However, due to the fact that the original project code was not a paradigm syntax, many
problems were encountered during the development process, leading to the idea of implementing a paradigm version of
Gods. Gods generic mainly reimagines data structures such as Sets, Lists, Stacks, Maps, Trees, and Queues using a
paradigm approach, and removes avltree (which I personally do not like) and the serialization method of data
structures (due to time issues)

---

## Data Structures

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
        - [RedBlackTree](#redblacktree)
        - [AVLTree](#avltree)
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

### License

This library is distributed under the BSD-style license found in
the [LICENSE](https://github.com/geange/gods-generic/blob/main/LICENSE) file.
