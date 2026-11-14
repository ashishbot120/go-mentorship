package main

import (
	"fmt"
	// "strings"
	// "sort"
)

// func wordCount(text string) map[string]int {
// 	words := strings.Fields(text)
// 	count := make(map[string]int)
// 	for _, word := range words {
// 		count[word]++
// 	}
// 	return count
// }

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

	// a := []int{1, 2, 3} // len 3, cap 3 (full)
	// b := append(a, 4)   // no room, so Go reallocates
	// b[0] = 100
	// fmt.Println(a) // [1 2 3]
	// fmt.Println(b) // [100 2 3 4]
	// fmt.Println(len(a), cap(a))
	// fmt.Println(len(b), cap(b))

	// s := make([]int, 3, 5) // [0 0 0], cap 5
	// t := append(s, 4)      // room, so same array
	// s[0] = 100
	// fmt.Println(s) // [100 0 0]  (a changed!)
	// fmt.Println(t) // [100 0 0 4] (b changed!)
	// fmt.Println(len(s), cap(s))
	// fmt.Println(len(t), cap(t))

	// ages := map[string]int{"Asha": 30}
	// ages["Ravi"] = 23
	// delete(ages, "Asha")
	// fmt.Println(len(ages))
	// v, ok := ages["Asha"]
	// fmt.Println(v, ok)

	// nums := []int{5, 6, 7}
	// total := 0
	// for i, n := range nums {
	// 	fmt.Println(i, n)
	// 	total += n
	// }
	// fmt.Println("Total:", total)

	// text := "go is fun and go is fast and go is simple"
	// count := wordCount(text)

	// sortedWords := make([]string, 0, len(count))
	// for word := range count {
	// 	sortedWords = append(sortedWords, word)
	// }
	// sort.Strings(sortedWords)
	// for _, word := range sortedWords {
	// 	fmt.Printf("%s: %d\n", word, count[word])
	// }

	s := []int{1, 2, 3}
	t := s
	t[0] = 9
	s = append(s, 4)
	s[1] = 7
	fmt.Println(s, t)

}
