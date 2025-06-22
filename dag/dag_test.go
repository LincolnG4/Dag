package dag

import (
	"testing"
)

func newTestNode(id string, val int) *Node {
	return &Node{ID: id, Value: val}
}

func TestAddNodes(t *testing.T) {
	t.Run("Add valid nodes", func(t *testing.T) {
		d, err := NewDag()
		if err != nil {
			t.Fatal(err)
		}
		n1 := newTestNode("node1", 1)
		n2 := newTestNode("node2", 1)

		err = d.AddNodes(n1, n2)
		if err != nil {
			t.Errorf("expected nodes to be added, got error: %s", err)
		}
	})

	t.Run("Add duplicate nodes should fail", func(t *testing.T) {
		n1 := newTestNode("node1", 1)
		d, _ := NewDag(n1)

		err := d.AddNodes(n1)
		if err == nil {
			t.Errorf("expected error for duplicate node")
		}
	})

	t.Run("NewDag with nil node should fail", func(t *testing.T) {
		_, err := NewDag(nil)
		if err == nil {
			t.Errorf("expected error for nil node")
		}
	})
}

func TestRemoveNodeByID(t *testing.T) {
	n1 := newTestNode("n1", 1)
	n2 := newTestNode("n2", 2)

	d, _ := NewDag(n1, n2)
	err := d.RemoveNodeByID("n2")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if d.Len() != 1 {
		t.Errorf("expected 1 node after removal, got %d", d.Len())
	}
}

func TestNodeExists(t *testing.T) {
	n := newTestNode("n", 10)
	d, _ := NewDag(n)

	t.Run("Existing node", func(t *testing.T) {
		if !d.NodeExists("n") {
			t.Errorf("expected node to exist")
		}
	})

	t.Run("Non-existent node", func(t *testing.T) {
		if d.NodeExists("missing") {
			t.Errorf("expected node to not exist")
		}
	})
}

func TestGetNodeByID(t *testing.T) {
	n1 := newTestNode("n1", 1)
	d, _ := NewDag(n1)

	t.Run("Get existing node", func(t *testing.T) {
		node, err := d.GetNodeByID("n1")
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if node.ID != "n1" {
			t.Errorf("expected ID n1, got %s", node.ID)
		}
	})

	t.Run("Get non-existent node", func(t *testing.T) {
		_, err := d.GetNodeByID("missing")
		if err == nil {
			t.Errorf("expected error for missing node")
		}
	})
}

func TestRemoveEdgeByID(t *testing.T) {
	n1 := &Node{ID: "n1", Value: 1}
	n2 := &Node{ID: "n2", Value: 2}
	n3 := &Node{ID: "n3", Value: 3}

	dg, err := NewDag(n1, n2, n3)
	if err != nil {
		t.Fatalf("failed to create dag: %v", err)
	}

	t.Run("Remove existing edge", func(t *testing.T) {
		err := dg.AddEdge(n1.ID, n2.ID)
		if err != nil {
			t.Fatalf("failed to add edge: %v", err)
		}

		err = dg.RemoveEdgeByID(n1.ID, n2.ID)
		if err != nil {
			t.Errorf("expected edge to be removed, got error: %v", err)
		}

		// Confirm edge is actually removed
		if len(n1.edgeTo) != 0 {
			t.Errorf("expected edgeTo to be empty, got %d edges", len(n1.edgeTo))
		}
	})

	t.Run("Remove non-existent edge", func(t *testing.T) {
		err := dg.RemoveEdgeByID(n1.ID, n2.ID)
		if err == nil {
			t.Errorf("expected error when removing non-existent edge")
		}
	})

	t.Run("Remove edge with non-existent from node", func(t *testing.T) {
		err := dg.RemoveEdgeByID("invalid", n2.ID)
		if err == nil {
			t.Errorf("expected error for invalid from node ID")
		}
	})

	t.Run("Remove edge with non-existent to node", func(t *testing.T) {
		err := dg.RemoveEdgeByID(n1.ID, "invalid")
		if err == nil {
			t.Errorf("expected error for invalid to node ID")
		}
	})
}

func TestGetAllNodes(t *testing.T) {
	t.Run("get all nodes", func(t *testing.T) {
		d, err := NewDag()
		if err != nil {
			t.Fatal(err)
		}
		n1 := newTestNode("node1", 1)
		n2 := newTestNode("node2", 1)

		err = d.AddNodes(n1, n2)
		if err != nil {
			t.Errorf("expected nodes to be added, got error: %s", err)
		}
		nodes := d.GetAllNodes()
		if len(nodes) != 2 {
			t.Errorf("expected 2 nodes and got %d", len(nodes))
		}
	})

	t.Run("get nodes from no-nodes dag", func(t *testing.T) {
		d, err := NewDag()
		if err != nil {
			t.Fatal(err)
		}

		nodes := d.GetAllNodes()
		if len(nodes) != 0 {
			t.Errorf("expected 0 nodes and got %d", len(nodes))
		}
	})
}

