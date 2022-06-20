package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (this *Node) Insert(k int) {
	if k < this.Key {
		// move left
		if this.Left == nil {
			this.Left = &Node{Key: k}
		} else if this.Left != nil {
			this.Left.Insert(k)
		}
	} else if k > this.Key {
		// move right
		if this.Right == nil {
			this.Right = &Node{Key: k}
		} else if this.Right != nil {
			this.Right.Insert(k)
		}
	}
}

func (this *Node) Search(k int) bool {
	if this == nil {
		return false
	}

	if k < this.Key {
		// move left
		return this.Left.Search(k)
	} else if k > this.Key {
		// move right
		return this.Right.Search(k)
	}

	return true
}

func main() {
	Tree := &Node{Key: 100}

	Tree.Insert(50)
	Tree.Insert(150)
	Tree.Insert(25)
	Tree.Insert(75)
	Tree.Insert(200)

	fmt.Println(Tree.Search(51))
}
