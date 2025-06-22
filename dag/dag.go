package dag

import (
	"fmt"

	"github.com/google/uuid"
)

// Dag create a Directed Acyclic graph (Dag)
type Dag struct {
	nodes    map[string]*Node  // Nodes connected to n by edges pointing from n
	nameToID map[string]string // key: Name, value: uuid
}

// Create the Dag of nodes without edges. Let it empty to start empty Dag
func NewDag(nodes ...*Node) (Dag, error) {
	d := Dag{
		nodes:    make(map[string]*Node),
		nameToID: make(map[string]string),
	}

	err := d.AddNodes(nodes...)
	if err != nil {
		return Dag{}, err
	}
	return d, nil
}

func (d *Dag) NewNode(name string, value int) (*Node, error) {
	node := &Node{
		id:    uuid.New().String(),
		Name:  name,
		Value: value,
	}

	if err := d.AddNodes(node); err != nil {
		return nil, err
	}

	return node, nil
}

func (d *Dag) GetAllNodes() []*Node {
	n := make([]*Node, 0, len(d.nodes))
	for _, v := range d.nodes {
		n = append(n, v)
	}
	return n
}

// Adds an edge from 'fromid' to 'toid' in the graph.
func (d *Dag) AddEdgeByID(fromid, toid string) error {
	from, err := d.GetNodeByID(fromid)
	if err != nil {
		return err
	}

	to, err := d.GetNodeByID(toid)
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
func (d *Dag) NodeExists(id string) bool {
	_, exist := d.nodes[id]
	return exist
}

func (d *Dag) AddEdgeByName(fromName, toName string) error {
	from, err := d.GetNodeByName(fromName)
	if err != nil {
		return err
	}
	to, err := d.GetNodeByName(toName)
	if err != nil {
		return err
	}
	return from.ConnectNode(to)
}

func (d *Dag) GetNodeByName(name string) (*Node, error) {
	id, ok := d.nameToID[name]
	if !ok {
		return nil, fmt.Errorf("node with name '%s' not found", name)
	}
	return d.GetNodeByID(id)
}

// get node by id
func (d *Dag) GetNodeByID(id string) (*Node, error) {
	if !d.NodeExists(id) {
		return nil, fmt.Errorf("node '%s' not added to the graph", id)
	}
	return d.nodes[id], nil
}

// add 1 or more node into the diapraph
func (d *Dag) AddNodes(nodes ...*Node) error {
	var err error
	for _, node := range nodes {
		err = d.addNode(node)
		if err != nil {
			return err
		}
		d.nameToID[node.Name] = node.id
	}
	return nil
}

func (d *Dag) addNode(node *Node) error {
	if node == nil {
		return fmt.Errorf("node is nil")
	}
	if _, exist := d.nodes[node.id]; exist {
		return fmt.Errorf("node id '%s' already exist in the dag", node.id)
	}
	d.nodes[node.id] = node
	return nil
}

// remove node by id
func (d *Dag) RemoveNodeByID(id string) error {
	if !d.NodeExists(id) {
		return fmt.Errorf("node '%s' not in the dag", id)
	}

	// Remove all edges to this node
	for _, node := range d.nodes {
		node.DisconnectNode(d.nodes[id])
	}

	delete(d.nodes, id)
	return nil
}

// Remove an edge from 'fromid' to 'toid' in the graph.
func (d *Dag) RemoveEdgeByID(fromid, toid string) error {
	from, err := d.GetNodeByID(fromid)
	if err != nil {
		return fmt.Errorf("from node %s not found: %w", fromid, err)
	}

	to, err := d.GetNodeByID(toid)
	if err != nil {
		return fmt.Errorf("to node %s not found: %w", toid, err)
	}

	err = from.DisconnectNode(to)
	if err != nil {
		return fmt.Errorf("to node %s not found: %w", toid, err)
	}

	return nil
}

func (d *Dag) Len() int {
	return len(d.nodes)
}

func (d *Dag) HasCycle() bool {
	return len(d.FindCycle()) > 0
}

func (d *Dag) FindCycle() []*Node {
	visited := make(map[string]bool)
	onStack := make(map[string]bool)
	parent := make(map[string]string)
	var cycle []string

	var dfs func(n *Node) bool
	dfs = func(n *Node) bool {
		id := n.id
		visited[id] = true
		onStack[id] = true

		for _, neighbor := range n.edgeTo {
			nid := neighbor.id

			if !visited[nid] {
				parent[nid] = id
				if dfs(neighbor) {
					return true
				}
			} else if onStack[nid] {
				// Cycle detected
				cycle = []string{}
				for x := id; x != nid; x = parent[x] {
					cycle = append([]string{x}, cycle...)
				}
				cycle = append([]string{nid}, cycle...)
				cycle = append(cycle, nid)
				return true
			}
		}
		onStack[id] = false
		return false
	}

	for _, node := range d.GetAllNodes() {
		if !visited[node.id] {
			if dfs(node) {
				break
			}
		}
	}

	result := []*Node{}
	for _, id := range cycle {
		if n, err := d.GetNodeByID(id); err == nil {
			result = append(result, n)
		}
	}
	return result
}

func (d *Dag) Validate() error {
	_, err := d.TopologicalSort()
	return err
}

// Calculate the inDegree of each vertex
func (d *Dag) inDegree() map[string]int {
	inDegree := make(map[string]int)
	for _, from := range d.GetAllNodes() {
		inDegree[from.id] = 0

	}

	for _, from := range d.GetAllNodes() {
		for _, to := range from.edgeTo {
			inDegree[to.id]++
		}

	}
	return inDegree
}

// Kahn's algorithm for Topological Sorting
func (d *Dag) TopologicalSort() ([]*Node, error) {
	if d.HasCycle() {
		return nil, fmt.Errorf("graph has a cycle, no topological sort possible")
	}

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
			inDegree[neighbor.id]--
			if inDegree[neighbor.id] == 0 {
				q = append(q, neighbor)
			}
		}
	}

	if len(res) != len(d.nodes) {
		return nil, fmt.Errorf("graph has a cycle, no topological sort possible")
	}
	return res, nil
}
