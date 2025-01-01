package advanced

import (
	"math/rand"
	"time"
)

const maxLevel = 16
const probability = 0.5

// SkipListNode represents a node in the skip list
type SkipListNode struct {
	key     int
	value   interface{}
	forward []*SkipListNode
}

// SkipList represents the skip list data structure
type SkipList struct {
	header *SkipListNode
	level  int
	random *rand.Rand
}

// NewSkipList creates a new skip list
func NewSkipList() *SkipList {
	return &SkipList{
		header: &SkipListNode{forward: make([]*SkipListNode, maxLevel)},
		level:  0,
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// randomLevel generates a random level for a new node
func (sl *SkipList) randomLevel() int {
	level := 0
	for level < maxLevel-1 && sl.random.Float64() < probability {
		level++
	}
	return level
}

// Insert adds a new key-value pair to the skip list
func (sl *SkipList) Insert(key int, value interface{}) {
	update := make([]*SkipListNode, maxLevel)
	current := sl.header

	// Find the position to insert
	for i := sl.level; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		update[i] = current
	}
	current = current.forward[0]

	// If key exists, update value
	if current != nil && current.key == key {
		current.value = value
		return
	}

	// Generate random level and create new node
	newLevel := sl.randomLevel()
	if newLevel > sl.level {
		for i := sl.level + 1; i <= newLevel; i++ {
			update[i] = sl.header
		}
		sl.level = newLevel
	}

	newNode := &SkipListNode{
		key:     key,
		value:   value,
		forward: make([]*SkipListNode, newLevel+1),
	}

	// Update forward pointers
	for i := 0; i <= newLevel; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
}

// Search finds a value by key in the skip list
func (sl *SkipList) Search(key int) (interface{}, bool) {
	current := sl.header

	// Start from the highest level and move down
	for i := sl.level; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
	}

	current = current.forward[0]
	if current != nil && current.key == key {
		return current.value, true
	}
	return nil, false
}

// Delete removes a key-value pair from the skip list
func (sl *SkipList) Delete(key int) bool {
	update := make([]*SkipListNode, maxLevel)
	current := sl.header

	// Find the node to delete
	for i := sl.level; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		update[i] = current
	}
	current = current.forward[0]

	// If key doesn't exist
	if current == nil || current.key != key {
		return false
	}

	// Update forward pointers
	for i := 0; i <= sl.level; i++ {
		if update[i].forward[i] != current {
			break
		}
		update[i].forward[i] = current.forward[i]
	}

	// Update level if necessary
	for sl.level > 0 && sl.header.forward[sl.level] == nil {
		sl.level--
	}

	return true
}
