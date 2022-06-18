// DATA STRUCTURES
// ARRAY

package main

import (
	"errors"
	"fmt"
	"log"
)

type Larik struct {
	Data []int
}

func main() {
	Data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Create new Larik
	myData := NewLarik(Data)
	myData.Rotate(11)
	fmt.Println("result of data rotated", myData)
}

func NewLarik(Data []int) Larik {
	myData := Larik{Data: Data}
	return myData
}

func (this *Larik) Rotate(n int) {

	// Check, it will return an error if amount of data exceed existing data
	if n > len(this.Data) {
		err := errors.New("Amount of data you want to rotated exceeded existing data")
		log.Println(err)
		return
	}
	tmpData := make([]int, n)

	// Insert amount of data you want to rotate to tmpData from myData
	for i := 0; i < n; i++ {
		tmpData[i] = this.Data[i]
	}
	fmt.Println(tmpData)

	// Shift remaining data of myData
	for i := 0; i < len(this.Data); i++ {
		if (n + i) == len(this.Data) {
			break
		}
		this.Data[i] = this.Data[n+i]
	}
	fmt.Println(this.Data)

	// Insert tmpData to myData
	for i := 0; i < n; i++ {
		this.Data[len(this.Data)-n+i] = tmpData[i]
	}
}
