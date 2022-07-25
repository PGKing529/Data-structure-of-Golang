package main

import (
	"fmt"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	tree := rbt.NewWithIntComparator()
	tree.Put(1, 'a')
	tree.Put(2, 'b')
	tree.Put(3, 'c')
	tree.Put(4, 'd')
	tree.Put(5, 'e')

	fmt.Println(tree.Left())

	fmt.Println(tree.Right())

	fmt.Println(tree.Ceiling(2))
}
