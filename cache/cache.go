package main

import (
	"container/list"
	"errors"
)

type Cache struct {
	capacity int
	queue    *list.List
	keyMap   map[string]*list.Element
}

type entry struct {
	key   string
	value string
}

func NewCache(capacity int) Cache {
	mapp := make(map[string]*list.Element, capacity)
	return Cache{
		capacity: capacity,
		queue:    list.New(),
		keyMap:   mapp,
	}
}

func (c *Cache) Get(key string) (value string, err error) {
	element, ok := c.keyMap[key]
	if !ok {
		err = errors.New("miss cache")
		return
	}
	value = element.Value.(entry).value
	c.queue.MoveToFront(element)
	return
}

func (c *Cache) Put(key string, value string) (err error) {
	element, ok := c.keyMap[key]
	if ok {
		c.queue.MoveToFront(element)
		element.Value = value
		return
	}
	if len(c.keyMap) >= c.capacity {
		first := c.queue.Back()
		delete(c.keyMap, first.Value.(entry).key)
		c.queue.Remove(first)
	}
	e := entry{
		key:   key,
		value: value,
	}
	c.keyMap[key] = c.queue.PushFront(e)
	return
}

func (c *Cache) Remove(key string) (val string, err error) {
	val = ""
	element, ok := c.keyMap[key]
	if ok {
		delete(c.keyMap, key)
		c.queue.Remove(element)
		return
	}
	return
}
