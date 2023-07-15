//go:build !go1.18 && !go1.19 && !go1.20

package cmp

import (
	"cmp"
)

// Ordered use cmp.Ordered
type Ordered cmp.Ordered

// Compare use cmp.Compare
type Compare cmp.Compare
