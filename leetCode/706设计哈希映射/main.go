package main

import "container/list"

const base = 769

type MyHashMap struct {
	data []list.List
}

type pair struct {
	key   int
	value int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, 769)}
}

/** value will always be non-negative. */
func (m *MyHashMap) Put(key int, value int) {
	h := hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		p := e.Value.(*pair)
		if p.key == key {
			p.value = value
			return
		}
	}
	m.data[h].PushBack(&pair{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (m *MyHashMap) Get(key int) int {
	h := hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		p := e.Value.(*pair)
		if p.key == key {
			return p.value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (m *MyHashMap) Remove(key int) {
	h := hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		p := e.Value.(*pair)
		if p.key == key {
			m.data[h].Remove(e)
			return
		}
	}
}

func hash(key int) int {
	return key % base
}
