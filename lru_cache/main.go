// 需要维护一个 map 用来存储结果，还需要一个数据结构来维护key 的权重，快速定位到最少使用的key
// 从 lru_cache 中获取 key 的时候，key 的命中
package main

import "fmt"

type LRUCache struct {
	maxSize  int
	used     int
	store    map[string]*LinkedList
	rootNode *LinkedList // 最近使用的节点
}

func NewLRUCache(maxSize int) *LRUCache {
	return &LRUCache{
		maxSize:  maxSize,
		used:     0,
		store:    make(map[string]*LinkedList),
		rootNode: NewLinkedList("root", nil),
	}
}

func (lru *LRUCache) Put(key string, val interface{}) {
	// 超出最大容量时，删除最少使用的那个节点
	if lru.used+1 > lru.maxSize {
		delete(lru.store, lru.rootNode.next.key)
		lru.rootNode.next.Delete()
		lru.used -= 1
	}

	var node *LinkedList
	node = lru.store[key]
	if node != nil {
		node.val = val
		node.hit = 0
	} else {
		node = NewLinkedList(key, val)
		lru.store[key] = node
		lru.used += 1
	}
	lru.rootNode.AddPrev(node)
}

func (lru *LRUCache) Get(key string) interface{} {
	node := lru.store[key]
	if node == nil {
		return nil
	}
	// 移动到链表最前方
	lru.rootNode.AddPrev(node)
	node.hit += 1
	return node.val
}

// LinkedList 环形双向链表，链表头和尾总是相连的
type LinkedList struct {
	prev *LinkedList
	next *LinkedList
	hit  int
	key  string
	val  interface{}
}

func NewLinkedList(key string, val interface{}) *LinkedList {
	return &LinkedList{
		prev: nil,
		next: nil,
		hit:  0,
		key:  key,
		val:  val,
	}
}

// AddPrev 将新的节点添加到自身的前面
func (l *LinkedList) AddPrev(node *LinkedList) {

	// 如果 l 存在前驱节点，将 l 的前驱节点的后继节点置为 node， node 的前驱节点置为 l 的前驱节点
	// 这是 l 仍然指向前驱节点，但前驱节点已经指向了 node
	if l.prev != nil {
		l.prev.next = node
		node.prev = l.prev
	} else {
		// 如果 l 没有前驱节点（初始化时），直接将 node 的前驱节点指向 l，为了形成闭环
		node.prev = l
	}

	switch l.next {
	case nil:
		// 如果 l 没有后继节点（初始化时），将 l 的后继节点置为 node，为了形成闭环
		l.next = node
	case node:
		// 如果 l 的后继节点等于 node，需要把 node 的后继节点置为 l 的后继节点
		l.next = node.next
		// 并将 node 的后继节点的前驱置为 l，把 node 从 l 和 node 的后继节点中剥离出来
		node.next.prev = l
	}

	// node 的后继节点置为 l
	node.next = l
	// l 的前驱节点置为 node
	l.prev = node
}

// Delete 将自身从链表中删除
func (l *LinkedList) Delete() {
	l.next.prev = l.prev
	l.prev.next = l.next
}

func main() {
	lru := NewLRUCache(3)
	lru.Put("A", "aaa") // root.next = a  root.prev = a	a.next = root 	a.prev = root
	lru.Put("B", "bbb") // a = root  root.next = a
	fmt.Print(lru.Get("A"))
	lru.Put("C", "ccc")
	lru.Put("D", "ddd")
}
