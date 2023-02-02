package leetcode

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	left     *Node
	right    *Node
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		capacity: capacity,
		cache:    map[int]*Node{},

		left:  &Node{key: 0, value: 0},
		right: &Node{key: 0, value: 0},
	}
	c.left.next, c.right.prev = c.right, c.left
	return c
}

func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.cache[key]; ok {
		lru.remove(lru.cache[key])
		lru.insert(lru.cache[key])
		return lru.cache[key].value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if _, ok := lru.cache[key]; ok {
		lru.remove(lru.cache[key])
	}
	lru.cache[key] = &Node{key: key, value: value}
	lru.insert(lru.cache[key])

	if len(lru.cache) > lru.capacity {
		last := lru.left.next
		lru.remove(last)
		delete(lru.cache, last.key)
	}
}

func (lru *LRUCache) insert(node *Node) {
	prev, next := lru.right.prev, lru.right
	next.prev = node
	prev.next = next.prev
	node.next, node.prev = next, prev
}

func (lru *LRUCache) remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
