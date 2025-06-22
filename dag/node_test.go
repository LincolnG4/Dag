package dag

import (
	"testing"
)

func TestConnect(t *testing.T) {
	n1 := &Node{id: "n1", Value: 10}
	n2 := &Node{id: "n2", Value: 20}
	n3 := &Node{id: "n3", Value: 30}

	t.Run("Connect new edge", func(t *testing.T) {
		err := n1.ConnectNode(n2)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(n1.edgeTo) != 1 || n1.edgeTo[0] != n2 {
			t.Errorf("expected n2 to be in edgeTo list")
		}
	})

	t.Run("Connect same edge again should fail", func(t *testing.T) {
		err := n1.ConnectNode(n2)
		if err == nil {
			t.Errorf("expected error on duplicate edge")
		}
	})

	t.Run("Connect nil node should fail", func(t *testing.T) {
		err := n1.ConnectNode(nil)
		if err == nil {
			t.Errorf("expected error when connecting to nil")
		}
	})

	t.Run("Connect another unique edge", func(t *testing.T) {
		err := n1.ConnectNode(n3)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(n1.edgeTo) != 2 {
			t.Errorf("expected 2 connected nodes, got %d", len(n1.edgeTo))
		}
	})

	t.Run("Connect node to itself should fail", func(t *testing.T) {
		err := n1.ConnectNode(n1)
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}

	})
}

func TestDisconnectNode(t *testing.T) {
	n1 := &Node{id: "n1", Value: 1}
	n2 := &Node{id: "n2", Value: 2}
	n3 := &Node{id: "n3", Value: 3}

	t.Run("Disconnect existing edge", func(t *testing.T) {
		err := n1.ConnectNode(n2)
		if err != nil {
			t.Fatalf("unexpected error while connecting: %v", err)
		}

		err = n1.DisconnectNode(n2)
		if err != nil {
			t.Errorf("expected to disconnect n2 from n1, got error: %v", err)
		}

		if len(n1.edgeTo) != 0 {
			t.Errorf("expected edgeTo to be empty after disconnect, got %d", len(n1.edgeTo))
		}
	})

	t.Run("Disconnect non-existing edge", func(t *testing.T) {
		err := n1.DisconnectNode(n3)
		if err == nil {
			t.Errorf("expected error when disconnecting non-existent edge")
		}
	})

	t.Run("Disconnect nil node", func(t *testing.T) {
		err := n1.DisconnectNode(nil)
		if err == nil {
			t.Errorf("expected error when disconnecting nil node")
		}
	})
}
