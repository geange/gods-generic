// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashbidimap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type option[K, V any] struct {
	key   K
	value V
	flag  bool
}

func TestMapPut(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	if actualValue := m.Size(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, m.Keys())
	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f", "g"}, m.Values())

	// key,expectedValue,expectedFound
	tests1 := []option[int, string]{
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
		actualValue, actualFound := m.Get(test.key)
		if actualValue != test.value || actualFound != test.flag {
			t.Errorf("Got %v expected %v", actualValue, test.flag)
		}
	}
}

func TestMapRemove(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	m.Remove(5)
	m.Remove(6)
	m.Remove(7)
	m.Remove(8)
	m.Remove(5)

	assert.Equal(t, []int{1, 2, 3, 4}, m.Keys())
	assert.Equal(t, []string{"a", "b", "c", "d"}, m.Values())
	assert.Equal(t, 4, m.Size())

	tests2 := []option[int, string]{
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
		actualValue, actualFound := m.Get(test.key)
		if actualValue != test.value || actualFound != test.flag {
			t.Errorf("Got %v expected %v", actualValue, test.value)
		}
	}

	m.Remove(1)
	m.Remove(4)
	m.Remove(2)
	m.Remove(3)
	m.Remove(2)
	m.Remove(2)

	assert.Empty(t, m.Keys())
	assert.Empty(t, m.Values())
	assert.Equal(t, 0, m.Size())
	assert.True(t, m.Empty())
}

func TestMapGetKey(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	// key,expectedValue,expectedFound
	tests1 := []option[int, string]{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{0, "x", false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := m.GetKey(test.value)
		assert.Equal(t, test.flag, actualFound)
		assert.Equal(t, test.key, actualValue)
		//if actualValue != test.key || actualFound != test.flag {
		//	t.Errorf("Got %v expected %v", actualValue, test.key)
		//}
	}
}

//func TestMapSerialization(t *testing.T) {
//	m := New[string, float64]()
//	m.Put("a", 1.0)
//	m.Put("b", 2.0)
//	m.Put("c", 3.0)
//
//	var err error
//	assert := func() {
//		if actualValue, expectedValue := m.Keys(), []interface{}{"a", "b", "c"}; !sameElements(actualValue, expectedValue) {
//			t.Errorf("Got %v expected %v", actualValue, expectedValue)
//		}
//		if actualValue, expectedValue := m.Values(), []interface{}{1.0, 2.0, 3.0}; !sameElements(actualValue, expectedValue) {
//			t.Errorf("Got %v expected %v", actualValue, expectedValue)
//		}
//		if actualValue, expectedValue := m.Size(), 3; actualValue != expectedValue {
//			t.Errorf("Got %v expected %v", actualValue, expectedValue)
//		}
//		if err != nil {
//			t.Errorf("Got error %v", err)
//		}
//	}
//
//	assert()
//
//	bytes, err := m.ToJSON()
//	assert()
//
//	err = m.FromJSON(bytes)
//	assert()
//
//	bytes, err = json.Marshal([]interface{}{"a", "b", "c", m})
//	if err != nil {
//		t.Errorf("Got error %v", err)
//	}
//
//	err = json.Unmarshal([]byte(`{"a":1,"b":2}`), &m)
//	if err != nil {
//		t.Errorf("Got error %v", err)
//	}
//}
//
//func TestMapString(t *testing.T) {
//	c := New()
//	c.Put("a", 1)
//	if !strings.HasPrefix(c.String(), "HashBidiMap") {
//		t.Errorf("String should start with container name")
//	}
//}

func sameElements(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func benchmarkGet(b *testing.B, m *Map[int, int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, m *Map[int, int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Put(n, n)
		}
	}
}

func benchmarkRemove(b *testing.B, m *Map[int, int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Remove(n)
		}
	}
}

func BenchmarkHashBidiMapGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashBidiMapGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashBidiMapGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashBidiMapGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashBidiMapPut100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, int]()
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashBidiMapPut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashBidiMapPut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashBidiMapPut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashBidiMapRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkHashBidiMapRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkHashBidiMapRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkHashBidiMapRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}
