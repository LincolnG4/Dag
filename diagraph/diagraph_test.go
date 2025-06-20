package diagraph

import (
	"fmt"
	"testing"
)

func TestAddNodes(t *testing.T) {
	// create a empty dag and add node
	d, err := NewDiagraph()
	if err != nil {
		t.Fatal(err)
	}
	n1 := Node{
		ID:    "node1",
		Value: 1,
	}

	n2 := Node{
		ID:    "node2",
		Value: 1,
	}

	err = d.AddNodes(&n1, &n2)
	if err != nil {
		t.Errorf("expected multiple nodes to be add to the dag, err: %s", err.Error())
	}

	err = d.AddNodes(&n1, &n2)
	if err == nil {
		t.Errorf("should fail to add not unique nodes, err: %s", err)
	}

	// create a dag with nodes and add new nodes
	d2, err := NewDiagraph(&n1, &n2)
	if err != nil {
		t.Fatal(err)
	}

	if d2.Len() != 2 {
		t.Errorf("n should expected 2, returned %d", d2.Len())
	}

	// it should fail for nil nodes
	_, err = NewDiagraph(nil)
	if err == nil {
		t.Fatal("NewDiagraph should return a error", err)
	}

}

func TestRemoveNode(t *testing.T) {
	n1 := Node{
		ID:    "node1",
		Value: 1,
	}

	n2 := Node{
		ID:    "node2",
		Value: 1,
	}

	// create a dag with nodes and add new nodes
	d, err := NewDiagraph(&n1, &n2)
	if err != nil {
		t.Fatal(err)
	}

	d.RemoveNodeByID(n2.ID)
	fmt.Println(d.nodes)
	if d.Len() != 1 {
		t.Errorf("n should expected 1, returned %d", d.Len())
	}

}

func TestCheckIfNodeExist(t *testing.T) {

}

func TestGetNodeByID(t *testing.T) {
	n1 := Node{
		ID:    "node1",
		Value: 1,
	}

	n2 := Node{
		ID:    "node2",
		Value: 1,
	}

	// create a dag with nodes and add new nodes
	d, err := NewDiagraph(&n1, &n2)
	if err != nil {
		t.Fatal(err)
	}

	// node should exist
	if !d.CheckIfNodeExist(n1.ID) {
		t.Fatal(err)
	}

	err = d.RemoveNodeByID(n1.ID)
	if err != nil {
		t.Fatal(err)
	}

	// node should not exist
	if d.CheckIfNodeExist(n1.ID) {
		t.Fatal(err)
	}
}
