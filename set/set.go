package main

import (
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 2, 3, 4}

	set := hashset.New()
	for _, v := range a {
		set.Add(v)
	}
	fmt.Println(set.Size())
	fmt.Println(set.Contains(1))
	set.Remove(1)
	fmt.Println(set.Contains(1))
	fmt.Println(set.Values())

}
