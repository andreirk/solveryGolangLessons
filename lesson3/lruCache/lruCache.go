package lrucache

import (
	"time"
)

type node struct {
	prev    *node
	next    *node
	val     int
	key     string
	expires time.Time
}

type LRUCache struct {
	m        map[string]*node
	fakeHead *node
	fakeTail *node
	cap      int
	len      int
}

func NewLRUCache(capacity int) LRUCache {
	fakeHead := &node{val: -1, prev: nil, next: nil}
	fakeTail := &node{val: -2, prev: nil, next: nil}
	fakeHead.next = fakeTail
	fakeTail.prev = fakeHead

	return LRUCache{
		cap:      capacity,
		len:      0,
		m:        make(map[string]*node, capacity),
		fakeHead: fakeHead,
		fakeTail: fakeTail,
	}
}

func (c *LRUCache) Get(key string) (int, bool) {
	foundNode, ok := c.m[key]

	if !ok {
		return -1, false
	}
	timeUntil := time.Until(foundNode.expires)

	if timeUntil < 0 {
		prev := foundNode.prev
		next := foundNode.next

		prev.next = next
		next.prev = prev

		delete(c.m, foundNode.key)
		return -1, false
	}

	if c.fakeHead.next == foundNode {
		return foundNode.val, true
	}

	c.moveToHead(foundNode)

	return foundNode.val, true
}

func (c *LRUCache) moveToHead(v *node) {
	frst := c.fakeHead.next
	next := v.next
	prev := v.prev

	c.fakeHead.next = v
	v.prev = c.fakeHead
	prev.next = next
	next.prev = prev
	frst.prev = v
	v.next = frst
}

func (c *LRUCache) Add(key string, value int, dur time.Duration) {
	now := time.Now()
	got, _ := c.Get(key)
	if got != -1 {
		c.fakeHead.next.val = value
		c.fakeHead.next.expires = now.Add(dur)
		c.m[key] = c.fakeHead.next
		return
	}

	var cur *node
	if c.len >= c.cap {
		// delete from queue and map
		c.len--
		cur = c.fakeTail.prev
		cur.prev.next = c.fakeTail
		c.fakeTail.prev = cur.prev
		delete(c.m, cur.key)
	} else {
		cur = &node{}
	}

	cur.expires = now.Add(dur)

	c.len++
	next := c.fakeHead.next
	cur.val = value
	cur.key = key
	cur.next = next
	cur.prev = c.fakeHead
	c.fakeHead.next = cur
	next.prev = cur
	c.m[key] = cur
}

func (c *LRUCache) DeleteTail() {
	c.len--
	cur := c.fakeTail.prev
	cur.prev.next = c.fakeTail
	c.fakeTail.prev = cur.prev
	delete(c.m, cur.key)

}
