package main

import "fmt"

type Queue struct {
	array []int
}

func (this *Queue) Push(n int) {
	this.array = append(this.array, n)
}

func (this *Queue) Pop() int {
	n := this.array[0]
	this.array = this.array[1:]
	return n
}

func main() {
	myQueue := Queue{}

	myQueue.Push(4)
	myQueue.Push(6)
	myQueue.Push(2)

	fmt.Println(myQueue.array)

	myQueue.Pop()

	fmt.Println(myQueue.array)
}
