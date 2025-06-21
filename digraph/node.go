package digraph

import (
	"fmt"
	"slices"
)

type Node struct {
	ID    string
	Value int

	edgeTo []*Node // node points to
}

// Create a new node
func NewNode(id string, value int) *Node {
	return &Node{
		ID:    id,
		Value: value,
	}
}

func (n *Node) EdgeTo() []*Node {
	return n.edgeTo
}

func (n *Node) ConnectNode(to *Node) error {
	if to == nil {
		return fmt.Errorf("cannot connect to a nil node")
	}

	if n.ID == to.ID {
		return fmt.Errorf("cannot connect node to itself")
	}

	if slices.Contains(n.edgeTo, to) {
		return fmt.Errorf("from %s already contains %s", n.ID, to.ID)
	}
	n.edgeTo = append(n.edgeTo, to)
	return nil
}

func (n *Node) DisconnectNode(to *Node) error {
	if to == nil {
		return fmt.Errorf("cannot disconnect to a nil node")
	}

	index := -1
	for i, node := range n.edgeTo {
		if node == to {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("edge from %s to %s does not exist", n.ID, to.ID)
	}

	// Remove edge from slice
	n.edgeTo = append(n.edgeTo[:index], n.edgeTo[index+1:]...)
	return nil
}
