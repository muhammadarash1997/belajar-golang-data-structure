package main

import "fmt"

type Heap struct {
	array []int
}

func (this *Heap) Push(n int) {
	this.array = append(this.array, n)
	this.heapifyUp(len(this.array)-1)
}

func (this *Heap) Pop() int {
	if len(this.array) == 0 {
		fmt.Println("cannot pop because array length is 0")
		return -1
	}

	n := this.array[0]
	l := len(this.array)-1

	this.array[0] = this.array[l]
	this.array = this.array[:l]

	this.heapifyDown(0)
	return n
}

func (this *Heap) heapifyUp(index int) {
	for this.array[parent(index)] < this.array[index] {
		this.swap(parent(index), index)
		index = parent(index)
	}
}

func (this *Heap) heapifyDown(index int) {
	
	lastIndex := len(this.array)-1
	l, r := left(index), right(index)
	var childTocompare int

	for l <= lastIndex {
		if l == lastIndex {	// when left child is the only child
			childTocompare = l
		} else if this.array[l] > this.array[r] {	// when left child is larger
			childTocompare = l
		} else { // when right child is larger
			childTocompare = r
		}

		if this.array[index] < this.array[childTocompare] {
			this.swap(index, childTocompare)
			index = childTocompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (this *Heap) swap(i, j int) {
	this.array[i], this.array[j] = this.array[j], this.array[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func main() {
	myHeap := &Heap{}

	array := []int{5, 12, 20, 36, 4, 11, 9, 5, 25}
	for _, v := range array {
		myHeap.Push(v)
		fmt.Println(myHeap)
	}

	for i := 0; i < 5; i++ {
		myHeap.Pop()
		fmt.Println(myHeap)
	}
}