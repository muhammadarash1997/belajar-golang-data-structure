package main

// Hash Table is actually combination of Array and Linked List

// HashTable is for store array
type HashTable struct {
	array [26]*bucket
}

// bucket is for store head of linked list
type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key string
	next *bucketNode
}