package main

import "fmt"

type Node struct {
	Children [26]*Node
	IsEnd    bool	// bisa saja kita tidak menggunakan IsEnd, tetapi yang jadi masalah jika kita sudah menyimpan kata seperti misal "orc", ketika kita search "or" maka dikatakan "or" tsb ada dalam memory kita padahal kita tidak pernah menyimpan kata tsb, maka dari itu kita harus menggunakan IsEnd
}

type Tree struct {
	Root *Node
}

func (this *Tree) Insert(word string) {
	currentNode := this.Root
	length := len(word)
	for i := 0; i < length; i++ {
		charIndex := word[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			currentNode.Children[charIndex] = &Node{}
		}
		currentNode = currentNode.Children[charIndex]
	}
	currentNode.IsEnd = true
}

func (this *Tree) Search(word string) bool {
	currentNode := this.Root
	length := len(word)
	for i := 0; i < length; i++ {
		charIndex := word[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.Children[charIndex]
	}
	if currentNode.IsEnd == true {
		return true
	}
	return false
}

func InitTree() *Tree {
	newTree := &Tree{Root: &Node{}}
	return newTree
}

func main() {
	myTree := InitTree()

	myTree.Insert("orc")
	myTree.Insert("orci")

	fmt.Println(myTree.Search("orc"))
	fmt.Println(myTree.Search("orci"))
	fmt.Println(myTree.Search("or"))
}
