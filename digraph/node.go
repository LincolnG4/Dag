package digraph

import (
	"fmt"
	"slices"
)

type nodeID string

type Node struct {
	ID    nodeID
	Value int

	edgeTo []*Node // node points to
}

func (n *Node) Connect(to *Node) error {
	if to == nil {
		return fmt.Errorf("cannot connect to a nil node")
	}

	if slices.Contains(n.edgeTo, to) {
		return fmt.Errorf("from %s already contains %s", n.ID, to.ID)
	}
	n.edgeTo = append(n.edgeTo, to)
	return nil
}
