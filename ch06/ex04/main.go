package main

import (
	"bytes"
	"fmt"
	"math/bits"
)

// IntSetは負ではない小さな整数のセットです。
// そのゼロ値は空セットを表しています。
type IntSet struct {
	words []uint64
}

// Hasは負ではない値xをセットが含んでいるか否かを報告します。
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Addはセットに負ではない値xを追加します。
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWithは、sとtの和集合をsに設定します。
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Stringは"{1 2 3}"の形式の文字列としてセットを返します。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 { // そのwordの中で立っているbitが無い。つまり集合に含まれる値はないので読み飛ばす
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") { // 最初の文字を書き込むときにスペースをいれたくなくてこうしてる
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// 要素数を返します
func (s *IntSet) Len() int {
	l := 0
	for _, w := range s.words {
		l += bits.OnesCount64(w)

	}
	return l
}

// セットからxを取り除きます
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] &^= 1 << bit
}

// セットからすべての要素を取り除きます
func (s *IntSet) Clear() {
	s.words = nil
}

// セットのコピーを返します
func (s *IntSet) Copy() *IntSet {
	var r IntSet
	for _, w := range s.words {
		r.words = append(r.words, w)
	}
	return &r
}

func (s *IntSet) AddAll(vals ...int) {
	for _, v := range vals {
		s.Add(v)
	}
}

// 積集合をsに設定します
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, _ := range s.words {
		if len(t.words)-1 < i {
			s.words[i] = 0
		} else {
			s.words[i] &= t.words[i]
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	u := s.Copy()
	u.IntersectWith(t)
	for i, _ := range u.words {
		u.words[i] = ^u.words[i]
	}
	s.IntersectWith(u)
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	u := s.Copy()
	u.IntersectWith(t)
	s.UnionWith(t)
	s.DifferenceWith(u)
}

func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 { // そのwordの中で立っているbitが無い。つまり集合に含まれる値はないので読み飛ばす
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 64*i+j)
			}
		}
	}
	return elems
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
}
