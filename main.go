package main

import (
	"dag/dag"
	"fmt"
	"log"
)

func main() {
	// Define number of vertices and edge list
	V := 6
	edges := [][2]string{
		{"2", "3"},
		{"3", "1"},
		{"4", "0"},
		{"4", "1"},
		{"5", "0"},
		{"5", "2"},
		{"2", "5"},
	}

	// Create nodes with string IDs and dummy value = int(ID)
	nodes := make([]*dag.Node, 0, V)
	for i := 0; i < V; i++ {
		id := fmt.Sprintf("%d", i)
		node := dag.NewNode(id, i)
		nodes = append(nodes, node)
	}

	// Create the graph
	dg, err := dag.NewDag(nodes...)
	if err != nil {
		log.Fatalf("Failed to create dag: %v", err)
	}

	// Add edges
	for _, edge := range edges {
		fromID, toID := edge[0], edge[1]
		err := dg.AddEdge(fromID, toID)
		if err != nil {
			log.Fatalf("Error adding edge %s -> %s: %v", fromID, toID, err)
		}
	}

	fmt.Println("DAG constructed successfully.")
	fmt.Println("Nodes and their connections:")

	// Print the graph structure
	for _, node := range dg.GetAllNodes() {
		fmt.Printf("Node %s points to: ", node.ID)
		for _, neighbor := range node.EdgeTo() {
			fmt.Printf("%s ", neighbor.ID)
		}
		fmt.Println()
	}

	tp, _ := dg.TopologicalSort()
	for _, node := range tp {
		fmt.Print(node.ID, ",")
	}

	fmt.Println(dg.HasCycle())

}
