// DATA STRUCTURES
// ARRAY

package main

import (
	"errors"
	"fmt"
	"log"
)

type Larik []int

func main() {

	// Create new Larik
	myData := Larik{1, 1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 9, 9, 10, 10}
	myData.Rotate(4)
	myData.Distinct()
	fmt.Println("result of data rotated", myData)
}

func (this *Larik) Rotate(n int) {

	// Check, it will return an error if amount of data (n) exceed existing data
	if n > len(*this) {
		err := errors.New("Amount of data you want to rotated exceeded existing data")
		log.Println(err)
		return
	}
	tmpData := make([]int, n)

	// Insert n data you want to rotate to tmpData from myData
	for i := 0; i < n; i++ {
		tmpData[i] = (*this)[i]
	}

	// Shift remaining data of myData
	for i := 0; i < len(*this); i++ {

		// jika n+i melebihi panjang dari slice maka berhenti
		if (n + i) == len(*this) {
			break
		}

		// geser nilai yang ada di index n+i ke i
		(*this)[i] = (*this)[n+i]
	}

	// Insert tmpData to myData
	for i := 0; i < n; i++ {

		// masukkan seluruh data yang ada di tmpData ke belakang slice kita
		(*this)[len(*this)-n+i] = tmpData[i]
	}
}

func (this *Larik) Distinct() {

	// buat penyimpanan sementara untuk menyimpan data yang tidak memiliki kembaran
	var tmpData []int

	// cek satu persatu
	for i := 0; i < len(*this); i++ {
		
		// cek satu persatu
		for j := 0; j < len(*this); j++ {
			
			if i != j {
				// cek apakah data sudah ada di tmpData, jika ya maka abaikan dan jika tidak maka tambahkan
				if (*this)[i] != (*this)[j] && isNotAlreadyExist((*this)[i], tmpData) {
					tmpData = append(tmpData, (*this)[i])
				}
			}
		}
	}

	*this = tmpData
}

// This function is for checking value whether or not already exist in
// your slice, it will return true if the value is already exist and
// vice versa
//	isNotAlreadyExist(value, yourSlice)
func isNotAlreadyExist(val int, yourSlice []int) bool {
	
	for i := 0; i < len(yourSlice); i++ {

		// jika value sudah ada di slice kita maka return false
		if val == yourSlice[i] {
			return false
		}
	}

	return true
}