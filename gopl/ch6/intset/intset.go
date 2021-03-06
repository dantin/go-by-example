// Package intset is a implementation  a set of small non-negative integer.
// version 1.0 2018-08-13
package intset

import (
	"bytes"
	"fmt"
)

// wordSize is the word size on machine.
const wordSize = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	// To encapsulate an object, we must make it a struct.
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/wordSize, uint(x%wordSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/wordSize, uint(x%wordSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll add all the nums to the set.
func (s *IntSet) AddAll(nums ...int) {
	for _, n := range nums {
		s.Add(n)
	}
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

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// DifferenceWith set s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// SymmetricDifference set s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
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
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", wordSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func popcount(x uint) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}

// Len returns the number of elements.
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popcount(word)
	}
	return count
}

// Remove removes x from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/wordSize, uint(x%wordSize)
	s.words[word] &^= 1 << bit
}

// Clear removes all elements from the set.
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// Copy returns a copy of the set.
func (s *IntSet) Copy() *IntSet {
	newOne := &IntSet{}
	newOne.words = make([]uint, len(s.words))
	copy(newOne.words, s.words)
	return newOne
}

// Elems returns elements of int slice.
func (s *IntSet) Elems() []int {
	e := make([]int, 0)
	for i, word := range s.words {
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				e = append(e, i*wordSize+j)
			}
		}
	}
	return e
}
