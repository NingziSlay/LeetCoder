package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 20; i++ {
		fmt.Println(FibMemCached(i))
	}
	fmt.Printf("count: %d\n", count)
}

// count Fib 函数每执行一次运算，自增 1
var count = 0
// memCache 内存缓存
var memCache = make(map[int]int)
var lruCache = New(10)

func Fib(n int) int {
	count += 1
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

// Fib 计算斐波那契数列第 N 个数的值
func FibMemCached(n int) int {
	if r, ok := memCache[n]; ok {
		return r
	}
	count += 1
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	r := FibMemCached(n-1) + FibMemCached(n-2)
	memCache[n] = r
	return r
}

func FibLRUCached(n int) int {
	v := lruCache.Get(n)
	if v != nil {
		return v.(int)
	}
	count += 1
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	r := FibLRUCached(n-1) + FibLRUCached(n-2)
	lruCache.Put(n, r)
	return r
}

// LRUCache Least Recently Used Cache
type LRUCache struct {
	root     *Ring                 // 链表 root 节点，指向最后一次使用的节点
	cache    map[interface{}]*Ring // 用户 get
	capacity int                   // 最大容量
	length   int                   // 已使用
}

// New LRUCache
func New(capacity int) *LRUCache {
	return &LRUCache{
		root:     NewRing("", nil),
		cache:    make(map[interface{}]*Ring),
		capacity: capacity,
		length:   0,
	}
}

func (l *LRUCache) Put(key interface{}, value interface{}) {
	if l.length >= l.capacity {
		l.delLRU()
	}

	ring := NewRing(key, value)

	// 设置 root 的 prev 节点
	l.root.SetNext(ring)
	// 存储 key value 值
	l.cache[key] = ring
	// 使用计数
	l.length = len(l.cache)
}

func (l *LRUCache) Get(key interface{}) interface{} {
	ring, ok := l.cache[key]
	if ok {
		return ring.Value
	}
	return nil
}

func (l *LRUCache) delLRU() {
	key := l.root.prev.Key
	l.root.prev.Del()
	delete(l.cache, key)
	l.length = len(l.cache)
}

/*
lru_cache 内部维护一个双向环形链表和一个 map
链表中的每个节点存储一个 lruCache，和其他的一下信息
map 的 key 为 cache_key, value 为链表的一个 node

这样可以实现增删改查都能做到 O(1) 的复杂度

最大容量为 6

    root
        a  ⬅️➡️  b  ⬅️➡️  c
        ⬆️             ⬆️
        ⬇️             ⬇️
        f  ⬅️➡️  e  ⬅️➡️  d

添加新的元素
    超出了最大容量，将最少使用的元素删除
    即将 root.prev 删除
    新的元素成为 root

    root
        g  ⬅️➡️  a  ⬅️➡️  b
        ⬆️             ⬆️
        ⬇️             ⬇️
        e  ⬅️➡️  d  ⬅️➡️  c

使用元素 c
    最近使用过的元素成为 root

    root
        c  ⬅️➡️  g  ⬅️➡️  a
        ⬆️             ⬆️
        ⬇️             ⬇️
        e  ⬅️➡️  d  ⬅️➡️  b
*/

type Ring struct {
	prev, next *Ring
	Key        interface{}
	Value      interface{}
}

func NewRing(key interface{}, v interface{}) *Ring {
	r := &Ring{
		Key:   key,
		Value: v,
		prev:  nil,
		next:  nil,
	}
	r.prev = r
	r.next = r
	return r
}

// SetNext
// a.SetNext(c)
// a  ⬅️➡️  b
// a  ⬅️➡️  c ⬅️➡️ b
func (r *Ring) SetNext(next *Ring) {
	next.next = r.next
	next.prev = r

	r.next.prev = next
	r.next = next
}

// SetPrev
// b.SetPrev(c)
// a  ⬅️➡️  b
// a  ⬅️➡️  c ⬅️➡️ b
func (r *Ring) SetPrev(prev *Ring) {
	prev.next = r
	prev.prev = r.prev

	r.prev.next = prev
	r.prev = prev
}

// Del
// c.Del()
// a  ⬅️➡️  c ⬅️➡️ b
// a  ⬅️➡️  b
func (r *Ring) Del() {
	r.next.prev = r.prev
	r.prev.next = r.next
}
