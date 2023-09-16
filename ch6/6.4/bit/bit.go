package bit

import (
	"bytes"
	"fmt"
)

const (
	uint64Max uint64 = 18446744073709551615
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func lowbit(x uint64) uint64 {
	return x & -x
}

func (s *IntSet) Len() int {
	res := 0

	for _, val := range s.words {
		for val > 0 {
			val -= lowbit(val)
			res++
		}
	}

	return res
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, x%64

	if len(s.words) <= word {
		return
	}

	mask := uint64Max ^ (uint64(1) << uint64(bit))
	s.words[word] &= mask
}

func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

func (s *IntSet) Copy() *IntSet {
	return s
}

func (s *IntSet) AddAll(alls ...int) {
	for _, x := range alls {
		s.Add(x)
	}
}

func (s *IntSet) IntersectWith(x *IntSet) {
	for i, val := range s.words {
		for j := 0; j < 64; j++ {
			if ((val >> j) & 1) == 1 {
				num := i*64 + j
				if !x.Has(num) {
					s.Remove(num)
				}
			}
		}
	}
}

func (s *IntSet) DifferenceWith(x *IntSet) {
	for i, val := range s.words {
		for j := 0; j < 64; j++ {
			if ((val >> j) & 1) == 1 {
				num := i*64 + j
				if x.Has(num) {
					s.Remove(num)
				}
			}
		}
	}
}

func (x *IntSet) SymmetricDifference(y *IntSet) *IntSet {
	var res IntSet

	for i, val := range x.words {
		for j := 0; j < 64; j++ {
			if ((val >> j) & 1) == 1 {
				num := i*64 + j
				if !y.Has(num) {
					res.Add(num)
				}
			}
		}
	}

	for i, val := range y.words {
		for j := 0; j < 64; j++ {
			if ((val >> j) & 1) == 1 {
				num := i*64 + j
				if !x.Has(num) {
					res.Add(num)
				}
			}
		}
	}

	return &res
}

func (s *IntSet) Elems() []int {
	var res []int

	for i, val := range s.words {
		for j := 0; j < 64; j++ {
			if (val >> j & 1) == 1 {
				res = append(res, i*64+j)
			}
		}
	}

	return res
}
