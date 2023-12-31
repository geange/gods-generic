// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/geange/gods-generic/cmp"
	"github.com/geange/gods-generic/lists/arraylist"
)

// ArrayListExample to demonstrate basic usage of ArrayList
func main() {
	eq := func(a, b string) bool {
		return a == b
	}

	list := arraylist.New[string]()
	list.Add("a")                             // ["a"]
	list.Add("c", "b")                        // ["a","c","b"]
	list.Sort(cmp.Compare[string])            // ["a","b","c"]
	_, _ = list.Get(0)                        // "a",true
	_, _ = list.Get(100)                      // nil,false
	_ = list.Contains(eq, "a", "b", "c")      // true
	_ = list.Contains(eq, "a", "b", "c", "d") // false
	list.Swap(0, 1)                           // ["b","a",c"]
	list.Remove(2)                            // ["b","a"]
	list.Remove(1)                            // ["b"]
	list.Remove(0)                            // []
	list.Remove(0)                            // [] (ignored)
	_ = list.Empty()                          // true
	_ = list.Size()                           // 0
	list.Add("a")                             // ["a"]
	list.Clear()                              // []
}
