package diagraph

import "fmt"

// Create a Directed graph (diagraph)
type Diagraph struct {
	nodes map[nodeID]*Node // Nodes connected to n by edges pointing from n
}

// Create the diagraph of nodes without edges. Let it empty to start empty diagraph
func NewDiagraph(nodes ...*Node) (Diagraph, error) {
	d := Diagraph{
		nodes: make(map[nodeID]*Node, 0),
	}

	err := d.AddNodes(nodes...)
	if err != nil {
		return Diagraph{}, err
	}
	return d, nil
}

// adds edge from->to this digraph
func (d *Diagraph) AddEdge(fromID, toID nodeID) error {
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
func (d *Diagraph) CheckIfNodeExist(id nodeID) bool {
	_, exist := d.nodes[id]
	if !exist {
		return false
	}
	return true
}

// get node by id
func (d *Diagraph) GetNodeByID(id nodeID) (*Node, error) {
	if !d.CheckIfNodeExist(id) {
		return nil, fmt.Errorf("node %s not add to dag", id)
	}
	return d.nodes[id], nil
}

// add 1 or more node into the diapraph
func (d *Diagraph) AddNodes(nodes ...*Node) error {
	var err error
	for _, node := range nodes {
		err = d.add(node)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Diagraph) add(node *Node) error {
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
func (d *Diagraph) RemoveNodeByID(id nodeID) error {
	if !d.CheckIfNodeExist(id) {
		return fmt.Errorf("node '%s' not in the dag", id)
	}
	delete(d.nodes, id)
	return nil
}

func (d *Diagraph) Len() int {
	return len(d.nodes)
}

func HasCycle() bool {
	return false
}

func Cycle() {}
