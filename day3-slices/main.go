package main

import "fmt"

func main() {
	// s := []int{1, 2, 3}
	// s = append(s, 4)
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(s[0])
	// var t []int
	// fmt.Println(t == nil)
	// t = append(t, 9)
	// fmt.Println(t)

	// s := make([]int, 2, 4)
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))
	// s = append(s, 7)
	// fmt.Println(s)

	a := []int{1, 2, 3} // len 3, cap 3 (full)
	b := append(a, 4)   // no room, so Go reallocates
	b[0] = 100
	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [100 2 3 4]
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))

	s := make([]int, 3, 5) // [0 0 0], cap 5
	t := append(s, 4)      // room, so same array
	s[0] = 100
	fmt.Println(s) // [100 0 0]  (a changed!)
	fmt.Println(t) // [100 0 0 4] (b changed!)
	fmt.Println(len(s), cap(s))
	fmt.Println(len(t), cap(t))

}
