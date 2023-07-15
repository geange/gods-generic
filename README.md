[![GoDoc](https://godoc.org/github.com/geange/gods-generic?status.svg)](https://godoc.org/github.com/geange/gods-generic)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/geange/gods-generic/tree/main.svg?style=shield)](https://dl.circleci.com/status-badge/redirect/gh/geange/gods-generic/tree/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/geange/gods-generic)](https://goreportcard.com/report/github.com/geange/gods-generic)
[![codecov](https://codecov.io/gh/geange/gods-generic/branch/master/graph/badge.svg)](https://codecov.io/gh/geange/gods-generic)
[![Sourcegraph](https://sourcegraph.com/github.com/geange/gods-generic/-/badge.svg)](https://sourcegraph.com/github.com/geange/gods-generic?badge)
[![Release](https://img.shields.io/github/release/geange/gods-generic.svg?style=flat-square)](https://github.com/geange/gods-generic/releases)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=gods&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=gods)
[![PyPI](https://img.shields.io/badge/License-BSD_2--Clause-green.svg)](https://github.com/geange/gods-generic/blob/main/LICENSE)

# GoDS (Golang范型数据结构)

> go1.18+

该项目基于[gods](https://github.com/emirpasic/gods)项目进行开发。在开发[lucene-go](https://github.com/geange/lucene-go)
的过程中使用了大量的go1.18+范型语法，在使用gods的过程中，由于原项目代码并非范型语法，开发过程中遇到不少问题，
便萌生想法实现一个范型版本的gods。

gods-generic主要使用范型的方式重新实现Sets, Lists, Stacks, Maps, Trees, Queues等数据结构，并移除了avltree（个人不太喜欢这种数据格式），
并移除了数据结构的序列化方式（由于时间问题，尚未重新进行设计）

---

[English](README_en.md)

Implementation of various data structures and algorithms in Go.

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

### License

This library is distributed under the BSD-style license found in
the [LICENSE](https://github.com/geange/gods-generic/blob/main/LICENSE) file.
