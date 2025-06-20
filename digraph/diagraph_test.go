package digraph

import (
	"testing"
)

func newTestNode(id nodeID, val int) *Node {
	return &Node{ID: id, Value: val}
}

func TestAddNodes(t *testing.T) {
	t.Run("Add valid nodes", func(t *testing.T) {
		d, err := NewDigraph()
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
		d, _ := NewDigraph(n1)

		err := d.AddNodes(n1)
		if err == nil {
			t.Errorf("expected error for duplicate node")
		}
	})

	t.Run("NewDigraph with nil node should fail", func(t *testing.T) {
		_, err := NewDigraph(nil)
		if err == nil {
			t.Errorf("expected error for nil node")
		}
	})
}

func TestRemoveNodeByID(t *testing.T) {
	n1 := newTestNode("n1", 1)
	n2 := newTestNode("n2", 2)

	d, _ := NewDigraph(n1, n2)
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
	d, _ := NewDigraph(n)

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
	d, _ := NewDigraph(n1)

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
