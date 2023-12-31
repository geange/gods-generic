// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbtree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type item[K, V any] struct {
	k    K
	v    V
	flag bool
}

func TestRedBlackTreeGet(t *testing.T) {
	tree := New[int, string]()

	if actualValue := tree.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	if actualValue := tree.GetNode(2).Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

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

	if actualValue := tree.Size(); actualValue != 6 {
		t.Errorf("Got %v expected %v", actualValue, 6)
	}

	if actualValue := tree.GetNode(4).Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}

	if actualValue := tree.GetNode(2).Size(); actualValue != 6 {
		t.Errorf("Got %v expected %v", actualValue, 6)
	}

	if actualValue := tree.GetNode(8).Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestRedBlackTreePut(t *testing.T) {
	tree := New[int, string]()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite

	if actualValue := tree.Size(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, tree.Keys())
	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f", "g"}, tree.Values())

	tests1 := []item[int, string]{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, "", false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := tree.Get(test.k)
		if actualValue != test.v || actualFound != test.flag {
			t.Errorf("Got %v expected %v", actualValue, test.v)
		}
	}
}

func TestRedBlackTreeRemove(t *testing.T) {
	tree := New[int, string]()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite

	tree.Remove(5)
	tree.Remove(6)
	tree.Remove(7)
	tree.Remove(8)
	tree.Remove(5)

	assert.Equal(t, []int{1, 2, 3, 4}, tree.Keys())
	assert.Equal(t, []string{"a", "b", "c", "d"}, tree.Values())
	assert.Equal(t, 4, tree.Size())

	tests2 := []item[int, string]{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "", false},
		{6, "", false},
		{7, "", false},
		{8, "", false},
	}

	for _, test := range tests2 {
		actualValue, actualFound := tree.Get(test.k)
		if actualValue != test.v || actualFound != test.flag {
			t.Errorf("Got %v expected %v", actualValue, test.v)
		}
	}

	tree.Remove(1)
	tree.Remove(4)
	tree.Remove(2)
	tree.Remove(3)
	tree.Remove(2)
	tree.Remove(2)

	assert.Equal(t, []int{}, tree.Keys())
	assert.Equal(t, []string{}, tree.Values())
	assert.Equal(t, 0, tree.Size())
}

func TestRedBlackTreeLeftAndRight(t *testing.T) {
	tree := New[int, string]()

	assert.Nil(t, tree.Left())
	assert.Nil(t, tree.Right())

	tree.Put(1, "a")
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x") // overwrite
	tree.Put(2, "b")

	assert.Equal(t, 1, tree.Left().Key)
	assert.Equal(t, "x", tree.Left().Value)
	assert.Equal(t, 7, tree.Right().Key)
	assert.Equal(t, "g", tree.Right().Value)
}

func TestRedBlackTreeCeilingAndFloor(t *testing.T) {
	tree := New[int, string]()

	if node, found := tree.Floor(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}
	if node, found := tree.Ceiling(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}

	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")

	if node, found := tree.Floor(4); node.Key != 4 || !found {
		t.Errorf("Got %v expected %v", node.Key, 4)
	}
	if node, found := tree.Floor(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}

	if node, found := tree.Ceiling(4); node.Key != 4 || !found {
		t.Errorf("Got %v expected %v", node.Key, 4)
	}
	if node, found := tree.Ceiling(8); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}
}

func TestRedBlackTreeIteratorNextOnEmpty(t *testing.T) {
	tree := New[int, struct{}]()
	it := tree.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty tree")
	}
}

func TestRedBlackTreeIteratorPrevOnEmpty(t *testing.T) {
	tree := New[int, struct{}]()
	it := tree.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty tree")
	}
}

