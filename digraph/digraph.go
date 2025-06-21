package digraph

import "fmt"

// Create a Directed graph (Digraph)
type Digraph struct {
	nodes map[string]*Node // Nodes connected to n by edges pointing from n
}

// Create the Digraph of nodes without edges. Let it empty to start empty Digraph
func NewDigraph(nodes ...*Node) (Digraph, error) {
	d := Digraph{
		nodes: make(map[string]*Node),
	}

	err := d.AddNodes(nodes...)
	if err != nil {
		return Digraph{}, err
	}
	return d, nil
}

func (d *Digraph) GetAllNodes() []*Node {
	n := make([]*Node, 0, len(d.nodes))
	for _, v := range d.nodes {
		n = append(n, v)
	}
	return n
}

// Adds an edge from 'fromID' to 'toID' in the graph.
func (d *Digraph) AddEdge(fromID, toID string) error {
	from, err := d.GetNodeByID(fromID)
	if err != nil {
		return err
	}

	to, err := d.GetNodeByID(toID)
	if err != nil {
		return err
	}

	err = from.ConnectNode(to)
	if err != nil {
		return err
	}

	return nil
}

// check if node exist in dag
func (d *Digraph) NodeExists(id string) bool {
	_, exist := d.nodes[id]
	return exist
}

// get node by id
func (d *Digraph) GetNodeByID(id string) (*Node, error) {
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
func (d *Digraph) RemoveNodeByID(id string) error {
	if !d.NodeExists(id) {
		return fmt.Errorf("node '%s' not in the dag", id)
	}
	delete(d.nodes, id)
	return nil
}

// Remove an edge from 'fromID' to 'toID' in the graph.
func (d *Digraph) RemoveEdgeByID(fromID, toID string) error {
	from, err := d.GetNodeByID(fromID)
	if err != nil {
		return fmt.Errorf("from node %s not found: %w", fromID, err)
	}

	to, err := d.GetNodeByID(toID)
	if err != nil {
		return fmt.Errorf("to node %s not found: %w", toID, err)
	}

	err = from.DisconnectNode(to)
	if err != nil {
		return fmt.Errorf("to node %s not found: %w", toID, err)
	}

	return nil
}

func (d *Digraph) Len() int {
	return len(d.nodes)
}

func (d *Digraph) HasCycle() bool {
	return false
}

func (d *Digraph) Cycle() []*Node {
	return nil
}

// Calculate the inDegree of each vertex
func (d *Digraph) inDegree() map[string]int {
	inDegree := make(map[string]int)
	for _, from := range d.GetAllNodes() {
		inDegree[from.ID] = 0

	}
	for _, from := range d.GetAllNodes() {
		for _, to := range from.edgeTo {
			inDegree[to.ID]++
		}

	}

	return inDegree
}

// Kahn's algorithm for Topological Sorting
func (d *Digraph) TopologicalSort() ([]*Node, error) {
	inDegree := d.inDegree()
	q := make([]*Node, 0)

	// add to indegree 0 to stack
	for id, v := range inDegree {
		if v == 0 {
			q = append(q, d.nodes[id])
		}
	}

	res := make([]*Node, 0)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			break
		}

		res = append(res, node)

		for _, neighbor := range node.edgeTo {
			inDegree[neighbor.ID]--
			if inDegree[neighbor.ID] == 0 {
				q = append(q, neighbor)
			}
		}
	}

	if len(res) != len(d.nodes) {
		return nil, fmt.Errorf("graph has a cycle, no topological sort possible")
	}
	return res, nil
}
