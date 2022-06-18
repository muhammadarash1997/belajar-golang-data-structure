package main

import (
	"errors"
	"fmt"
	"log"
)

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (this *Vertex) Print() {
	for _, vertex := range this.adjacent {
		fmt.Print(vertex.key, " ")
	}
}

type Graph struct {
	vertices []*Vertex
}

// Add vertex to the graph
func (this *Graph) AddVertex(k int) {
	if this.Contains(k) {
		return
	}
	this.vertices = append(this.vertices, &Vertex{key: k})
}

// Print all vertex in graph
func (this *Graph) Print() {
	for _, vertex := range this.vertices {
		fmt.Println(vertex)
	}
}

// Check whether or not input key already exist in graph
func (this *Graph) Contains(key int) bool {
	for _, vertex := range this.vertices {
		if key == vertex.key {
			err := errors.New("this key has been used")
			log.Println(err)
			return true
		}
	}
	return false
}

// Add edge
func (this *Graph) AddEdge(keyOrigin int, keyDestination int) {
	fromVertex := this.GetVertex(keyOrigin)
	toVertex := this.GetVertex(keyDestination)

	if fromVertex == nil || toVertex == nil {
		err := errors.New("key origin or key destination not-found")
		log.Println(err)
		return
	}
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
}

// Get vertex from graph
func (this *Graph) GetVertex(key int) *Vertex {
	for _, vertex := range this.vertices {
		if key == vertex.key {
			return vertex
		}
	}
	return nil
}

// Print all vertex along its all adjacent
func (this *Graph) PrintAllAdjacentOfEachVertex() {
	for _, vertex := range this.vertices {
		fmt.Print("Vertex ", vertex.key, " : ")
		vertex.Print()
		fmt.Println()
	}
}

// Create new graph
func InitGraph() *Graph {
	return &Graph{}
}

func main() {
	myGraph := InitGraph()
	
	myGraph.AddVertex(1)
	myGraph.AddVertex(2)
	myGraph.AddVertex(3)

	myGraph.AddEdge(1, 1)
	myGraph.AddEdge(1, 2)
	myGraph.AddEdge(1, 3)

	myGraph.AddEdge(2, 1)
	myGraph.AddEdge(2, 2)
	myGraph.AddEdge(2, 3)

	myGraph.AddEdge(3, 1)
	myGraph.AddEdge(3, 2)
	myGraph.AddEdge(3, 3)

	myGraph.Print()

	fmt.Println()

	myGraph.PrintAllAdjacentOfEachVertex()
}
