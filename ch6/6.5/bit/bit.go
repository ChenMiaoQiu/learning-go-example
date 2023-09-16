package bit

import (
	"bytes"
	"fmt"
)

const (
	platformLimit uint = 32 << (^uint(0) >> 63)
	numMax        uint = ^uint(0)
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/int(platformLimit), uint(x%int(platformLimit))
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/int(platformLimit), uint(x%int(platformLimit))
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
		for j := 0; j < int(platformLimit); j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", int(platformLimit)*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func lowbit(x uint) uint {
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
	word, bit := x/int(platformLimit), x%int(platformLimit)

	if len(s.words) <= word {
		return
	}

	mask := numMax ^ (1 << bit)
	s.words[word] &= mask
}

func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
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
		for j := 0; j < int(platformLimit); j++ {
			if ((val >> j) & 1) == 1 {
				num := i*int(platformLimit) + j
				if !x.Has(num) {
					s.Remove(num)
				}
			}
		}
	}
}

func (s *IntSet) DifferenceWith(x *IntSet) {
	for i, val := range s.words {
		for j := 0; j < int(platformLimit); j++ {
			if ((val >> j) & 1) == 1 {
				num := i*int(platformLimit) + j
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
		for j := 0; j < int(platformLimit); j++ {
			if ((val >> j) & 1) == 1 {
				num := i*int(platformLimit) + j
				if !y.Has(num) {
					res.Add(num)
				}
			}
		}
	}

	for i, val := range y.words {
		for j := 0; j < int(platformLimit); j++ {
			if ((val >> j) & 1) == 1 {
				num := i*int(platformLimit) + j
				if !x.Has(num) {
					res.Add(num)
				}
			}
		}
	}

	return &res
}
