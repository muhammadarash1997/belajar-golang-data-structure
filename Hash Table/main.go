package main

import "fmt"

// Hash Table is actually combination of Array and Linked List

const ArraySIze = 26

// HashTable is for store array
type HashTable struct {
	array [ArraySIze]*bucket
}

func (this *HashTable) Insert(key string) {
	index := hash(key)
	this.array[index].insert(key)
}

func (this *HashTable) Search(key string) bool {
	index := hash(key)
	return this.array[index].search(key)
}

// bucket is for store head of linked list
type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func (this *bucket) insert(key string) {
	newNode := &bucketNode{key: key}
	newNode.next = this.head
	this.head = newNode
}

func (this *bucket) search(key string) bool {
	current := this.head

	for current != nil {
		if current.key == key {
			return true
		}
		current = current.next
	}

	return false
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySIze
}

func main() {
	iHashTable := Init()

	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		iHashTable.Insert(v)
	}

	fmt.Println(iHashTable.Search("ERIC"))	// true
	fmt.Println(iHashTable.Search("Eric"))	// false
}
