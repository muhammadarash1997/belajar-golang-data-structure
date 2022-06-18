package main

import "fmt"

type Stack struct {
	array []int
}

func (this *Stack) Push(n int) {
	this.array = append(this.array, n)
}

func (this *Stack) Pop() int {
	n := this.array[len(this.array)-1]
	this.array = this.array[:len(this.array)-1]
	return n
}

func main() {
	myStack := Stack{}

	myStack.Push(4)
	myStack.Push(6)
	myStack.Push(2)

	fmt.Println(myStack.array)

	myStack.Pop()

	fmt.Println(myStack.array)
}