func TestRedBlackTreeIterator1Next(t *testing.T) {
	tree := New[int, string]()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite
	// │   ┌── 7
	// └── 6
	//     │   ┌── 5
	//     └── 4
	//         │   ┌── 3
	//         └── 2
	//             └── 1
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		if actualValue, expectedValue := key, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator1Prev(t *testing.T) {
	tree := New[int, string]()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite
	// │   ┌── 7
	// └── 6
	//     │   ┌── 5
	//     └── 4
	//         │   ┌── 3
	//         └── 2
	//             └── 1
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator2Next(t *testing.T) {
	tree := New[int, string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		if actualValue, expectedValue := key, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator2Prev(t *testing.T) {
	tree := New[int, string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator3Next(t *testing.T) {
	tree := New[int, string]()
	tree.Put(1, "a")
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		if actualValue, expectedValue := key, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator3Prev(t *testing.T) {
	tree := New[int, string]()
	tree.Put(1, "a")
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator4Next(t *testing.T) {
	tree := New[int, int]()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	// │           ┌── 27
	// │       ┌── 25
	// │       │   └── 22
	// │   ┌── 17
	// │   │   └── 15
	// └── 13
	//     │   ┌── 11
	//     └── 8
	//         │   ┌── 6
	//         └── 1
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		value := it.Value()
		if actualValue, expectedValue := value, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator4Prev(t *testing.T) {
	tree := New[int, int]()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	// │           ┌── 27
	// │       ┌── 25
	// │       │   └── 22
	// │   ┌── 17
	// │   │   └── 15
	// └── 13
	//     │   ┌── 11
	//     └── 8
	//         │   ┌── 6
	//         └── 1
	it := tree.Iterator()
	count := tree.Size()
	for it.Next() {
	}
	for it.Prev() {
		value := it.Value()
		if actualValue, expectedValue := value, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		count--
	}
	if actualValue, expectedValue := count, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIteratorBegin(t *testing.T) {
	tree := New[int, string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.Begin()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	for it.Next() {
	}

	it.Begin()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.Next()
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestRedBlackTreeIteratorEnd(t *testing.T) {
	tree := New[int, string]()
	it := tree.Iterator()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.End()
	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it.End()
	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.Prev()
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestRedBlackTreeIteratorFirst(t *testing.T) {
	tree := New[int, string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestRedBlackTreeIteratorLast(t *testing.T) {
	tree := New[int, string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestRedBlackTreeIteratorNextTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// NextTo (empty)
	{
		tree := New[int, string]()
		it := tree.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty tree")
		}
	}

	// NextTo (not found)
	{
		tree := New[int, string]()
		tree.Put(0, "xx")
		tree.Put(1, "yy")
		it := tree.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty tree")
		}
	}

	// NextTo (found)
	{
		tree := New[int, string]()
		tree.Put(2, "cc")
		tree.Put(0, "aa")
		tree.Put(1, "bb")
		it := tree.Iterator()
		it.Begin()
		if !it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty tree")
		}
		if index, value := it.Key(), it.Value(); index != 1 || value != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Next() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Key(), it.Value(); index != 2 || value != "cc" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "cc")
		}
		if it.Next() {
			t.Errorf("Should not go past last element")
		}
	}
}

func TestRedBlackTreeIteratorPrevTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// PrevTo (empty)
	{
		tree := New[int, string]()
		it := tree.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty tree")
		}
	}

	// PrevTo (not found)
	{
		tree := New[int, string]()
		tree.Put(0, "xx")
		tree.Put(1, "yy")
		it := tree.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty tree")
		}
	}

	// PrevTo (found)
	{
		tree := New[int, string]()
		tree.Put(2, "cc")
		tree.Put(0, "aa")
		tree.Put(1, "bb")
		it := tree.Iterator()
		it.End()
		if !it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty tree")
		}
		if index, value := it.Key(), it.Value(); index != 1 || value != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Prev() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Key(), it.Value(); index != 0 || value != "aa" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "aa")
		}
		if it.Prev() {
			t.Errorf("Should not go before first element")
		}
	}
}

//func TestRedBlackTreeSerialization(t *testing.T) {
//	tree := New[string, string]()
//	tree.Put("c", "3")
//	tree.Put("b", "2")
//	tree.Put("a", "1")
//
//	var err error
//	assert := func() {
//		if actualValue, expectedValue := tree.Size(), 3; actualValue != expectedValue {
//			t.Errorf("Got %v expected %v", actualValue, expectedValue)
//		}
//		if actualValue := tree.Keys(); actualValue[0] != "a" || actualValue[1] != "b" || actualValue[2] != "c" {
//			t.Errorf("Got %v expected %v", actualValue, "[a,b,c]")
//		}
//		if actualValue := tree.Values(); actualValue[0] != "1" || actualValue[1] != "2" || actualValue[2] != "3" {
//			t.Errorf("Got %v expected %v", actualValue, "[1,2,3]")
//		}
//		if err != nil {
//			t.Errorf("Got error %v", err)
//		}
//	}
//
//	assert()
//
//	bytes, err := tree.ToJSON()
//	assert()
//
//	err = tree.FromJSON(bytes)
//	assert()
//
//	bytes, err = json.Marshal([]interface{}{"a", "b", "c", tree})
//	if err != nil {
//		t.Errorf("Got error %v", err)
//	}
//
//	err = json.Unmarshal([]byte(`{"a":1,"b":2}`), &tree)
//	if err != nil {
//		t.Errorf("Got error %v", err)
//	}
//}

func TestRedBlackTreeString(t *testing.T) {
	c := New[string, int]()
	c.Put("a", 1)
	if !strings.HasPrefix(c.String(), "RedBlackTree") {
		t.Errorf("String should start with container name")
	}
}

func benchmarkGet(b *testing.B, tree *Tree[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, tree *Tree[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Put(n, struct{}{})
		}
	}
}

func benchmarkRemove(b *testing.B, tree *Tree[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Remove(n)
		}
	}
}

func BenchmarkRedBlackTreeGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreePut100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := New[int, struct{}]()
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreeRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := New[int, struct{}]()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}
