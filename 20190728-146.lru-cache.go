package main

/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// Node146 : 双向链表结点
type Node146 struct {
	key        int
	val        int
	prev, next *Node146
}

// LRUCache :
type LRUCache struct {
	capacity   int
	cache      map[int]*Node146
	head, tail *Node146
}

// Constructor :
func Constructor146(capacity int) LRUCache {
	m := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node146, capacity),
		head:     new(Node146),
		tail:     new(Node146),
	}
	m.head.next, m.tail.prev = m.tail, m.head
	return m
}

// Get : 读
func (m *LRUCache) Get(key int) int {
	if node, ok := m.cache[key]; ok {
		m.moveToEnd(node)
		return node.val
	} else {
		return -1
	}
}

// Put : 写
func (m *LRUCache) Put(key int, value int) {
	if node, ok := m.cache[key]; ok {
		node.val = value
		m.moveToEnd(node)
	} else {
		if len(m.cache) == m.capacity {
			// delete first
			first := m.head.next
			delete(m.cache, first.key)
			m.head.next, first.next.prev = first.next, m.head
		}
		node := &Node146{
			key: key,
			val: value,
		}
		m.cache[key] = node
		m.insertToEnd(node)
	}
}

func (m *LRUCache) insertToEnd(node *Node146) {
	node.prev, node.next = m.tail.prev, m.tail
	node.prev.next, node.next.prev = node, node
}

func (m *LRUCache) moveToEnd(node *Node146) {
	// unlink
	node.prev.next, node.next.prev = node.next, node.prev
	m.insertToEnd(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
