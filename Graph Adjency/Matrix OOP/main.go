//     matrix := [][]int{
// 		{0, 0, 0, 0},	// <-- vertex pertama dimana belum ada edge
// 		{0, 0, 0, 0},	// <-- vertex kedua dimana belum ada edge
// 		{0, 0, 0, 0},	// <-- vertex ketiga dimana belum ada edge
// 		{0, 0, 0, 0},	// <-- vertex keempat dimana belum ada edge
// 	}

package main

import "fmt"

type Graph [][]int

func (this *Graph) AddVertex() {
	length := len((*this)[0])

	// tambah 0 pada setiap akhir dari setiap baris
	for i := 0; i < length; i++ {
		(*this)[i] = append((*this)[i], 0)
	}

	// buat sebuah array 1D dengan ukuran lebih besar 1 ukuran dari myGraph yang nanti array 1D ini akan ditambahkan ke akhir dari myGraph
	newVertex := []int{}
	for i := 0; i < length+1; i++ {
		newVertex = append(newVertex, 0)
	}
	
	// tambahkan array 1D di atas ke dalam myGraph
	*this = append(*this, newVertex)
}

func (this *Graph) DeleteVertex(index int) {
	*this = append((*this)[:index], (*this)[index+1:]...) // append(slicePenampung, dataYangInginDitambah)

	for i := 0; i < len((*this)[0]); i++ {
		(*this)[i] = RemoveIndex1D((*this)[i], index)
	}
}

func (this Graph) AddEdge(vertexOrigin int, vertexDestination int) {
	this[vertexOrigin][vertexDestination] = 1
}

func (this Graph) DeleteEdge(vertexOrigin int, vertexDestination int) {
	this[vertexOrigin][vertexDestination] = 0
}

func RemoveIndex1D(s []int, index int) []int {
	s = append(s[:index], s[index+1:]...)
	return s
}

func Create1D(size int, myGraph Graph) Graph {
	myVertex := []int{}
	for i := 0; i < size; i++ {
		myVertex = append(myVertex, 0)
	}
	myGraph = append(myGraph, myVertex)
	return myGraph
}

func InitGraph(size int) Graph {
	var myGraph Graph
	for i := 0; i < size; i++ {
		myGraph = Create1D(size, myGraph)
	}

	return myGraph
}

func main() {
	// Membuat graph
	// myGraph := [][]int{{}}	<-- kalau seperti ini akan terjadi error di fungsi penambahan vertex, jadi ketika membuah matrix baru minimal harus memasukkan 1 angka seperti ini [][]int{{0}}
	// myGraph := [][]int{{0}}	<-- kalau seperti ini tidak akan terjadi error di fungsi penambahan vertex, jadi ketika membuah matrix baru minimal harus memasukkan 1 angka seperti ini [][]int{{0}}
	myGraph := InitGraph(3)

	fmt.Println(myGraph)

	myGraph.AddVertex()

	fmt.Println(myGraph)
	
	myGraph.AddEdge(0, 3)

	fmt.Println(myGraph)
	
	myGraph.DeleteVertex(3)

	fmt.Println(myGraph)
}