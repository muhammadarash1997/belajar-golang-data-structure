//     matrix := [][]int{
// 		{0, 0, 0, 0},	// <-- vertex pertama dimana belum ada edge
// 		{0, 0, 0, 0},	// <-- vertex kedua dimana belum ada edge
// 		{0, 0, 0, 0},	// <-- vertex ketiga dimana belum ada edge
// 		{0, 0, 0, 0},	// <-- vertex keempat dimana belum ada edge
// 	}

package main

import "fmt"

func InitGraph(size int) [][]int {
	myVertex := []int{}
	for i := 0; i < size; i++ {
		myVertex = append(myVertex, 0)
	}

	var myGraph [][]int
	for i := 0; i < size; i++ {
		myGraph = append(myGraph, myVertex)
	}
	return myGraph
}

type Graph [][]int

func main() {
	// Membuat graph
	// myGraph := [][]int{{}}	<-- kalau seperti ini akan terjadi error di fungsi penambahan vertex, jadi ketika membuah matrix baru minimal harus memasukkan 1 angka seperti ini [][]int{{0}}
	myGraph := InitGraph(5)

	fmt.Println(myGraph)

	myGraph = AddVertex(myGraph)

	fmt.Println(myGraph)

	
}

func AddVertex(myGraph [][]int) [][]int {
	length := len(myGraph[0])

	// tambah 0 pada setiap akhir dari setiap baris
	for i := 0; i < length; i++ {
		myGraph[i] = append(myGraph[i], 0)
	}

	// buat sebuah array 1D dengan ukuran lebih besar 1 ukuran dari myGraph yang nanti array 1D ini akan ditambahkan ke akhir dari myGraph
	newVertex := []int{}
	for i := 0; i < length+1; i++ {
		newVertex = append(newVertex, 0)
	}

	// tambahkan array 1D di atas ke dalam myGraph
	myGraph = append(myGraph, newVertex)

	// kembalikan myGraph
	return myGraph
}