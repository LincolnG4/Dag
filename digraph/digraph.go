package digraph

import "fmt"

// Create a Directed graph (Digraph)
type Digraph struct {
	nodes map[nodeID]*Node // Nodes connected to n by edges pointing from n
}

// Create the Digraph of nodes without edges. Let it empty to start empty Digraph
func NewDigraph(nodes ...*Node) (Digraph, error) {
	d := Digraph{
		nodes: make(map[nodeID]*Node),
	}

	err := d.AddNodes(nodes...)
	if err != nil {
		return Digraph{}, err
	}
	return d, nil
}

// Adds an edge from 'fromID' to 'toID' in the graph.
func (d *Digraph) AddEdge(fromID, toID nodeID) error {
	from, err := d.GetNodeByID(fromID)
	if err != nil {
		return err
	}

	to, err := d.GetNodeByID(toID)
	if err != nil {
		return err
	}

	err = from.Connect(to)
	if err != nil {
		return err
	}

	return nil
}

// check if node exist in dag
func (d *Digraph) NodeExists(id nodeID) bool {
	_, exist := d.nodes[id]
	return exist
}

// get node by id
func (d *Digraph) GetNodeByID(id nodeID) (*Node, error) {
	if !d.NodeExists(id) {
		return nil, fmt.Errorf("node '%s' not added to the graph", id)
	}
	return d.nodes[id], nil
}

// add 1 or more node into the diapraph
func (d *Digraph) AddNodes(nodes ...*Node) error {
	var err error
	for _, node := range nodes {
		err = d.addNode(node)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Digraph) addNode(node *Node) error {
	if node == nil {
		return fmt.Errorf("node is nil")
	}
	if _, exist := d.nodes[node.ID]; exist {
		return fmt.Errorf("node id '%s' already exist in the dag", node.ID)
	}
	d.nodes[node.ID] = node
	return nil
}

// remove node by id
func (d *Digraph) RemoveNodeByID(id nodeID) error {
	if !d.NodeExists(id) {
		return fmt.Errorf("node '%s' not in the dag", id)
	}
	delete(d.nodes, id)
	return nil
}

func (d *Digraph) Len() int {
	return len(d.nodes)
}

func HasCycle() bool {
	return false
}

func Cycle() {}
