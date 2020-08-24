package main

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	key := "test"
	val := 1234
	node := NewLinkedList(key, val)
	if node.key != key {
		t.Fatalf("expect key: %s, got: %s\n", key, node.key)
	}
	if node.val != val {
		t.Fatalf("expect val: %v, got: %v\n", val, node.val)
	}
	if node.hit != 0 {
		t.Fatalf("new node's hit should be 0.")
	}
	if node.prev != nil || node.next != nil {
		t.Fatalf("new node don't have prev or next node.")
	}
}

func TestLinkedList_AddPrev(t *testing.T) {
	node1 := NewLinkedList("Node1", "1")
	node2 := NewLinkedList("Node2", []int{1, 2, 3})
	node3 := NewLinkedList("Node3", map[string]interface{}{"key": 123})

	node1.AddPrev(node2)

	if node1.prev != node2 {
		t.Fatalf("1")
	}
	if node1.next != node2 {
		t.Fatalf("2")
	}

	if node2.next != node1 {
		t.Fatalf("3")
	}
	if node2.prev != node1 {
		t.Fatalf("4")
	}

	node2.AddPrev(node3)
	if node2.prev != node3 {
		t.Fatalf("5")
	}
	if node2.next != node1 {
		t.Fatalf("6")
	}
	if node3.next != node2 {
		t.Fatalf("7")
	}
	if node3.prev != node1 {
		t.Fatalf("8")
	}

	node1.AddPrev(node3)
	if node1.prev != node3 {
		t.Fatalf("9")
	}
	if node1.next != node2 {
		t.Fatalf("10")
	}
	if node3.prev != node2 {
		t.Fatalf("11")
	}
	if node3.next != node1 {
		t.Fatalf("12")
	}
	if node2.prev != node1 {
		t.Fatalf("13")
	}
	if node2.next != node3 {
		t.Fatalf("14")
	}
}

func TestLinkedList_Delete(t *testing.T) {
	node1 := NewLinkedList("Node1", "1")
	node2 := NewLinkedList("Node2", []int{1, 2, 3})
	node3 := NewLinkedList("Node3", map[string]interface{}{"key": 123})

	node1.AddPrev(node2)
	node2.AddPrev(node3)

	node1.Delete()
	if node2.next != node3 {
		t.Fatalf("")
	}
	if node3.prev != node2 {
		t.Fatalf("")
	}
}

func TestNewLRUCache(t *testing.T) {
	cache := NewLRUCache(3)
	if cache.maxSize != 3 {
		t.Fatal("maxSize wrong")
	}
	if cache.rootNode == nil {
		t.Fatal("rootNode can't be nil")
	}
}

func TestLRUCache_Put(t *testing.T) {
	cache := NewLRUCache(2)
	key := "node1"
	val := "1"
	key2 := "node2"
	val2 := "2"
	key3 := "node3"
	val3 := "3"

	cache.Put(key, val)
	cache.Put(key2, val2)
	if cache.used != 2 {
		t.Fatal("")
	}
	if cache.rootNode.prev.key != key2 {
		t.Fatal("")
	}
	if cache.rootNode.next.key != key {
		t.Fatal("")
	}

	cache.Put(key3, val3)
	if cache.used != 2 {
		t.Fatal("")
	}
	if cache.rootNode.prev.key != key3 {
		t.Fatal("")
	}
	if cache.rootNode.next.key != key2 {
		t.Fatal("")
	}

	// node1 should be removed
	if cache.Get(key) != nil {
		t.Fatal("")
	}
}
