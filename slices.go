package main

import (
	"fmt"
	"slices"
)

func main() {
	// [6] => array | [] => slice
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// slice: dynamic memory
	// make(type, length, capacity)
	s = make([]string, 3)
	// it's dynamic but it has capacity that we can preset
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	
	// put append's return val back in s, because we may 
	// get a different slice as the memory changes
	s = append(s, "d")
	s = append(s, "e", "f") // you can append >1 elems
	fmt.Println("apd:", s)

	// copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	
	// 2 incl, 5 excl (python)
	// but unlike python, it shares the same memory
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// array to slice through slicing
	x := [3]string{"holi1", "holi2", "holi3"}
	s2 := x[:]

	// grow s to its capacity by slicing
	s = s[:cap(s)]

	// NOTE: slices cannot be grown beyong its capacity by themselves
	// they need copy/append

	// append one slice to another
	a := []string{"john", "paul"}
	b := []string{"george", "ringo", "pete"}
	a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"

	// length of inner slices can vary
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		// you allocate inside the for
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
