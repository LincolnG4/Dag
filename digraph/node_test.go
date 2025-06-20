package digraph

import (
	"testing"
)

func TestConnect(t *testing.T) {
	n1 := &Node{ID: "n1", Value: 10}
	n2 := &Node{ID: "n2", Value: 20}
	n3 := &Node{ID: "n3", Value: 30}

	t.Run("Connect new edge", func(t *testing.T) {
		err := n1.Connect(n2)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(n1.edgeTo) != 1 || n1.edgeTo[0] != n2 {
			t.Errorf("expected n2 to be in edgeTo list")
		}
	})

	t.Run("Connect same edge again should fail", func(t *testing.T) {
		err := n1.Connect(n2)
		if err == nil {
			t.Errorf("expected error on duplicate edge")
		}
	})

	t.Run("Connect nil node should fail", func(t *testing.T) {
		err := n1.Connect(nil)
		if err == nil {
			t.Errorf("expected error when connecting to nil")
		}
	})

	t.Run("Connect another unique edge", func(t *testing.T) {
		err := n1.Connect(n3)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(n1.edgeTo) != 2 {
			t.Errorf("expected 2 connected nodes, got %d", len(n1.edgeTo))
		}
	})
}