func TestHasCycle(t *testing.T) {
	t.Run("check cycle in no-cycle", func(t *testing.T) {
		d, err := NewDag()
		if err != nil {
			t.Fatal(err)
		}
		n1 := newTestNode("node1", 1)
		n2 := newTestNode("node2", 1)

		err = d.AddNodes(n1, n2)
		if err != nil {
			t.Errorf("expected nodes to be added, got error: %s", err)
		}

		err = d.AddEdge(n1.ID, n2.ID)
		if err != nil {
			t.Fatalf("failed to add edge: %v", err)
		}

		if d.HasCycle() {
			t.Errorf("expected no cycles")
		}
	})

	t.Run("check cycle in cycled dag", func(t *testing.T) {
		d, err := NewDag()
		if err != nil {
			t.Fatal(err)
		}
		n1 := newTestNode("node1", 1)
		n2 := newTestNode("node2", 1)

		err = d.AddNodes(n1, n2)
		if err != nil {
			t.Errorf("expected nodes to be added, got error: %s", err)
		}

		err = d.AddEdge(n1.ID, n2.ID)
		if err != nil {
			t.Fatalf("failed to add edge: %v", err)
		}

		err = d.AddEdge(n2.ID, n1.ID)
		if err != nil {
			t.Fatalf("failed to add edge: %v", err)
		}

		if !d.HasCycle() {
			t.Errorf("expected cycles")
		}
	})
}

func TestFindCycle(t *testing.T) {
	t.Run("check cycle in no-cycle", func(t *testing.T) {
		d, err := NewDag()
		if err != nil {
			t.Fatal(err)
		}
		n1 := newTestNode("node1", 1)
		n2 := newTestNode("node2", 1)

		err = d.AddNodes(n1, n2)
		if err != nil {
			t.Errorf("expected nodes to be added, got error: %s", err)
		}

		err = d.AddEdge(n1.ID, n2.ID)
		if err != nil {
			t.Fatalf("failed to add edge: %v", err)
		}

		nodes := d.FindCycle()
		if len(nodes) > 0 {
			t.Fatalf("expected 0 nodes, it got %v", nodes)
		}

	})

	t.Run("check cycle in cycled dag", func(t *testing.T) {
		d, err := NewDag()
		if err != nil {
			t.Fatal(err)
		}
		n1 := newTestNode("node1", 1)
		n2 := newTestNode("node2", 1)

		err = d.AddNodes(n1, n2)
		if err != nil {
			t.Errorf("expected nodes to be added, got error: %s", err)
		}

		err = d.AddEdge(n1.ID, n2.ID)
		if err != nil {
			t.Fatalf("failed to add edge: %v", err)
		}

		err = d.AddEdge(n2.ID, n1.ID)
		if err != nil {
			t.Fatalf("failed to add edge: %v", err)
		}

		nodes := d.FindCycle()
		if len(nodes) != 3 {
			t.Fatalf("expected 3 nodes, it got %v", len(nodes))
		}
	})
}
func TestInDegree(t *testing.T) {
	t.Run("inDegree counts are correct", func(t *testing.T) {
		n1 := newTestNode("n1", 1)
		n2 := newTestNode("n2", 2)
		n3 := newTestNode("n3", 3)
		n4 := newTestNode("n4", 4)

		d, err := NewDag(n1, n2, n3, n4)
		if err != nil {
			t.Fatalf("failed to create dag: %v", err)
		}

		_ = d.AddEdge("n1", "n2")
		_ = d.AddEdge("n1", "n3")
		_ = d.AddEdge("n2", "n3")
		_ = d.AddEdge("n3", "n4")

		expected := map[string]int{
			"n1": 0,
			"n2": 1,
			"n3": 2,
			"n4": 1,
		}

		actual := d.inDegree()
		for id, expectedVal := range expected {
			if actual[id] != expectedVal {
				t.Errorf("expected inDegree[%s] = %d, got %d", id, expectedVal, actual[id])
			}
		}
	})
}

func TestTopologicalSort(t *testing.T) {
	t.Run("Topological sort returns correct order", func(t *testing.T) {
		n1 := newTestNode("n1", 1)
		n2 := newTestNode("n2", 2)
		n3 := newTestNode("n3", 3)
		n4 := newTestNode("n4", 4)

		d, err := NewDag(n1, n2, n3, n4)
		if err != nil {
			t.Fatalf("failed to create dag: %v", err)
		}

		_ = d.AddEdge("n1", "n2")
		_ = d.AddEdge("n1", "n3")
		_ = d.AddEdge("n2", "n3")
		_ = d.AddEdge("n3", "n4")

		sorted, err := d.TopologicalSort()
		if err != nil {
			t.Fatalf("unexpected error from TopologicalSort: %v", err)
		}

		order := make(map[string]int)
		for i, node := range sorted {
			order[node.ID] = i
		}

		// Validate topological order
		checkEdges := [][2]string{
			{"n1", "n2"},
			{"n1", "n3"},
			{"n2", "n3"},
			{"n3", "n4"},
		}

		for _, edge := range checkEdges {
			from, to := edge[0], edge[1]
			if order[from] >= order[to] {
				t.Errorf("topological order invalid: %s should come before %s", from, to)
			}
		}
	})

	t.Run("Topological sort fails on cycle", func(t *testing.T) {
		n1 := newTestNode("n1", 1)
		n2 := newTestNode("n2", 2)

		d, err := NewDag(n1, n2)
		if err != nil {
			t.Fatalf("failed to create dag: %v", err)
		}

		_ = d.AddEdge("n1", "n2")
		_ = d.AddEdge("n2", "n1") // introduces cycle

		_, err = d.TopologicalSort()
		if err == nil {
			t.Fatalf("expected error due to cycle, but got nil")
		}
	})
}
