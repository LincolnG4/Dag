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
	}

	// Create the graph
	dg, err := dag.NewDag()
	if err != nil {
		log.Fatalf("Failed to create dag: %v", err)
	}

	// Create nodes with string ids and dummy value = int(id)
	nodes := make([]*dag.Node, 0, V)
	for i := 0; i < V; i++ {
		id := fmt.Sprintf("%d", i)
		node, err := dg.NewNode(id, i)
		if err != nil {
			log.Fatalf("Error adding node %s : %v", node.Name, err)
		}
		nodes = append(nodes, node)
	}

	// Add edges
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		err := dg.AddEdgeByName(from, to)
		if err != nil {
			log.Fatalf("Error adding edge %s -> %s: %v", from, to, err)
		}
	}

	fmt.Println("DAG constructed successfully.")
	fmt.Println("Nodes and their connections:")

	// Print the graph structure
	for _, node := range dg.GetAllNodes() {
		fmt.Printf("Node %s points to: ", node.Name)
		for _, neighbor := range node.EdgeTo() {
			fmt.Printf("%s ", neighbor.Name)
		}
		fmt.Println()
	}

	tp, _ := dg.TopologicalSort()
	for _, node := range tp {
		fmt.Print(node.Name, ",")
	}

	fmt.Println(dg.HasCycle())

}
