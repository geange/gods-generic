package btree

import "fmt"

// Entry represents the key-value pair contained within nodes
type Entry[K, V any] struct {
	Key   K
	Value V
}

// String entry format to string
func (entry *Entry[K, V]) String() string {
	return fmt.Sprintf("%v", entry.Key)
}
